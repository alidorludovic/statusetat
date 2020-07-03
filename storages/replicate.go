package storages

import (
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/ArthurHlt/statusetat/models"
	log "github.com/sirupsen/logrus"
)

const (
	created recordAction = iota
	updated
	deleted
)

type recordAction int

type record struct {
	storeUrl string
	action   recordAction
	guid     string
	incident models.Incident
	deleted  bool
}

type Replicate struct {
	storesInit []Store
	stores     map[string]Store
	records    *[]*record
	mu         sync.Mutex
}

func NewReplicate(storesInit []Store) *Replicate {
	records := make([]*record, 0)
	repl := &Replicate{
		storesInit: storesInit,
		stores:     make(map[string]Store),
		records:    &records,
		mu:         sync.Mutex{},
	}

	go repl.clean()
	go repl.replay()
	return repl
}

func (m Replicate) Creator() func(u *url.URL) (Store, error) {
	return func(u *url.URL) (Store, error) {
		for _, s := range m.storesInit {
			if !s.Detect(u) {
				continue
			}
			store, err := s.Creator()(u)
			if err != nil {
				return nil, err
			}
			m.stores[u.String()] = store
			return m, nil
		}
		return nil, fmt.Errorf("No valid store can be found")
	}
}

func (m Replicate) Detect(u *url.URL) bool {
	return true
}

func (m Replicate) Create(incident models.Incident) (models.Incident, error) {
	var err error
	allInError := true
	for storeUrl, s := range m.stores {
		incident, err = s.Create(incident)
		if err != nil {
			m.addRecord(storeUrl, created, "", incident)
			continue
		}
		allInError = false
	}
	if allInError {
		return incident, err
	}
	return incident, nil
}

func (m Replicate) Update(guid string, incident models.Incident) (models.Incident, error) {
	var err error
	allInError := true
	for storeUrl, s := range m.stores {
		incident, err = s.Update(guid, incident)
		if err != nil {
			m.addRecord(storeUrl, created, guid, incident)
			continue
		}
		allInError = false
	}
	if allInError {
		return incident, err
	}
	return incident, nil
}

func (m Replicate) Delete(guid string) error {
	var err error
	allInError := true
	for storeUrl, s := range m.stores {
		err = s.Delete(guid)
		if err != nil {
			m.addRecord(storeUrl, deleted, guid, models.Incident{})
			continue
		}
		allInError = false
	}
	if allInError {
		return err
	}
	return nil
}

func (m *Replicate) addRecord(storeUrl string, action recordAction, guid string, incident models.Incident) {
	log.
		WithField("action", action).
		WithField("url", storeUrl).
		Debug("Add record to replay")
	records := *m.records
	records = append(records, &record{
		storeUrl: storeUrl,
		action:   action,
		guid:     guid,
		incident: incident,
		deleted:  false,
	})
	m.mu.Lock()
	defer m.mu.Unlock()
	*m.records = records
}

func (m Replicate) clean() {
	for {
		m.mu.Lock()
		finalRecord := make([]*record, 0)
		for _, record := range *m.records {
			if record.deleted {
				continue
			}
			finalRecord = append(finalRecord, record)
		}
		*m.records = finalRecord
		m.mu.Unlock()
		time.Sleep(1 * time.Hour)
	}
}

func (m Replicate) replay() {
	for {
		m.mu.Lock()
		todo := make([]*record, 0)
		for _, record := range *m.records {
			if record.deleted {
				continue
			}
			todo = append(todo, record)
		}
		m.mu.Unlock()

		for _, record := range todo {
			log.
				WithField("action", record.action).
				WithField("url", record.storeUrl).
				Debug("Playing record ...")
			store := m.stores[record.storeUrl]
			switch record.action {
			case created:
				_, err := store.Create(record.incident)
				if err != nil {
					continue
				}
			case updated:
				_, err := store.Update(record.guid, record.incident)
				if err != nil {
					continue
				}
			case deleted:
				err := store.Delete(record.guid)
				if err != nil {
					continue
				}
			}
			record.deleted = true
		}

		time.Sleep(90 * time.Second)
	}
}

func (m Replicate) Read(guid string) (models.Incident, error) {
	var incident models.Incident
	var err error
	for _, s := range m.stores {
		incident, err = s.Read(guid)
		if err != nil {
			continue
		}
		return incident, nil
	}
	return incident, err
}

func (m Replicate) ByDate(from, to time.Time) ([]models.Incident, error) {
	incidents := make([]models.Incident, 0)
	var err error
	for _, s := range m.stores {
		incidents, err = s.ByDate(from, to)
		if err != nil {
			continue
		}
		return incidents, nil
	}
	return incidents, err
}

func (m Replicate) Ping() error {
	return nil
}

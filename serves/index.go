package serves

import (
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/ArthurHlt/statusetat/config"
	"github.com/ArthurHlt/statusetat/models"
	"github.com/gorilla/mux"
)

type IndexData struct {
	GroupComponentState map[string]models.ComponentState
	ComponentStatesData map[string][]*ComponentStateData
	Timeline            map[string][]models.Incident
	TimelineDates       []string
	Scheduled           []models.Incident
	BaseInfo            config.BaseInfo
	Timezone            string
}

type timeSlice []string

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	first, err := time.Parse("Jan 02, 2006", p[i])
	if err != nil {
		fmt.Println(err.Error())
	}

	second, err := time.Parse("Jan 02, 2006", p[j])
	if err != nil {
		fmt.Println(err.Error())
	}
	return first.Before(second)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type ComponentStateData struct {
	Name        string
	Description string
	State       models.ComponentState
}

func (a Serve) Index(w http.ResponseWriter, req *http.Request) {

	incidents, err := a.incidentsByParamsDate(req)
	if err != nil {
		HTMLError(w, err, http.StatusInternalServerError)
		return
	}

	componentStatesByGroup := make(map[string][]*ComponentStateData)
	componentStateMap := make(map[string]*ComponentStateData)
	for _, component := range a.components {
		componentState := &ComponentStateData{
			Name:        component.Name,
			Description: component.Description,
			State:       models.Operational,
		}
		_, ok := componentStatesByGroup[component.Group]
		if ok {
			componentStatesByGroup[component.Group] = append(componentStatesByGroup[component.Group], componentState)
		} else {
			componentStatesByGroup[component.Group] = []*ComponentStateData{componentState}
		}

		componentStateMap[component.String()] = componentState
	}

	compStateGroup := make(map[string]models.ComponentState)
	for k, _ := range a.components.Regroups() {
		compStateGroup[k] = models.Operational
	}

	timeline := make(map[string][]models.Incident)
	timelineDates := make(timeSlice, 0)
	for _, incident := range incidents {
		date := a.timelineFormat(incident.CreatedAt)
		if _, ok := timeline[date]; ok {
			timeline[date] = append(timeline[date], incident)
		} else {
			timeline[date] = []models.Incident{incident}
			timelineDates = append(timelineDates, date)
		}
		if incident.State == models.Resolved || incident.Components == nil {
			continue
		}

		for _, component := range *incident.Components {
			componentState, ok := componentStateMap[component.String()]
			if !ok {
				continue
			}
			if incident.ComponentState > componentState.State {
				componentState.State = incident.ComponentState
			}

			if compGroupState, ok := compStateGroup[component.Group]; ok && incident.ComponentState > compGroupState {
				compStateGroup[component.Group] = incident.ComponentState
			}
		}

	}

	from, err := a.parseDate(req, "from", time.Now().In(a.Location(req)))
	if err != nil {
		HTMLError(w, err, http.StatusInternalServerError)
		return
	}

	for i := 0; i <= 6; i++ {
		date := a.timelineFormat(from.AddDate(0, 0, -i))
		if _, ok := timeline[date]; !ok {
			timeline[date] = []models.Incident{}
			timelineDates = append(timelineDates, date)
		}
	}

	scheduled, err := a.scheduled(req)
	if err != nil {
		HTMLError(w, err, http.StatusInternalServerError)
		return
	}

	timezone := ""
	if !a.IsDefaultLocation(req) {
		timezone = a.Location(req).String()
	}
	sort.Sort(sort.Reverse(timelineDates))
	err = a.xt.ExecuteTemplate(w, "incidents.html", IndexData{
		BaseInfo:            a.baseInfo,
		GroupComponentState: compStateGroup,
		ComponentStatesData: componentStatesByGroup,
		Timeline:            timeline,
		TimelineDates:       timelineDates,
		Scheduled:           scheduled,
		Timezone:            timezone,
	})
	if err != nil {
		HTMLError(w, err, http.StatusInternalServerError)
		return
	}
}

func (a Serve) ShowIncident(w http.ResponseWriter, req *http.Request) {
	v := mux.Vars(req)
	guid := v["guid"]
	incident, err := a.store.Read(guid)
	if err != nil {
		HTMLError(w, err, http.StatusInternalServerError)
		return
	}
	timezone := ""
	if !a.IsDefaultLocation(req) {
		timezone = a.Location(req).String()
	}
	err = a.xt.ExecuteTemplate(w, "one_incident.html", struct {
		BaseInfo config.BaseInfo
		Incident models.Incident
		Timezone string
	}{
		BaseInfo: a.baseInfo,
		Incident: incident,
		Timezone: timezone,
	})
	if err != nil {
		HTMLError(w, err, http.StatusInternalServerError)
		return
	}
}

func (a Serve) History(w http.ResponseWriter, req *http.Request) {
	loc := a.Location(req)
	from, err := a.parseDate(req, "from", time.Now().In(loc))
	if err != nil {
		HTMLError(w, err, http.StatusInternalServerError)
		return
	}

	after := from.Add(7 * 24 * time.Hour)

	incidents, err := a.incidentsByParamsDate(req)
	if err != nil {
		HTMLError(w, err, http.StatusInternalServerError)
		return
	}

	timeline := make(map[string][]models.Incident)
	timelineDates := make(timeSlice, 0)
	for _, incident := range incidents {
		date := a.timelineFormat(incident.CreatedAt)
		if _, ok := timeline[date]; ok {
			timeline[date] = append(timeline[date], incident)
		} else {
			timeline[date] = []models.Incident{incident}
			timelineDates = append(timelineDates, date)
		}
		if incident.State == models.Resolved || incident.Components == nil {
			continue
		}

	}

	var before time.Time
	var subDate time.Time
	for i := 0; i <= 6; i++ {
		subDate = from.AddDate(0, 0, -i)
		date := a.timelineFormat(subDate)
		if _, ok := timeline[date]; !ok {
			timeline[date] = []models.Incident{}
			timelineDates = append(timelineDates, date)
		}
	}
	before = subDate.AddDate(0, 0, -1)

	timezone := ""
	if !a.IsDefaultLocation(req) {
		timezone = a.Location(req).String()
	}

	sort.Sort(sort.Reverse(timelineDates))
	err = a.xt.ExecuteTemplate(w, "history.html", struct {
		IndexData
		Before time.Time
		After  time.Time
	}{
		IndexData: IndexData{
			BaseInfo:      a.baseInfo,
			Timeline:      timeline,
			TimelineDates: timelineDates,
			Timezone:      timezone,
		},
		After:  after,
		Before: before,
	})
	if err != nil {
		HTMLError(w, err, http.StatusInternalServerError)
		return
	}
}

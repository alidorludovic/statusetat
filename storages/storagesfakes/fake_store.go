// Code generated by counterfeiter. DO NOT EDIT.
package storagesfakes

import (
	"net/url"
	"sync"
	"time"

	"github.com/orange-cloudfoundry/statusetat/models"
	"github.com/orange-cloudfoundry/statusetat/storages"
)

type FakeStore struct {
	ByDateStub        func(time.Time, time.Time) ([]models.Incident, error)
	byDateMutex       sync.RWMutex
	byDateArgsForCall []struct {
		arg1 time.Time
		arg2 time.Time
	}
	byDateReturns struct {
		result1 []models.Incident
		result2 error
	}
	byDateReturnsOnCall map[int]struct {
		result1 []models.Incident
		result2 error
	}
	CreateStub        func(models.Incident) (models.Incident, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 models.Incident
	}
	createReturns struct {
		result1 models.Incident
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 models.Incident
		result2 error
	}
	CreatorStub        func() func(u *url.URL) (storages.Store, error)
	creatorMutex       sync.RWMutex
	creatorArgsForCall []struct {
	}
	creatorReturns struct {
		result1 func(u *url.URL) (storages.Store, error)
	}
	creatorReturnsOnCall map[int]struct {
		result1 func(u *url.URL) (storages.Store, error)
	}
	DeleteStub        func(string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DetectStub        func(*url.URL) bool
	detectMutex       sync.RWMutex
	detectArgsForCall []struct {
		arg1 *url.URL
	}
	detectReturns struct {
		result1 bool
	}
	detectReturnsOnCall map[int]struct {
		result1 bool
	}
	PersistentsStub        func() ([]models.Incident, error)
	persistentsMutex       sync.RWMutex
	persistentsArgsForCall []struct {
	}
	persistentsReturns struct {
		result1 []models.Incident
		result2 error
	}
	persistentsReturnsOnCall map[int]struct {
		result1 []models.Incident
		result2 error
	}
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct {
	}
	pingReturns struct {
		result1 error
	}
	pingReturnsOnCall map[int]struct {
		result1 error
	}
	ReadStub        func(string) (models.Incident, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		arg1 string
	}
	readReturns struct {
		result1 models.Incident
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 models.Incident
		result2 error
	}
	SubscribeStub        func(string) error
	subscribeMutex       sync.RWMutex
	subscribeArgsForCall []struct {
		arg1 string
	}
	subscribeReturns struct {
		result1 error
	}
	subscribeReturnsOnCall map[int]struct {
		result1 error
	}
	SubscribersStub        func() ([]string, error)
	subscribersMutex       sync.RWMutex
	subscribersArgsForCall []struct {
	}
	subscribersReturns struct {
		result1 []string
		result2 error
	}
	subscribersReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	UnsubscribeStub        func(string) error
	unsubscribeMutex       sync.RWMutex
	unsubscribeArgsForCall []struct {
		arg1 string
	}
	unsubscribeReturns struct {
		result1 error
	}
	unsubscribeReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(string, models.Incident) (models.Incident, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 string
		arg2 models.Incident
	}
	updateReturns struct {
		result1 models.Incident
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 models.Incident
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStore) ByDate(arg1 time.Time, arg2 time.Time) ([]models.Incident, error) {
	fake.byDateMutex.Lock()
	ret, specificReturn := fake.byDateReturnsOnCall[len(fake.byDateArgsForCall)]
	fake.byDateArgsForCall = append(fake.byDateArgsForCall, struct {
		arg1 time.Time
		arg2 time.Time
	}{arg1, arg2})
	fake.recordInvocation("ByDate", []interface{}{arg1, arg2})
	fake.byDateMutex.Unlock()
	if fake.ByDateStub != nil {
		return fake.ByDateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.byDateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) ByDateCallCount() int {
	fake.byDateMutex.RLock()
	defer fake.byDateMutex.RUnlock()
	return len(fake.byDateArgsForCall)
}

func (fake *FakeStore) ByDateCalls(stub func(time.Time, time.Time) ([]models.Incident, error)) {
	fake.byDateMutex.Lock()
	defer fake.byDateMutex.Unlock()
	fake.ByDateStub = stub
}

func (fake *FakeStore) ByDateArgsForCall(i int) (time.Time, time.Time) {
	fake.byDateMutex.RLock()
	defer fake.byDateMutex.RUnlock()
	argsForCall := fake.byDateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) ByDateReturns(result1 []models.Incident, result2 error) {
	fake.byDateMutex.Lock()
	defer fake.byDateMutex.Unlock()
	fake.ByDateStub = nil
	fake.byDateReturns = struct {
		result1 []models.Incident
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) ByDateReturnsOnCall(i int, result1 []models.Incident, result2 error) {
	fake.byDateMutex.Lock()
	defer fake.byDateMutex.Unlock()
	fake.ByDateStub = nil
	if fake.byDateReturnsOnCall == nil {
		fake.byDateReturnsOnCall = make(map[int]struct {
			result1 []models.Incident
			result2 error
		})
	}
	fake.byDateReturnsOnCall[i] = struct {
		result1 []models.Incident
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) Create(arg1 models.Incident) (models.Incident, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 models.Incident
	}{arg1})
	fake.recordInvocation("Create", []interface{}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeStore) CreateCalls(stub func(models.Incident) (models.Incident, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeStore) CreateArgsForCall(i int) models.Incident {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) CreateReturns(result1 models.Incident, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 models.Incident
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) CreateReturnsOnCall(i int, result1 models.Incident, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 models.Incident
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 models.Incident
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) Creator() func(u *url.URL) (storages.Store, error) {
	fake.creatorMutex.Lock()
	ret, specificReturn := fake.creatorReturnsOnCall[len(fake.creatorArgsForCall)]
	fake.creatorArgsForCall = append(fake.creatorArgsForCall, struct {
	}{})
	fake.recordInvocation("Creator", []interface{}{})
	fake.creatorMutex.Unlock()
	if fake.CreatorStub != nil {
		return fake.CreatorStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.creatorReturns
	return fakeReturns.result1
}

func (fake *FakeStore) CreatorCallCount() int {
	fake.creatorMutex.RLock()
	defer fake.creatorMutex.RUnlock()
	return len(fake.creatorArgsForCall)
}

func (fake *FakeStore) CreatorCalls(stub func() func(u *url.URL) (storages.Store, error)) {
	fake.creatorMutex.Lock()
	defer fake.creatorMutex.Unlock()
	fake.CreatorStub = stub
}

func (fake *FakeStore) CreatorReturns(result1 func(u *url.URL) (storages.Store, error)) {
	fake.creatorMutex.Lock()
	defer fake.creatorMutex.Unlock()
	fake.CreatorStub = nil
	fake.creatorReturns = struct {
		result1 func(u *url.URL) (storages.Store, error)
	}{result1}
}

func (fake *FakeStore) CreatorReturnsOnCall(i int, result1 func(u *url.URL) (storages.Store, error)) {
	fake.creatorMutex.Lock()
	defer fake.creatorMutex.Unlock()
	fake.CreatorStub = nil
	if fake.creatorReturnsOnCall == nil {
		fake.creatorReturnsOnCall = make(map[int]struct {
			result1 func(u *url.URL) (storages.Store, error)
		})
	}
	fake.creatorReturnsOnCall[i] = struct {
		result1 func(u *url.URL) (storages.Store, error)
	}{result1}
}

func (fake *FakeStore) Delete(arg1 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeStore) DeleteCalls(stub func(string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeStore) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) Detect(arg1 *url.URL) bool {
	fake.detectMutex.Lock()
	ret, specificReturn := fake.detectReturnsOnCall[len(fake.detectArgsForCall)]
	fake.detectArgsForCall = append(fake.detectArgsForCall, struct {
		arg1 *url.URL
	}{arg1})
	fake.recordInvocation("Detect", []interface{}{arg1})
	fake.detectMutex.Unlock()
	if fake.DetectStub != nil {
		return fake.DetectStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.detectReturns
	return fakeReturns.result1
}

func (fake *FakeStore) DetectCallCount() int {
	fake.detectMutex.RLock()
	defer fake.detectMutex.RUnlock()
	return len(fake.detectArgsForCall)
}

func (fake *FakeStore) DetectCalls(stub func(*url.URL) bool) {
	fake.detectMutex.Lock()
	defer fake.detectMutex.Unlock()
	fake.DetectStub = stub
}

func (fake *FakeStore) DetectArgsForCall(i int) *url.URL {
	fake.detectMutex.RLock()
	defer fake.detectMutex.RUnlock()
	argsForCall := fake.detectArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) DetectReturns(result1 bool) {
	fake.detectMutex.Lock()
	defer fake.detectMutex.Unlock()
	fake.DetectStub = nil
	fake.detectReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStore) DetectReturnsOnCall(i int, result1 bool) {
	fake.detectMutex.Lock()
	defer fake.detectMutex.Unlock()
	fake.DetectStub = nil
	if fake.detectReturnsOnCall == nil {
		fake.detectReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.detectReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeStore) Persistents() ([]models.Incident, error) {
	fake.persistentsMutex.Lock()
	ret, specificReturn := fake.persistentsReturnsOnCall[len(fake.persistentsArgsForCall)]
	fake.persistentsArgsForCall = append(fake.persistentsArgsForCall, struct {
	}{})
	fake.recordInvocation("Persistents", []interface{}{})
	fake.persistentsMutex.Unlock()
	if fake.PersistentsStub != nil {
		return fake.PersistentsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.persistentsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) PersistentsCallCount() int {
	fake.persistentsMutex.RLock()
	defer fake.persistentsMutex.RUnlock()
	return len(fake.persistentsArgsForCall)
}

func (fake *FakeStore) PersistentsCalls(stub func() ([]models.Incident, error)) {
	fake.persistentsMutex.Lock()
	defer fake.persistentsMutex.Unlock()
	fake.PersistentsStub = stub
}

func (fake *FakeStore) PersistentsReturns(result1 []models.Incident, result2 error) {
	fake.persistentsMutex.Lock()
	defer fake.persistentsMutex.Unlock()
	fake.PersistentsStub = nil
	fake.persistentsReturns = struct {
		result1 []models.Incident
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) PersistentsReturnsOnCall(i int, result1 []models.Incident, result2 error) {
	fake.persistentsMutex.Lock()
	defer fake.persistentsMutex.Unlock()
	fake.PersistentsStub = nil
	if fake.persistentsReturnsOnCall == nil {
		fake.persistentsReturnsOnCall = make(map[int]struct {
			result1 []models.Incident
			result2 error
		})
	}
	fake.persistentsReturnsOnCall[i] = struct {
		result1 []models.Incident
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) Ping() error {
	fake.pingMutex.Lock()
	ret, specificReturn := fake.pingReturnsOnCall[len(fake.pingArgsForCall)]
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct {
	}{})
	fake.recordInvocation("Ping", []interface{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pingReturns
	return fakeReturns.result1
}

func (fake *FakeStore) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeStore) PingCalls(stub func() error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = stub
}

func (fake *FakeStore) PingReturns(result1 error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) PingReturnsOnCall(i int, result1 error) {
	fake.pingMutex.Lock()
	defer fake.pingMutex.Unlock()
	fake.PingStub = nil
	if fake.pingReturnsOnCall == nil {
		fake.pingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) Read(arg1 string) (models.Incident, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Read", []interface{}{arg1})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.readReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeStore) ReadCalls(stub func(string) (models.Incident, error)) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = stub
}

func (fake *FakeStore) ReadArgsForCall(i int) string {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	argsForCall := fake.readArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) ReadReturns(result1 models.Incident, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 models.Incident
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) ReadReturnsOnCall(i int, result1 models.Incident, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 models.Incident
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 models.Incident
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) Subscribe(arg1 string) error {
	fake.subscribeMutex.Lock()
	ret, specificReturn := fake.subscribeReturnsOnCall[len(fake.subscribeArgsForCall)]
	fake.subscribeArgsForCall = append(fake.subscribeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Subscribe", []interface{}{arg1})
	fake.subscribeMutex.Unlock()
	if fake.SubscribeStub != nil {
		return fake.SubscribeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.subscribeReturns
	return fakeReturns.result1
}

func (fake *FakeStore) SubscribeCallCount() int {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	return len(fake.subscribeArgsForCall)
}

func (fake *FakeStore) SubscribeCalls(stub func(string) error) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = stub
}

func (fake *FakeStore) SubscribeArgsForCall(i int) string {
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	argsForCall := fake.subscribeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) SubscribeReturns(result1 error) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = nil
	fake.subscribeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) SubscribeReturnsOnCall(i int, result1 error) {
	fake.subscribeMutex.Lock()
	defer fake.subscribeMutex.Unlock()
	fake.SubscribeStub = nil
	if fake.subscribeReturnsOnCall == nil {
		fake.subscribeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.subscribeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) Subscribers() ([]string, error) {
	fake.subscribersMutex.Lock()
	ret, specificReturn := fake.subscribersReturnsOnCall[len(fake.subscribersArgsForCall)]
	fake.subscribersArgsForCall = append(fake.subscribersArgsForCall, struct {
	}{})
	fake.recordInvocation("Subscribers", []interface{}{})
	fake.subscribersMutex.Unlock()
	if fake.SubscribersStub != nil {
		return fake.SubscribersStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.subscribersReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) SubscribersCallCount() int {
	fake.subscribersMutex.RLock()
	defer fake.subscribersMutex.RUnlock()
	return len(fake.subscribersArgsForCall)
}

func (fake *FakeStore) SubscribersCalls(stub func() ([]string, error)) {
	fake.subscribersMutex.Lock()
	defer fake.subscribersMutex.Unlock()
	fake.SubscribersStub = stub
}

func (fake *FakeStore) SubscribersReturns(result1 []string, result2 error) {
	fake.subscribersMutex.Lock()
	defer fake.subscribersMutex.Unlock()
	fake.SubscribersStub = nil
	fake.subscribersReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) SubscribersReturnsOnCall(i int, result1 []string, result2 error) {
	fake.subscribersMutex.Lock()
	defer fake.subscribersMutex.Unlock()
	fake.SubscribersStub = nil
	if fake.subscribersReturnsOnCall == nil {
		fake.subscribersReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.subscribersReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) Unsubscribe(arg1 string) error {
	fake.unsubscribeMutex.Lock()
	ret, specificReturn := fake.unsubscribeReturnsOnCall[len(fake.unsubscribeArgsForCall)]
	fake.unsubscribeArgsForCall = append(fake.unsubscribeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Unsubscribe", []interface{}{arg1})
	fake.unsubscribeMutex.Unlock()
	if fake.UnsubscribeStub != nil {
		return fake.UnsubscribeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.unsubscribeReturns
	return fakeReturns.result1
}

func (fake *FakeStore) UnsubscribeCallCount() int {
	fake.unsubscribeMutex.RLock()
	defer fake.unsubscribeMutex.RUnlock()
	return len(fake.unsubscribeArgsForCall)
}

func (fake *FakeStore) UnsubscribeCalls(stub func(string) error) {
	fake.unsubscribeMutex.Lock()
	defer fake.unsubscribeMutex.Unlock()
	fake.UnsubscribeStub = stub
}

func (fake *FakeStore) UnsubscribeArgsForCall(i int) string {
	fake.unsubscribeMutex.RLock()
	defer fake.unsubscribeMutex.RUnlock()
	argsForCall := fake.unsubscribeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStore) UnsubscribeReturns(result1 error) {
	fake.unsubscribeMutex.Lock()
	defer fake.unsubscribeMutex.Unlock()
	fake.UnsubscribeStub = nil
	fake.unsubscribeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) UnsubscribeReturnsOnCall(i int, result1 error) {
	fake.unsubscribeMutex.Lock()
	defer fake.unsubscribeMutex.Unlock()
	fake.UnsubscribeStub = nil
	if fake.unsubscribeReturnsOnCall == nil {
		fake.unsubscribeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unsubscribeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStore) Update(arg1 string, arg2 models.Incident) (models.Incident, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 string
		arg2 models.Incident
	}{arg1, arg2})
	fake.recordInvocation("Update", []interface{}{arg1, arg2})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeStore) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeStore) UpdateCalls(stub func(string, models.Incident) (models.Incident, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeStore) UpdateArgsForCall(i int) (string, models.Incident) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeStore) UpdateReturns(result1 models.Incident, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 models.Incident
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) UpdateReturnsOnCall(i int, result1 models.Incident, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 models.Incident
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 models.Incident
		result2 error
	}{result1, result2}
}

func (fake *FakeStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.byDateMutex.RLock()
	defer fake.byDateMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.creatorMutex.RLock()
	defer fake.creatorMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.detectMutex.RLock()
	defer fake.detectMutex.RUnlock()
	fake.persistentsMutex.RLock()
	defer fake.persistentsMutex.RUnlock()
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.subscribeMutex.RLock()
	defer fake.subscribeMutex.RUnlock()
	fake.subscribersMutex.RLock()
	defer fake.subscribersMutex.RUnlock()
	fake.unsubscribeMutex.RLock()
	defer fake.unsubscribeMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStore) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ storages.Store = new(FakeStore)

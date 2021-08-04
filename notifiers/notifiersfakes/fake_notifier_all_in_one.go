// Code generated by counterfeiter. DO NOT EDIT.
package notifiersfakes

import (
	"sync"

	"github.com/orange-cloudfoundry/statusetat/config"
	"github.com/orange-cloudfoundry/statusetat/models"
	"github.com/orange-cloudfoundry/statusetat/notifiers"
)

type FakeNotifierAllInOne struct {
	CreatorStub        func(map[string]interface{}, config.BaseInfo) (notifiers.Notifier, error)
	creatorMutex       sync.RWMutex
	creatorArgsForCall []struct {
		arg1 map[string]interface{}
		arg2 config.BaseInfo
	}
	creatorReturns struct {
		result1 notifiers.Notifier
		result2 error
	}
	creatorReturnsOnCall map[int]struct {
		result1 notifiers.Notifier
		result2 error
	}
	IdStub        func() string
	idMutex       sync.RWMutex
	idArgsForCall []struct {
	}
	idReturns struct {
		result1 string
	}
	idReturnsOnCall map[int]struct {
		result1 string
	}
	MetadataFieldsStub        func() []models.MetadataField
	metadataFieldsMutex       sync.RWMutex
	metadataFieldsArgsForCall []struct {
	}
	metadataFieldsReturns struct {
		result1 []models.MetadataField
	}
	metadataFieldsReturnsOnCall map[int]struct {
		result1 []models.MetadataField
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	NotifyStub        func(*models.NotifyRequest) error
	notifyMutex       sync.RWMutex
	notifyArgsForCall []struct {
		arg1 *models.NotifyRequest
	}
	notifyReturns struct {
		result1 error
	}
	notifyReturnsOnCall map[int]struct {
		result1 error
	}
	PreCheckStub        func(models.Incident) error
	preCheckMutex       sync.RWMutex
	preCheckArgsForCall []struct {
		arg1 models.Incident
	}
	preCheckReturns struct {
		result1 error
	}
	preCheckReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNotifierAllInOne) Creator(arg1 map[string]interface{}, arg2 config.BaseInfo) (notifiers.Notifier, error) {
	fake.creatorMutex.Lock()
	ret, specificReturn := fake.creatorReturnsOnCall[len(fake.creatorArgsForCall)]
	fake.creatorArgsForCall = append(fake.creatorArgsForCall, struct {
		arg1 map[string]interface{}
		arg2 config.BaseInfo
	}{arg1, arg2})
	fake.recordInvocation("Creator", []interface{}{arg1, arg2})
	fake.creatorMutex.Unlock()
	if fake.CreatorStub != nil {
		return fake.CreatorStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.creatorReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNotifierAllInOne) CreatorCallCount() int {
	fake.creatorMutex.RLock()
	defer fake.creatorMutex.RUnlock()
	return len(fake.creatorArgsForCall)
}

func (fake *FakeNotifierAllInOne) CreatorCalls(stub func(map[string]interface{}, config.BaseInfo) (notifiers.Notifier, error)) {
	fake.creatorMutex.Lock()
	defer fake.creatorMutex.Unlock()
	fake.CreatorStub = stub
}

func (fake *FakeNotifierAllInOne) CreatorArgsForCall(i int) (map[string]interface{}, config.BaseInfo) {
	fake.creatorMutex.RLock()
	defer fake.creatorMutex.RUnlock()
	argsForCall := fake.creatorArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNotifierAllInOne) CreatorReturns(result1 notifiers.Notifier, result2 error) {
	fake.creatorMutex.Lock()
	defer fake.creatorMutex.Unlock()
	fake.CreatorStub = nil
	fake.creatorReturns = struct {
		result1 notifiers.Notifier
		result2 error
	}{result1, result2}
}

func (fake *FakeNotifierAllInOne) CreatorReturnsOnCall(i int, result1 notifiers.Notifier, result2 error) {
	fake.creatorMutex.Lock()
	defer fake.creatorMutex.Unlock()
	fake.CreatorStub = nil
	if fake.creatorReturnsOnCall == nil {
		fake.creatorReturnsOnCall = make(map[int]struct {
			result1 notifiers.Notifier
			result2 error
		})
	}
	fake.creatorReturnsOnCall[i] = struct {
		result1 notifiers.Notifier
		result2 error
	}{result1, result2}
}

func (fake *FakeNotifierAllInOne) Id() string {
	fake.idMutex.Lock()
	ret, specificReturn := fake.idReturnsOnCall[len(fake.idArgsForCall)]
	fake.idArgsForCall = append(fake.idArgsForCall, struct {
	}{})
	fake.recordInvocation("Id", []interface{}{})
	fake.idMutex.Unlock()
	if fake.IdStub != nil {
		return fake.IdStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.idReturns
	return fakeReturns.result1
}

func (fake *FakeNotifierAllInOne) IdCallCount() int {
	fake.idMutex.RLock()
	defer fake.idMutex.RUnlock()
	return len(fake.idArgsForCall)
}

func (fake *FakeNotifierAllInOne) IdCalls(stub func() string) {
	fake.idMutex.Lock()
	defer fake.idMutex.Unlock()
	fake.IdStub = stub
}

func (fake *FakeNotifierAllInOne) IdReturns(result1 string) {
	fake.idMutex.Lock()
	defer fake.idMutex.Unlock()
	fake.IdStub = nil
	fake.idReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeNotifierAllInOne) IdReturnsOnCall(i int, result1 string) {
	fake.idMutex.Lock()
	defer fake.idMutex.Unlock()
	fake.IdStub = nil
	if fake.idReturnsOnCall == nil {
		fake.idReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.idReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeNotifierAllInOne) MetadataFields() []models.MetadataField {
	fake.metadataFieldsMutex.Lock()
	ret, specificReturn := fake.metadataFieldsReturnsOnCall[len(fake.metadataFieldsArgsForCall)]
	fake.metadataFieldsArgsForCall = append(fake.metadataFieldsArgsForCall, struct {
	}{})
	fake.recordInvocation("MetadataFields", []interface{}{})
	fake.metadataFieldsMutex.Unlock()
	if fake.MetadataFieldsStub != nil {
		return fake.MetadataFieldsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.metadataFieldsReturns
	return fakeReturns.result1
}

func (fake *FakeNotifierAllInOne) MetadataFieldsCallCount() int {
	fake.metadataFieldsMutex.RLock()
	defer fake.metadataFieldsMutex.RUnlock()
	return len(fake.metadataFieldsArgsForCall)
}

func (fake *FakeNotifierAllInOne) MetadataFieldsCalls(stub func() []models.MetadataField) {
	fake.metadataFieldsMutex.Lock()
	defer fake.metadataFieldsMutex.Unlock()
	fake.MetadataFieldsStub = stub
}

func (fake *FakeNotifierAllInOne) MetadataFieldsReturns(result1 []models.MetadataField) {
	fake.metadataFieldsMutex.Lock()
	defer fake.metadataFieldsMutex.Unlock()
	fake.MetadataFieldsStub = nil
	fake.metadataFieldsReturns = struct {
		result1 []models.MetadataField
	}{result1}
}

func (fake *FakeNotifierAllInOne) MetadataFieldsReturnsOnCall(i int, result1 []models.MetadataField) {
	fake.metadataFieldsMutex.Lock()
	defer fake.metadataFieldsMutex.Unlock()
	fake.MetadataFieldsStub = nil
	if fake.metadataFieldsReturnsOnCall == nil {
		fake.metadataFieldsReturnsOnCall = make(map[int]struct {
			result1 []models.MetadataField
		})
	}
	fake.metadataFieldsReturnsOnCall[i] = struct {
		result1 []models.MetadataField
	}{result1}
}

func (fake *FakeNotifierAllInOne) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *FakeNotifierAllInOne) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeNotifierAllInOne) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeNotifierAllInOne) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeNotifierAllInOne) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeNotifierAllInOne) Notify(arg1 *models.NotifyRequest) error {
	fake.notifyMutex.Lock()
	ret, specificReturn := fake.notifyReturnsOnCall[len(fake.notifyArgsForCall)]
	fake.notifyArgsForCall = append(fake.notifyArgsForCall, struct {
		arg1 *models.NotifyRequest
	}{arg1})
	fake.recordInvocation("Notify", []interface{}{arg1})
	fake.notifyMutex.Unlock()
	if fake.NotifyStub != nil {
		return fake.NotifyStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.notifyReturns
	return fakeReturns.result1
}

func (fake *FakeNotifierAllInOne) NotifyCallCount() int {
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	return len(fake.notifyArgsForCall)
}

func (fake *FakeNotifierAllInOne) NotifyCalls(stub func(*models.NotifyRequest) error) {
	fake.notifyMutex.Lock()
	defer fake.notifyMutex.Unlock()
	fake.NotifyStub = stub
}

func (fake *FakeNotifierAllInOne) NotifyArgsForCall(i int) *models.NotifyRequest {
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	argsForCall := fake.notifyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNotifierAllInOne) NotifyReturns(result1 error) {
	fake.notifyMutex.Lock()
	defer fake.notifyMutex.Unlock()
	fake.NotifyStub = nil
	fake.notifyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNotifierAllInOne) NotifyReturnsOnCall(i int, result1 error) {
	fake.notifyMutex.Lock()
	defer fake.notifyMutex.Unlock()
	fake.NotifyStub = nil
	if fake.notifyReturnsOnCall == nil {
		fake.notifyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.notifyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNotifierAllInOne) PreCheck(arg1 models.Incident) error {
	fake.preCheckMutex.Lock()
	ret, specificReturn := fake.preCheckReturnsOnCall[len(fake.preCheckArgsForCall)]
	fake.preCheckArgsForCall = append(fake.preCheckArgsForCall, struct {
		arg1 models.Incident
	}{arg1})
	fake.recordInvocation("PreCheck", []interface{}{arg1})
	fake.preCheckMutex.Unlock()
	if fake.PreCheckStub != nil {
		return fake.PreCheckStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.preCheckReturns
	return fakeReturns.result1
}

func (fake *FakeNotifierAllInOne) PreCheckCallCount() int {
	fake.preCheckMutex.RLock()
	defer fake.preCheckMutex.RUnlock()
	return len(fake.preCheckArgsForCall)
}

func (fake *FakeNotifierAllInOne) PreCheckCalls(stub func(models.Incident) error) {
	fake.preCheckMutex.Lock()
	defer fake.preCheckMutex.Unlock()
	fake.PreCheckStub = stub
}

func (fake *FakeNotifierAllInOne) PreCheckArgsForCall(i int) models.Incident {
	fake.preCheckMutex.RLock()
	defer fake.preCheckMutex.RUnlock()
	argsForCall := fake.preCheckArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeNotifierAllInOne) PreCheckReturns(result1 error) {
	fake.preCheckMutex.Lock()
	defer fake.preCheckMutex.Unlock()
	fake.PreCheckStub = nil
	fake.preCheckReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNotifierAllInOne) PreCheckReturnsOnCall(i int, result1 error) {
	fake.preCheckMutex.Lock()
	defer fake.preCheckMutex.Unlock()
	fake.PreCheckStub = nil
	if fake.preCheckReturnsOnCall == nil {
		fake.preCheckReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.preCheckReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNotifierAllInOne) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.creatorMutex.RLock()
	defer fake.creatorMutex.RUnlock()
	fake.idMutex.RLock()
	defer fake.idMutex.RUnlock()
	fake.metadataFieldsMutex.RLock()
	defer fake.metadataFieldsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	fake.preCheckMutex.RLock()
	defer fake.preCheckMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNotifierAllInOne) recordInvocation(key string, args []interface{}) {
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

var _ notifiers.NotifierAllInOne = new(FakeNotifierAllInOne)
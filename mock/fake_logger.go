// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"drivinglicence"
	"sync"
)

type FakeLogger struct {
	LogStuffStub        func(string)
	logStuffMutex       sync.RWMutex
	logStuffArgsForCall []struct {
		arg1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLogger) LogStuff(arg1 string) {
	fake.logStuffMutex.Lock()
	fake.logStuffArgsForCall = append(fake.logStuffArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LogStuffStub
	fake.recordInvocation("LogStuff", []interface{}{arg1})
	fake.logStuffMutex.Unlock()
	if stub != nil {
		fake.LogStuffStub(arg1)
	}
}

func (fake *FakeLogger) LogStuffCallCount() int {
	fake.logStuffMutex.RLock()
	defer fake.logStuffMutex.RUnlock()
	return len(fake.logStuffArgsForCall)
}

func (fake *FakeLogger) LogStuffCalls(stub func(string)) {
	fake.logStuffMutex.Lock()
	defer fake.logStuffMutex.Unlock()
	fake.LogStuffStub = stub
}

func (fake *FakeLogger) LogStuffArgsForCall(i int) string {
	fake.logStuffMutex.RLock()
	defer fake.logStuffMutex.RUnlock()
	argsForCall := fake.logStuffArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeLogger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.logStuffMutex.RLock()
	defer fake.logStuffMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLogger) recordInvocation(key string, args []interface{}) {
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

var _ drivinglicence.Logger = new(FakeLogger)
// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/IBM/idemix/bccsp/schemes/dlog/handlers"
)

type Ecp struct {
	BytesStub        func() []byte
	bytesMutex       sync.RWMutex
	bytesArgsForCall []struct {
	}
	bytesReturns struct {
		result1 []byte
	}
	bytesReturnsOnCall map[int]struct {
		result1 []byte
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Ecp) Bytes() []byte {
	fake.bytesMutex.Lock()
	ret, specificReturn := fake.bytesReturnsOnCall[len(fake.bytesArgsForCall)]
	fake.bytesArgsForCall = append(fake.bytesArgsForCall, struct {
	}{})
	fake.recordInvocation("Bytes", []interface{}{})
	fake.bytesMutex.Unlock()
	if fake.BytesStub != nil {
		return fake.BytesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.bytesReturns
	return fakeReturns.result1
}

func (fake *Ecp) BytesCallCount() int {
	fake.bytesMutex.RLock()
	defer fake.bytesMutex.RUnlock()
	return len(fake.bytesArgsForCall)
}

func (fake *Ecp) BytesCalls(stub func() []byte) {
	fake.bytesMutex.Lock()
	defer fake.bytesMutex.Unlock()
	fake.BytesStub = stub
}

func (fake *Ecp) BytesReturns(result1 []byte) {
	fake.bytesMutex.Lock()
	defer fake.bytesMutex.Unlock()
	fake.BytesStub = nil
	fake.bytesReturns = struct {
		result1 []byte
	}{result1}
}

func (fake *Ecp) BytesReturnsOnCall(i int, result1 []byte) {
	fake.bytesMutex.Lock()
	defer fake.bytesMutex.Unlock()
	fake.BytesStub = nil
	if fake.bytesReturnsOnCall == nil {
		fake.bytesReturnsOnCall = make(map[int]struct {
			result1 []byte
		})
	}
	fake.bytesReturnsOnCall[i] = struct {
		result1 []byte
	}{result1}
}

func (fake *Ecp) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bytesMutex.RLock()
	defer fake.bytesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Ecp) recordInvocation(key string, args []interface{}) {
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

var _ handlers.Ecp = new(Ecp)

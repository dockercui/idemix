// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/IBM/idemix/bccsp/types"
)

type Issuer struct {
	BasesStub        func(types.IssuerPublicKey, types.CommitmentBasesRequest, int, int, int) (map[types.CommitmentType]interface{}, error)
	basesMutex       sync.RWMutex
	basesArgsForCall []struct {
		arg1 types.IssuerPublicKey
		arg2 types.CommitmentBasesRequest
		arg3 int
		arg4 int
		arg5 int
	}
	basesReturns struct {
		result1 map[types.CommitmentType]interface{}
		result2 error
	}
	basesReturnsOnCall map[int]struct {
		result1 map[types.CommitmentType]interface{}
		result2 error
	}
	NewKeyStub        func([]string) (types.IssuerSecretKey, error)
	newKeyMutex       sync.RWMutex
	newKeyArgsForCall []struct {
		arg1 []string
	}
	newKeyReturns struct {
		result1 types.IssuerSecretKey
		result2 error
	}
	newKeyReturnsOnCall map[int]struct {
		result1 types.IssuerSecretKey
		result2 error
	}
	NewKeyFromBytesStub        func([]byte, []string) (types.IssuerSecretKey, error)
	newKeyFromBytesMutex       sync.RWMutex
	newKeyFromBytesArgsForCall []struct {
		arg1 []byte
		arg2 []string
	}
	newKeyFromBytesReturns struct {
		result1 types.IssuerSecretKey
		result2 error
	}
	newKeyFromBytesReturnsOnCall map[int]struct {
		result1 types.IssuerSecretKey
		result2 error
	}
	NewPublicKeyFromBytesStub        func([]byte, []string) (types.IssuerPublicKey, error)
	newPublicKeyFromBytesMutex       sync.RWMutex
	newPublicKeyFromBytesArgsForCall []struct {
		arg1 []byte
		arg2 []string
	}
	newPublicKeyFromBytesReturns struct {
		result1 types.IssuerPublicKey
		result2 error
	}
	newPublicKeyFromBytesReturnsOnCall map[int]struct {
		result1 types.IssuerPublicKey
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Issuer) Bases(arg1 types.IssuerPublicKey, arg2 types.CommitmentBasesRequest, arg3 int, arg4 int, arg5 int) (map[types.CommitmentType]interface{}, error) {
	fake.basesMutex.Lock()
	ret, specificReturn := fake.basesReturnsOnCall[len(fake.basesArgsForCall)]
	fake.basesArgsForCall = append(fake.basesArgsForCall, struct {
		arg1 types.IssuerPublicKey
		arg2 types.CommitmentBasesRequest
		arg3 int
		arg4 int
		arg5 int
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.BasesStub
	fakeReturns := fake.basesReturns
	fake.recordInvocation("Bases", []interface{}{arg1, arg2, arg3, arg4, arg5})
	fake.basesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Issuer) BasesCallCount() int {
	fake.basesMutex.RLock()
	defer fake.basesMutex.RUnlock()
	return len(fake.basesArgsForCall)
}

func (fake *Issuer) BasesCalls(stub func(types.IssuerPublicKey, types.CommitmentBasesRequest, int, int, int) (map[types.CommitmentType]interface{}, error)) {
	fake.basesMutex.Lock()
	defer fake.basesMutex.Unlock()
	fake.BasesStub = stub
}

func (fake *Issuer) BasesArgsForCall(i int) (types.IssuerPublicKey, types.CommitmentBasesRequest, int, int, int) {
	fake.basesMutex.RLock()
	defer fake.basesMutex.RUnlock()
	argsForCall := fake.basesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *Issuer) BasesReturns(result1 map[types.CommitmentType]interface{}, result2 error) {
	fake.basesMutex.Lock()
	defer fake.basesMutex.Unlock()
	fake.BasesStub = nil
	fake.basesReturns = struct {
		result1 map[types.CommitmentType]interface{}
		result2 error
	}{result1, result2}
}

func (fake *Issuer) BasesReturnsOnCall(i int, result1 map[types.CommitmentType]interface{}, result2 error) {
	fake.basesMutex.Lock()
	defer fake.basesMutex.Unlock()
	fake.BasesStub = nil
	if fake.basesReturnsOnCall == nil {
		fake.basesReturnsOnCall = make(map[int]struct {
			result1 map[types.CommitmentType]interface{}
			result2 error
		})
	}
	fake.basesReturnsOnCall[i] = struct {
		result1 map[types.CommitmentType]interface{}
		result2 error
	}{result1, result2}
}

func (fake *Issuer) NewKey(arg1 []string) (types.IssuerSecretKey, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.newKeyMutex.Lock()
	ret, specificReturn := fake.newKeyReturnsOnCall[len(fake.newKeyArgsForCall)]
	fake.newKeyArgsForCall = append(fake.newKeyArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.NewKeyStub
	fakeReturns := fake.newKeyReturns
	fake.recordInvocation("NewKey", []interface{}{arg1Copy})
	fake.newKeyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Issuer) NewKeyCallCount() int {
	fake.newKeyMutex.RLock()
	defer fake.newKeyMutex.RUnlock()
	return len(fake.newKeyArgsForCall)
}

func (fake *Issuer) NewKeyCalls(stub func([]string) (types.IssuerSecretKey, error)) {
	fake.newKeyMutex.Lock()
	defer fake.newKeyMutex.Unlock()
	fake.NewKeyStub = stub
}

func (fake *Issuer) NewKeyArgsForCall(i int) []string {
	fake.newKeyMutex.RLock()
	defer fake.newKeyMutex.RUnlock()
	argsForCall := fake.newKeyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *Issuer) NewKeyReturns(result1 types.IssuerSecretKey, result2 error) {
	fake.newKeyMutex.Lock()
	defer fake.newKeyMutex.Unlock()
	fake.NewKeyStub = nil
	fake.newKeyReturns = struct {
		result1 types.IssuerSecretKey
		result2 error
	}{result1, result2}
}

func (fake *Issuer) NewKeyReturnsOnCall(i int, result1 types.IssuerSecretKey, result2 error) {
	fake.newKeyMutex.Lock()
	defer fake.newKeyMutex.Unlock()
	fake.NewKeyStub = nil
	if fake.newKeyReturnsOnCall == nil {
		fake.newKeyReturnsOnCall = make(map[int]struct {
			result1 types.IssuerSecretKey
			result2 error
		})
	}
	fake.newKeyReturnsOnCall[i] = struct {
		result1 types.IssuerSecretKey
		result2 error
	}{result1, result2}
}

func (fake *Issuer) NewKeyFromBytes(arg1 []byte, arg2 []string) (types.IssuerSecretKey, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.newKeyFromBytesMutex.Lock()
	ret, specificReturn := fake.newKeyFromBytesReturnsOnCall[len(fake.newKeyFromBytesArgsForCall)]
	fake.newKeyFromBytesArgsForCall = append(fake.newKeyFromBytesArgsForCall, struct {
		arg1 []byte
		arg2 []string
	}{arg1Copy, arg2Copy})
	stub := fake.NewKeyFromBytesStub
	fakeReturns := fake.newKeyFromBytesReturns
	fake.recordInvocation("NewKeyFromBytes", []interface{}{arg1Copy, arg2Copy})
	fake.newKeyFromBytesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Issuer) NewKeyFromBytesCallCount() int {
	fake.newKeyFromBytesMutex.RLock()
	defer fake.newKeyFromBytesMutex.RUnlock()
	return len(fake.newKeyFromBytesArgsForCall)
}

func (fake *Issuer) NewKeyFromBytesCalls(stub func([]byte, []string) (types.IssuerSecretKey, error)) {
	fake.newKeyFromBytesMutex.Lock()
	defer fake.newKeyFromBytesMutex.Unlock()
	fake.NewKeyFromBytesStub = stub
}

func (fake *Issuer) NewKeyFromBytesArgsForCall(i int) ([]byte, []string) {
	fake.newKeyFromBytesMutex.RLock()
	defer fake.newKeyFromBytesMutex.RUnlock()
	argsForCall := fake.newKeyFromBytesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Issuer) NewKeyFromBytesReturns(result1 types.IssuerSecretKey, result2 error) {
	fake.newKeyFromBytesMutex.Lock()
	defer fake.newKeyFromBytesMutex.Unlock()
	fake.NewKeyFromBytesStub = nil
	fake.newKeyFromBytesReturns = struct {
		result1 types.IssuerSecretKey
		result2 error
	}{result1, result2}
}

func (fake *Issuer) NewKeyFromBytesReturnsOnCall(i int, result1 types.IssuerSecretKey, result2 error) {
	fake.newKeyFromBytesMutex.Lock()
	defer fake.newKeyFromBytesMutex.Unlock()
	fake.NewKeyFromBytesStub = nil
	if fake.newKeyFromBytesReturnsOnCall == nil {
		fake.newKeyFromBytesReturnsOnCall = make(map[int]struct {
			result1 types.IssuerSecretKey
			result2 error
		})
	}
	fake.newKeyFromBytesReturnsOnCall[i] = struct {
		result1 types.IssuerSecretKey
		result2 error
	}{result1, result2}
}

func (fake *Issuer) NewPublicKeyFromBytes(arg1 []byte, arg2 []string) (types.IssuerPublicKey, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.newPublicKeyFromBytesMutex.Lock()
	ret, specificReturn := fake.newPublicKeyFromBytesReturnsOnCall[len(fake.newPublicKeyFromBytesArgsForCall)]
	fake.newPublicKeyFromBytesArgsForCall = append(fake.newPublicKeyFromBytesArgsForCall, struct {
		arg1 []byte
		arg2 []string
	}{arg1Copy, arg2Copy})
	stub := fake.NewPublicKeyFromBytesStub
	fakeReturns := fake.newPublicKeyFromBytesReturns
	fake.recordInvocation("NewPublicKeyFromBytes", []interface{}{arg1Copy, arg2Copy})
	fake.newPublicKeyFromBytesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Issuer) NewPublicKeyFromBytesCallCount() int {
	fake.newPublicKeyFromBytesMutex.RLock()
	defer fake.newPublicKeyFromBytesMutex.RUnlock()
	return len(fake.newPublicKeyFromBytesArgsForCall)
}

func (fake *Issuer) NewPublicKeyFromBytesCalls(stub func([]byte, []string) (types.IssuerPublicKey, error)) {
	fake.newPublicKeyFromBytesMutex.Lock()
	defer fake.newPublicKeyFromBytesMutex.Unlock()
	fake.NewPublicKeyFromBytesStub = stub
}

func (fake *Issuer) NewPublicKeyFromBytesArgsForCall(i int) ([]byte, []string) {
	fake.newPublicKeyFromBytesMutex.RLock()
	defer fake.newPublicKeyFromBytesMutex.RUnlock()
	argsForCall := fake.newPublicKeyFromBytesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *Issuer) NewPublicKeyFromBytesReturns(result1 types.IssuerPublicKey, result2 error) {
	fake.newPublicKeyFromBytesMutex.Lock()
	defer fake.newPublicKeyFromBytesMutex.Unlock()
	fake.NewPublicKeyFromBytesStub = nil
	fake.newPublicKeyFromBytesReturns = struct {
		result1 types.IssuerPublicKey
		result2 error
	}{result1, result2}
}

func (fake *Issuer) NewPublicKeyFromBytesReturnsOnCall(i int, result1 types.IssuerPublicKey, result2 error) {
	fake.newPublicKeyFromBytesMutex.Lock()
	defer fake.newPublicKeyFromBytesMutex.Unlock()
	fake.NewPublicKeyFromBytesStub = nil
	if fake.newPublicKeyFromBytesReturnsOnCall == nil {
		fake.newPublicKeyFromBytesReturnsOnCall = make(map[int]struct {
			result1 types.IssuerPublicKey
			result2 error
		})
	}
	fake.newPublicKeyFromBytesReturnsOnCall[i] = struct {
		result1 types.IssuerPublicKey
		result2 error
	}{result1, result2}
}

func (fake *Issuer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.basesMutex.RLock()
	defer fake.basesMutex.RUnlock()
	fake.newKeyMutex.RLock()
	defer fake.newKeyMutex.RUnlock()
	fake.newKeyFromBytesMutex.RLock()
	defer fake.newKeyFromBytesMutex.RUnlock()
	fake.newPublicKeyFromBytesMutex.RLock()
	defer fake.newPublicKeyFromBytesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Issuer) recordInvocation(key string, args []interface{}) {
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

var _ types.Issuer = new(Issuer)

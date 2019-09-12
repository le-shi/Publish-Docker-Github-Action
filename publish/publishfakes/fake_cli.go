// Code generated by counterfeiter. DO NOT EDIT.
package publishfakes

import (
	"context"
	"io"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/elgohr/Publish-Docker-Github-Action/publish"
)

type FakeCli struct {
	ImageBuildStub        func(context.Context, io.Reader, types.ImageBuildOptions) (types.ImageBuildResponse, error)
	imageBuildMutex       sync.RWMutex
	imageBuildArgsForCall []struct {
		arg1 context.Context
		arg2 io.Reader
		arg3 types.ImageBuildOptions
	}
	imageBuildReturns struct {
		result1 types.ImageBuildResponse
		result2 error
	}
	imageBuildReturnsOnCall map[int]struct {
		result1 types.ImageBuildResponse
		result2 error
	}
	ImagePullStub        func(context.Context, string, types.ImagePullOptions) (io.ReadCloser, error)
	imagePullMutex       sync.RWMutex
	imagePullArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 types.ImagePullOptions
	}
	imagePullReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	imagePullReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	ImagePushStub        func(context.Context, string, types.ImagePushOptions) (io.ReadCloser, error)
	imagePushMutex       sync.RWMutex
	imagePushArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 types.ImagePushOptions
	}
	imagePushReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	imagePushReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	RegistryLoginStub        func(context.Context, types.AuthConfig) (registry.AuthenticateOKBody, error)
	registryLoginMutex       sync.RWMutex
	registryLoginArgsForCall []struct {
		arg1 context.Context
		arg2 types.AuthConfig
	}
	registryLoginReturns struct {
		result1 registry.AuthenticateOKBody
		result2 error
	}
	registryLoginReturnsOnCall map[int]struct {
		result1 registry.AuthenticateOKBody
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCli) ImageBuild(arg1 context.Context, arg2 io.Reader, arg3 types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	fake.imageBuildMutex.Lock()
	ret, specificReturn := fake.imageBuildReturnsOnCall[len(fake.imageBuildArgsForCall)]
	fake.imageBuildArgsForCall = append(fake.imageBuildArgsForCall, struct {
		arg1 context.Context
		arg2 io.Reader
		arg3 types.ImageBuildOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("ImageBuild", []interface{}{arg1, arg2, arg3})
	fake.imageBuildMutex.Unlock()
	if fake.ImageBuildStub != nil {
		return fake.ImageBuildStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.imageBuildReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCli) ImageBuildCallCount() int {
	fake.imageBuildMutex.RLock()
	defer fake.imageBuildMutex.RUnlock()
	return len(fake.imageBuildArgsForCall)
}

func (fake *FakeCli) ImageBuildCalls(stub func(context.Context, io.Reader, types.ImageBuildOptions) (types.ImageBuildResponse, error)) {
	fake.imageBuildMutex.Lock()
	defer fake.imageBuildMutex.Unlock()
	fake.ImageBuildStub = stub
}

func (fake *FakeCli) ImageBuildArgsForCall(i int) (context.Context, io.Reader, types.ImageBuildOptions) {
	fake.imageBuildMutex.RLock()
	defer fake.imageBuildMutex.RUnlock()
	argsForCall := fake.imageBuildArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCli) ImageBuildReturns(result1 types.ImageBuildResponse, result2 error) {
	fake.imageBuildMutex.Lock()
	defer fake.imageBuildMutex.Unlock()
	fake.ImageBuildStub = nil
	fake.imageBuildReturns = struct {
		result1 types.ImageBuildResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCli) ImageBuildReturnsOnCall(i int, result1 types.ImageBuildResponse, result2 error) {
	fake.imageBuildMutex.Lock()
	defer fake.imageBuildMutex.Unlock()
	fake.ImageBuildStub = nil
	if fake.imageBuildReturnsOnCall == nil {
		fake.imageBuildReturnsOnCall = make(map[int]struct {
			result1 types.ImageBuildResponse
			result2 error
		})
	}
	fake.imageBuildReturnsOnCall[i] = struct {
		result1 types.ImageBuildResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCli) ImagePull(arg1 context.Context, arg2 string, arg3 types.ImagePullOptions) (io.ReadCloser, error) {
	fake.imagePullMutex.Lock()
	ret, specificReturn := fake.imagePullReturnsOnCall[len(fake.imagePullArgsForCall)]
	fake.imagePullArgsForCall = append(fake.imagePullArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 types.ImagePullOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("ImagePull", []interface{}{arg1, arg2, arg3})
	fake.imagePullMutex.Unlock()
	if fake.ImagePullStub != nil {
		return fake.ImagePullStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.imagePullReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCli) ImagePullCallCount() int {
	fake.imagePullMutex.RLock()
	defer fake.imagePullMutex.RUnlock()
	return len(fake.imagePullArgsForCall)
}

func (fake *FakeCli) ImagePullCalls(stub func(context.Context, string, types.ImagePullOptions) (io.ReadCloser, error)) {
	fake.imagePullMutex.Lock()
	defer fake.imagePullMutex.Unlock()
	fake.ImagePullStub = stub
}

func (fake *FakeCli) ImagePullArgsForCall(i int) (context.Context, string, types.ImagePullOptions) {
	fake.imagePullMutex.RLock()
	defer fake.imagePullMutex.RUnlock()
	argsForCall := fake.imagePullArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCli) ImagePullReturns(result1 io.ReadCloser, result2 error) {
	fake.imagePullMutex.Lock()
	defer fake.imagePullMutex.Unlock()
	fake.ImagePullStub = nil
	fake.imagePullReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCli) ImagePullReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.imagePullMutex.Lock()
	defer fake.imagePullMutex.Unlock()
	fake.ImagePullStub = nil
	if fake.imagePullReturnsOnCall == nil {
		fake.imagePullReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.imagePullReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCli) ImagePush(arg1 context.Context, arg2 string, arg3 types.ImagePushOptions) (io.ReadCloser, error) {
	fake.imagePushMutex.Lock()
	ret, specificReturn := fake.imagePushReturnsOnCall[len(fake.imagePushArgsForCall)]
	fake.imagePushArgsForCall = append(fake.imagePushArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 types.ImagePushOptions
	}{arg1, arg2, arg3})
	fake.recordInvocation("ImagePush", []interface{}{arg1, arg2, arg3})
	fake.imagePushMutex.Unlock()
	if fake.ImagePushStub != nil {
		return fake.ImagePushStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.imagePushReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCli) ImagePushCallCount() int {
	fake.imagePushMutex.RLock()
	defer fake.imagePushMutex.RUnlock()
	return len(fake.imagePushArgsForCall)
}

func (fake *FakeCli) ImagePushCalls(stub func(context.Context, string, types.ImagePushOptions) (io.ReadCloser, error)) {
	fake.imagePushMutex.Lock()
	defer fake.imagePushMutex.Unlock()
	fake.ImagePushStub = stub
}

func (fake *FakeCli) ImagePushArgsForCall(i int) (context.Context, string, types.ImagePushOptions) {
	fake.imagePushMutex.RLock()
	defer fake.imagePushMutex.RUnlock()
	argsForCall := fake.imagePushArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCli) ImagePushReturns(result1 io.ReadCloser, result2 error) {
	fake.imagePushMutex.Lock()
	defer fake.imagePushMutex.Unlock()
	fake.ImagePushStub = nil
	fake.imagePushReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCli) ImagePushReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.imagePushMutex.Lock()
	defer fake.imagePushMutex.Unlock()
	fake.ImagePushStub = nil
	if fake.imagePushReturnsOnCall == nil {
		fake.imagePushReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.imagePushReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCli) RegistryLogin(arg1 context.Context, arg2 types.AuthConfig) (registry.AuthenticateOKBody, error) {
	fake.registryLoginMutex.Lock()
	ret, specificReturn := fake.registryLoginReturnsOnCall[len(fake.registryLoginArgsForCall)]
	fake.registryLoginArgsForCall = append(fake.registryLoginArgsForCall, struct {
		arg1 context.Context
		arg2 types.AuthConfig
	}{arg1, arg2})
	fake.recordInvocation("RegistryLogin", []interface{}{arg1, arg2})
	fake.registryLoginMutex.Unlock()
	if fake.RegistryLoginStub != nil {
		return fake.RegistryLoginStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.registryLoginReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCli) RegistryLoginCallCount() int {
	fake.registryLoginMutex.RLock()
	defer fake.registryLoginMutex.RUnlock()
	return len(fake.registryLoginArgsForCall)
}

func (fake *FakeCli) RegistryLoginCalls(stub func(context.Context, types.AuthConfig) (registry.AuthenticateOKBody, error)) {
	fake.registryLoginMutex.Lock()
	defer fake.registryLoginMutex.Unlock()
	fake.RegistryLoginStub = stub
}

func (fake *FakeCli) RegistryLoginArgsForCall(i int) (context.Context, types.AuthConfig) {
	fake.registryLoginMutex.RLock()
	defer fake.registryLoginMutex.RUnlock()
	argsForCall := fake.registryLoginArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCli) RegistryLoginReturns(result1 registry.AuthenticateOKBody, result2 error) {
	fake.registryLoginMutex.Lock()
	defer fake.registryLoginMutex.Unlock()
	fake.RegistryLoginStub = nil
	fake.registryLoginReturns = struct {
		result1 registry.AuthenticateOKBody
		result2 error
	}{result1, result2}
}

func (fake *FakeCli) RegistryLoginReturnsOnCall(i int, result1 registry.AuthenticateOKBody, result2 error) {
	fake.registryLoginMutex.Lock()
	defer fake.registryLoginMutex.Unlock()
	fake.RegistryLoginStub = nil
	if fake.registryLoginReturnsOnCall == nil {
		fake.registryLoginReturnsOnCall = make(map[int]struct {
			result1 registry.AuthenticateOKBody
			result2 error
		})
	}
	fake.registryLoginReturnsOnCall[i] = struct {
		result1 registry.AuthenticateOKBody
		result2 error
	}{result1, result2}
}

func (fake *FakeCli) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.imageBuildMutex.RLock()
	defer fake.imageBuildMutex.RUnlock()
	fake.imagePullMutex.RLock()
	defer fake.imagePullMutex.RUnlock()
	fake.imagePushMutex.RLock()
	defer fake.imagePushMutex.RUnlock()
	fake.registryLoginMutex.RLock()
	defer fake.registryLoginMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCli) recordInvocation(key string, args []interface{}) {
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

var _ publish.Cli = new(FakeCli)

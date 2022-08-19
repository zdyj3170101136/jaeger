// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	storage_v1 "github.com/zdyj3170101136/jaeger/proto-gen/storage_v1"
)

// DependenciesReaderPluginServer is an autogenerated mock type for the DependenciesReaderPluginServer type
type DependenciesReaderPluginServer struct {
	mock.Mock
}

// GetDependencies provides a mock function with given fields: _a0, _a1
func (_m *DependenciesReaderPluginServer) GetDependencies(_a0 context.Context, _a1 *storage_v1.GetDependenciesRequest) (*storage_v1.GetDependenciesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *storage_v1.GetDependenciesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *storage_v1.GetDependenciesRequest) *storage_v1.GetDependenciesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage_v1.GetDependenciesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *storage_v1.GetDependenciesRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	storage_v1 "github.com/zdyj3170101136/jaeger/proto-gen/storage_v1"
)

// ArchiveSpanReaderPluginServer is an autogenerated mock type for the ArchiveSpanReaderPluginServer type
type ArchiveSpanReaderPluginServer struct {
	mock.Mock
}

// GetArchiveTrace provides a mock function with given fields: _a0, _a1
func (_m *ArchiveSpanReaderPluginServer) GetArchiveTrace(_a0 *storage_v1.GetTraceRequest, _a1 storage_v1.ArchiveSpanReaderPlugin_GetArchiveTraceServer) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*storage_v1.GetTraceRequest, storage_v1.ArchiveSpanReaderPlugin_GetArchiveTraceServer) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

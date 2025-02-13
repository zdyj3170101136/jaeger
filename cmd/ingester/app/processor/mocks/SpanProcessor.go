// Code generated by mockery v1.0.0

// Copyright (c) 2018 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mocks

import mock "github.com/stretchr/testify/mock"
import processor "github.com/zdyj3170101136/jaeger/cmd/ingester/app/processor"

// SpanProcessor is an autogenerated mock type for the SpanProcessor type
type SpanProcessor struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *SpanProcessor) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Process provides a mock function with given fields: input
func (_m *SpanProcessor) Process(input processor.Message) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(processor.Message) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

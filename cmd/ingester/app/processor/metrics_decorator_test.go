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

package processor_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zdyj3170101136/jaeger/cmd/ingester/app/processor"
	"github.com/zdyj3170101136/jaeger/cmd/ingester/app/processor/mocks"
	"github.com/zdyj3170101136/jaeger/internal/metricstest"
)

type fakeMsg struct{}

func (fakeMsg) Value() []byte {
	return nil
}

func TestProcess(t *testing.T) {
	p := &mocks.SpanProcessor{}
	msg := fakeMsg{}
	p.On("Process", msg).Return(nil)
	m := metricstest.NewFactory(0)
	proc := processor.NewDecoratedProcessor(m, p)

	proc.Process(msg)
	p.AssertExpectations(t)
	_, g := m.Snapshot()
	assert.Contains(t, g, "span-processor.latency.P90")
}

func TestProcessErr(t *testing.T) {
	p := &mocks.SpanProcessor{}
	msg := fakeMsg{}
	p.On("Process", msg).Return(errors.New("err"))
	m := metricstest.NewFactory(0)
	proc := processor.NewDecoratedProcessor(m, p)

	proc.Process(msg)
	p.AssertExpectations(t)
	c, g := m.Snapshot()
	assert.Contains(t, g, "span-processor.latency.P90")
	assert.Equal(t, int64(1), c["span-processor.errors"])
}

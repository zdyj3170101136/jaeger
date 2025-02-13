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

package static

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	ss "github.com/zdyj3170101136/jaeger/cmd/collector/app/sampling/strategystore"
	"github.com/zdyj3170101136/jaeger/pkg/config"
	"github.com/zdyj3170101136/jaeger/pkg/metrics"
	"github.com/zdyj3170101136/jaeger/plugin"
)

var (
	_ ss.Factory          = new(Factory)
	_ plugin.Configurable = new(Factory)
)

func TestFactory(t *testing.T) {
	f := NewFactory()
	v, command := config.Viperize(f.AddFlags)
	command.ParseFlags([]string{"--sampling.strategies-file=fixtures/strategies.json"})
	f.InitFromViper(v, zap.NewNop())

	assert.NoError(t, f.Initialize(metrics.NullFactory, nil, zap.NewNop()))
	_, _, err := f.CreateStrategyStore()
	assert.NoError(t, err)
}

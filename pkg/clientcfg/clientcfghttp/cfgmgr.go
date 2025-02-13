// Copyright (c) 2020 The Jaeger Authors.
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

package clientcfghttp

import (
	"context"
	"errors"

	"github.com/zdyj3170101136/jaeger/cmd/collector/app/sampling/strategystore"
	"github.com/zdyj3170101136/jaeger/thrift-gen/baggage"
	"github.com/zdyj3170101136/jaeger/thrift-gen/sampling"
)

// ConfigManager implements ClientConfigManager.
type ConfigManager struct {
	SamplingStrategyStore strategystore.StrategyStore
	BaggageManager        baggage.BaggageRestrictionManager
}

// GetSamplingStrategy implements ClientConfigManager.GetSamplingStrategy.
func (c *ConfigManager) GetSamplingStrategy(ctx context.Context, serviceName string) (*sampling.SamplingStrategyResponse, error) {
	return c.SamplingStrategyStore.GetSamplingStrategy(ctx, serviceName)
}

// GetBaggageRestrictions implements ClientConfigManager.GetBaggageRestrictions.
func (c *ConfigManager) GetBaggageRestrictions(ctx context.Context, serviceName string) ([]*baggage.BaggageRestriction, error) {
	if c.BaggageManager == nil {
		return nil, errors.New("baggage restrictions not implemented")
	}
	return c.BaggageManager.GetBaggageRestrictions(ctx, serviceName)
}

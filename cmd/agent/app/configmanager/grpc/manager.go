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

package grpc

import (
	"context"
	"errors"

	"google.golang.org/grpc"

	"github.com/zdyj3170101136/jaeger/model/converter/thrift/jaeger"
	"github.com/zdyj3170101136/jaeger/proto-gen/api_v2"
	"github.com/zdyj3170101136/jaeger/thrift-gen/baggage"
	"github.com/zdyj3170101136/jaeger/thrift-gen/sampling"
)

// SamplingManager returns sampling decisions from collector over gRPC.
type SamplingManager struct {
	client api_v2.SamplingManagerClient
}

// NewConfigManager creates gRPC sampling manager.
func NewConfigManager(conn *grpc.ClientConn) *SamplingManager {
	return &SamplingManager{
		client: api_v2.NewSamplingManagerClient(conn),
	}
}

// GetSamplingStrategy returns sampling strategies from collector.
func (s *SamplingManager) GetSamplingStrategy(ctx context.Context, serviceName string) (*sampling.SamplingStrategyResponse, error) {
	r, err := s.client.GetSamplingStrategy(ctx, &api_v2.SamplingStrategyParameters{ServiceName: serviceName})
	if err != nil {
		return nil, err
	}
	return jaeger.ConvertSamplingResponseFromDomain(r)
}

// GetBaggageRestrictions returns baggage restrictions from collector.
func (s *SamplingManager) GetBaggageRestrictions(_ context.Context, _ string) ([]*baggage.BaggageRestriction, error) {
	return nil, errors.New("baggage not implemented")
}

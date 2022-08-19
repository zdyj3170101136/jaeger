// Copyright (c) 2019 The Jaeger Authors.
// Copyright (c) 2017 Uber Technologies, Inc.
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

package memory

import (
	"flag"

	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/zdyj3170101136/jaeger/pkg/distributedlock"
	"github.com/zdyj3170101136/jaeger/pkg/metrics"
	"github.com/zdyj3170101136/jaeger/plugin"
	"github.com/zdyj3170101136/jaeger/storage/dependencystore"
	"github.com/zdyj3170101136/jaeger/storage/samplingstore"
	"github.com/zdyj3170101136/jaeger/storage/spanstore"
)

var _ plugin.Configurable = (*Factory)(nil)

// Factory implements storage.Factory and creates storage components backed by memory store.
type Factory struct {
	options        Options
	metricsFactory metrics.Factory
	logger         *zap.Logger
	store          *Store
}

// NewFactory creates a new Factory.
func NewFactory() *Factory {
	return &Factory{}
}

// AddFlags implements plugin.Configurable
func (f *Factory) AddFlags(flagSet *flag.FlagSet) {
	AddFlags(flagSet)
}

// InitFromViper implements plugin.Configurable
func (f *Factory) InitFromViper(v *viper.Viper, logger *zap.Logger) {
	f.options.InitFromViper(v)
}

// InitFromOptions initializes factory from the supplied options
func (f *Factory) InitFromOptions(opts Options) {
	f.options = opts
}

// Initialize implements storage.Factory
func (f *Factory) Initialize(metricsFactory metrics.Factory, logger *zap.Logger) error {
	f.metricsFactory, f.logger = metricsFactory, logger
	f.store = WithConfiguration(f.options.Configuration)
	logger.Info("Memory storage initialized", zap.Any("configuration", f.store.defaultConfig))
	f.publishOpts()

	return nil
}

// CreateSpanReader implements storage.Factory
func (f *Factory) CreateSpanReader() (spanstore.Reader, error) {
	return f.store, nil
}

// CreateSpanWriter implements storage.Factory
func (f *Factory) CreateSpanWriter() (spanstore.Writer, error) {
	return f.store, nil
}

// CreateDependencyReader implements storage.Factory
func (f *Factory) CreateDependencyReader() (dependencystore.Reader, error) {
	return f.store, nil
}

// CreateSamplingStore implements storage.SamplingStoreFactory
func (f *Factory) CreateSamplingStore(maxBuckets int) (samplingstore.Store, error) {
	return NewSamplingStore(maxBuckets), nil
}

// CreateLock implements storage.SamplingStoreFactory
func (f *Factory) CreateLock() (distributedlock.Lock, error) {
	return &lock{}, nil
}

func (f *Factory) publishOpts() {
	internalFactory := f.metricsFactory.Namespace(metrics.NSOptions{Name: "internal"})
	internalFactory.Gauge(metrics.Options{Name: limit}).
		Update(int64(f.options.Configuration.MaxTraces))
}

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

package env

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/zdyj3170101136/jaeger/plugin/metrics"
	"github.com/zdyj3170101136/jaeger/plugin/sampling/strategystore"
	"github.com/zdyj3170101136/jaeger/plugin/storage"
)

const (
	longTemplate = `
All command line options can be provided via environment variables by converting
their names to upper case and replacing punctuation with underscores. For example:

command line option                 environment variable
------------------------------------------------------------------
--cassandra.connections-per-host    CASSANDRA_CONNECTIONS_PER_HOST
--metrics-backend                   METRICS_BACKEND

The following configuration options are only available via environment variables:
%s
`
	storageTypeDescription = `The type of backend [%s] used for trace storage.
Multiple backends can be specified as comma-separated list, e.g. "cassandra,elasticsearch"
(currently only for writing spans). Note that "kafka" is only valid in jaeger-collector;
it is not a replacement for a proper storage backend, and only used as a buffer for spans
when Jaeger is deployed in the collector+ingester configuration.
`

	samplingTypeDescription = `The method [%s] used for determining the sampling rates served
to clients configured with remote sampling enabled. "file" uses a periodically reloaded file and
"adaptive" dynamically adjusts sampling rates based on current traffic.
`

	samplingStorageTypeDescription = `The type of backend [%s] used for adaptive sampling storage
when adaptive sampling is enabled via %s.
`

	metricsStorageTypeDescription = `The type of backend [%s] used as a metrics store with
Service Performance Monitoring (https://www.jaegertracing.io/docs/latest/spm/).
`
)

// Command creates `env` command
func Command() *cobra.Command {
	fs := new(pflag.FlagSet)
	fs.String(
		storage.SpanStorageTypeEnvVar,
		"cassandra",
		fmt.Sprintf(
			strings.ReplaceAll(storageTypeDescription, "\n", " "),
			strings.Join(storage.AllStorageTypes, ", "),
		),
	)
	fs.String(
		storage.DependencyStorageTypeEnvVar,
		"${SPAN_STORAGE_TYPE}",
		"The type of backend used for service dependencies storage.",
	)
	fs.String(
		strategystore.SamplingTypeEnvVar,
		"file",
		fmt.Sprintf(
			strings.ReplaceAll(samplingTypeDescription, "\n", " "),
			strings.Join(strategystore.AllSamplingTypes, ", "),
		),
	)
	fs.String(
		storage.SamplingStorageTypeEnvVar,
		"",
		fmt.Sprintf(
			strings.ReplaceAll(samplingStorageTypeDescription, "\n", " "),
			strings.Join(storage.AllSamplingStorageTypes(), ", "),
			strategystore.SamplingTypeEnvVar,
		),
	)
	fs.String(
		metrics.StorageTypeEnvVar,
		"",
		fmt.Sprintf(
			strings.ReplaceAll(metricsStorageTypeDescription, "\n", " "),
			strings.Join(metrics.AllStorageTypes, ", "),
		),
	)
	long := fmt.Sprintf(longTemplate, strings.Replace(fs.FlagUsagesWrapped(0), "      --", "\n", -1))
	return &cobra.Command{
		Use:   "env",
		Short: "Help about environment variables.",
		Long:  long,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(cmd.OutOrStdout(), long)
		},
	}
}

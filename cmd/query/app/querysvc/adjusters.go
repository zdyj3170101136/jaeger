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

package querysvc

import (
	"time"

	"github.com/zdyj3170101136/jaeger/model/adjuster"
)

// StandardAdjusters is a list of model adjusters applied by the query service
// before returning the data to the API clients.
func StandardAdjusters(maxClockSkewAdjust time.Duration) []adjuster.Adjuster {
	return []adjuster.Adjuster{
		adjuster.SpanIDDeduper(),
		adjuster.ClockSkew(maxClockSkewAdjust),
		adjuster.IPTagAdjuster(),
		adjuster.SortLogFields(),
		adjuster.SpanReferences(),
		adjuster.ParentReference(),
	}
}

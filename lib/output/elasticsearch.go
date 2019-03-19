// Copyright (c) 2018 Ashley Jeffs
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package output

import (
	"github.com/Jeffail/benthos/lib/log"
	"github.com/Jeffail/benthos/lib/metrics"
	"github.com/Jeffail/benthos/lib/output/writer"
	"github.com/Jeffail/benthos/lib/types"
)

//------------------------------------------------------------------------------

func init() {
	Constructors[TypeElasticsearch] = TypeSpec{
		constructor: NewElasticsearch,
		description: `
Publishes messages into an Elasticsearch index. This output currently does not
support creating the target index.

Both the ` + "`id` and `index`" + ` fields can be dynamically set using function
interpolations described [here](../config_interpolation.md#functions). When
sending batched messages these interpolations are performed per message part.`,
	}
}

//------------------------------------------------------------------------------

// NewElasticsearch creates a new Elasticsearch output type.
func NewElasticsearch(conf Config, mgr types.Manager, log log.Modular, stats metrics.Type) (Type, error) {
	elasticWriter, err := writer.NewElasticsearch(conf.Elasticsearch, log, stats)
	if err != nil {
		return nil, err
	}
	return NewWriter(
		"elasticsearch", elasticWriter, log, stats,
	)
}

//------------------------------------------------------------------------------
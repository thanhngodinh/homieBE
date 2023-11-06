// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/3b09f9d8e90178243f8a340a7bc324aab152c602

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// WeightedAverageValue type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3b09f9d8e90178243f8a340a7bc324aab152c602/specification/_types/aggregations/metric.ts#L448-L458
type WeightedAverageValue struct {
	// Field The field from which to extract the values or weights.
	Field *string `json:"field,omitempty"`
	// Missing A value or weight to use if the field is missing.
	Missing *Float64 `json:"missing,omitempty"`
	Script  Script   `json:"script,omitempty"`
}

func (s *WeightedAverageValue) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "missing":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Missing = &f
			case float64:
				f := Float64(v)
				s.Missing = &f
			}

		case "script":
			if err := dec.Decode(&s.Script); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewWeightedAverageValue returns a WeightedAverageValue.
func NewWeightedAverageValue() *WeightedAverageValue {
	r := &WeightedAverageValue{}

	return r
}

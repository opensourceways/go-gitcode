// Copyright (c) Huawei Technologies Co., Ltd. 2024. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package openapi

import (
	"time"
)

type timestamp time.Time

// UnmarshalJSON implements the json.Unmarshaler interface.
// Time is expected in RFC3339 or Unix format.
func (t *timestamp) UnmarshalJSON(data []byte) error {
	if len(data) == 2 {
		*t = timestamp(time.Time{})
		return nil
	}

	t1, err := time.Parse(`"`+time.RFC3339+`"`, string(data))
	*t = timestamp(t1)
	return err
}

func (t *timestamp) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(time.RFC3339)+2)
	b = append(b, '"')
	if !time.Time(*t).IsZero() {
		b = time.Time(*t).AppendFormat(b, time.RFC3339)
	}
	b = append(b, '"')
	return b, nil
}

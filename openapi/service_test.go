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
package openapi_test

import (
	"bytes"
	"encoding/gob"
	"net/http"
	"reflect"
	"testing"
)

func assertMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func assertReqBody(t *testing.T, r *http.Request, want any) {
	t.Helper()
	if got := r.Body; got != want {
		t.Errorf("Request body: %v, want %v", got, want)
	}
}

func assertDataLa(t *testing.T, got any, want any) {
	t.Helper()
	t1, t2 := reflect.TypeOf(got), reflect.TypeOf(want)
	if t1.String() != t2.String() {
		t.Errorf("mismatch data type, got: %v, want %v", t1.String(), t2.String())
	}

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(got); err != nil {
		t.Errorf("got: %v", err)
	}
	gotBytes := buf.Bytes()

	buf1 := new(bytes.Buffer)
	enc1 := gob.NewEncoder(buf1)
	if err := enc1.Encode(want); err != nil {
		t.Errorf("want: %v", err)
	}
	wantBytes := buf1.Bytes()

	if len(gotBytes) != len(wantBytes) {
		t.Errorf("mismatch data length, got: %v, want %v", len(gotBytes), len(wantBytes))
	}

	for i := 0; i < len(wantBytes); i++ {
		if gotBytes[i] != wantBytes[i] {
			t.Errorf("data different, got: %v, want %v", gotBytes, wantBytes)
			break
		}
	}
}

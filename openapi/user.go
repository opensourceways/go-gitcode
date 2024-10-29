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
	"context"
	"net/http"
)

// GetUserInfo 获取授权用户的资料
//
// api Docs: https://docs.gitcode.com/docs/openapi/users/#2-%e8%8e%b7%e5%8f%96%e6%8e%88%e6%9d%83%e7%94%a8%e6%88%b7%e7%9a%84%e8%b5%84%e6%96%99
func (s *UserService) GetUserInfo(ctx context.Context) (*User, bool, error) {
	req, err := newRequest(s.api, http.MethodGet, "user", nil)
	if err != nil {
		return nil, false, err
	}

	userInfo := new(User)
	resp, err := s.api.Do(ctx, req, userInfo)
	return userInfo, successGetData(resp), err
}

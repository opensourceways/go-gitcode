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
	"fmt"
	"net/http"
	"net/url"
)

// GetRepoAllMember 获取仓库的所有成员
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/member/#3-%e8%8e%b7%e5%8f%96%e4%bb%93%e5%ba%93%e7%9a%84%e6%89%80%e6%9c%89%e6%88%90%e5%91%98
func (s *RepositoryService) GetRepoAllMember(ctx context.Context, owner, repo, page string) ([]*User, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/collaborators", owner, repo)
	req, err := newRequest(s.api, http.MethodGet, urlStr,
		&url.Values{"page": []string{page}, "per_page": []string{"100"}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var members []*User
	resp, err := s.api.Do(ctx, req, &members)
	return members, successGetData(resp), err
}

// GetRepoMemberPermission 查看仓库成员的权限
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/member/#5-%e6%9f%a5%e7%9c%8b%e4%bb%93%e5%ba%93%e6%88%90%e5%91%98%e7%9a%84%e6%9d%83%e9%99%90
func (s *RepositoryService) GetRepoMemberPermission(ctx context.Context, owner, repo, login string) (*User, [2]bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/collaborators/%s/permission", owner, repo, login)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, [2]bool{false, false}, err
	}

	user := new(User)
	resp, err := s.api.Do(ctx, req, user)
	return user, [2]bool{successModified(resp), resp != nil && resp.StatusCode == http.StatusNotFound}, err
}

// CheckUserIsRepoMember 判断用户是否为仓库成员
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/member/#4-%e5%88%a4%e6%96%ad%e7%94%a8%e6%88%b7%e6%98%af%e5%90%a6%e4%b8%ba%e4%bb%93%e5%ba%93%e6%88%90%e5%91%98
func (s *RepositoryService) CheckUserIsRepoMember(ctx context.Context, owner, repo, username string) (bool, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/collaborators/%s", owner, repo, username)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return false, false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), resp != nil && resp.StatusCode == http.StatusNotFound, err
}

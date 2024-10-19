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

// GetRepoAllBranch 获取项目所有分支
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/branch/#1-%e8%8e%b7%e5%8f%96%e9%a1%b9%e7%9b%ae%e6%89%80%e6%9c%89%e5%88%86%e6%94%af
func (s *RepositoryService) GetRepoAllBranch(ctx context.Context, owner, repo, sort, direction, page string) ([]*Branch, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/branches", owner, repo)
	req, err := s.api.newRequest(http.MethodGet, urlStr,
		url.Values{"sort": []string{sort}, "direction": []string{direction}, "page": []string{page}, "per_page": []string{"100"}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var branches []*Branch
	resp, err := s.api.Do(ctx, req, &branches)
	return branches, successGetData(resp), err
}

// CreateRepoBranch 创建分支
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/branch/#2-%e5%88%9b%e5%bb%ba%e5%88%86%e6%94%af
func (s *RepositoryService) CreateRepoBranch(ctx context.Context, owner, repo, refs, name string) (*Branch, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/branches", owner, repo)
	req, err := s.api.newRequest(http.MethodPost, urlStr,
		url.Values{"refs": []string{refs}, "branch_name": []string{name}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	createdBranch := new(Branch)
	resp, err := s.api.Do(ctx, req, createdBranch)
	return createdBranch, successCreated(resp), err
}

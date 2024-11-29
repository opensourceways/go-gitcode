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

// GetRepoContributors 获取仓库贡献者
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/#9-%e8%8e%b7%e5%8f%96%e4%bb%93%e5%ba%93%e8%b4%a1%e7%8c%ae%e8%80%85
func (s *RepositoryService) GetRepoContributors(ctx context.Context, owner, repo, category string) ([]*Contributor, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/contributors", owner, repo)
	var query url.Values
	if category != "" {
		query = url.Values{}
		query.Set("type", category)
	}
	req, err := newRequest(s.api, http.MethodGet, urlStr, &query, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var contributors []*Contributor
	resp, err := s.api.Do(ctx, req, &contributors)
	return contributors, successGetData(resp), err
}

// GetRepoContentByPath 获取仓库具体路径下的内容
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/#2-%e8%8e%b7%e5%8f%96%e4%bb%93%e5%ba%93%e5%85%b7%e4%bd%93%e8%b7%af%e5%be%84%e4%b8%8b%e7%9a%84%e5%86%85%e5%ae%b9
func (s *RepositoryService) GetRepoContentByPath(ctx context.Context, owner, repo, path, ref string) (*RepositoryContent, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/contents/%s?ref=%s", owner, repo, path, ref)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	content := new(RepositoryContent)
	resp, err := s.api.Do(ctx, req, content)
	return content, successGetData(resp), err
}

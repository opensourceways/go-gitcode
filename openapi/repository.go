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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// GetRepoContentByPath 获取仓库具体路径下的内容
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/#2-%e8%8e%b7%e5%8f%96%e4%bb%93%e5%ba%93%e5%85%b7%e4%bd%93%e8%b7%af%e5%be%84%e4%b8%8b%e7%9a%84%e5%86%85%e5%ae%b9
func (s *RepositoryService) GetRepoContentByPath(ctx context.Context, owner, repo, path, ref string) ([]*RepositoryContent, bool, error) {
	if strings.Contains(path, "..") {
		return nil, false, pathForbiddenError
	}

	escapedPath := (&url.URL{Path: strings.TrimSuffix(path, "/")}).String()
	urlStr := fmt.Sprintf("repos/%s/%s/contents/%s", owner, repo, escapedPath)
	req, err := s.api.newRequest(http.MethodGet, urlStr, url.Values{"ref": []string{ref}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var rawJSON json.RawMessage
	resp, err := s.api.Do(ctx, req, &rawJSON)

	var directoryContent []*RepositoryContent
	if len(rawJSON) != 0 {
		if rawJSON[0] == '{' {
			fileContent := new(RepositoryContent)
			err = json.Unmarshal(rawJSON, fileContent)
			directoryContent = append(directoryContent, fileContent)
		} else {
			err = json.Unmarshal(rawJSON, &directoryContent)
		}
	}
	return directoryContent, successGetData(resp) && len(directoryContent) != 0, err
}

// GetRepoAllFileList 获取文件列表
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/#3-%e8%8e%b7%e5%8f%96%e6%96%87%e4%bb%b6%e5%88%97%e8%a1%a8
func (s *RepositoryService) GetRepoAllFileList(ctx context.Context, owner, repo, ref, file string) ([]string, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/file_list", owner, repo)
	req, err := s.api.newRequest(http.MethodGet, urlStr, url.Values{"ref_name": []string{ref}, "file_name": []string{file}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var fileList []string
	resp, err := s.api.Do(ctx, req, &fileList)
	return fileList, successGetData(resp), err
}

// GetRepoContributors 获取仓库贡献者
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/#9-%e8%8e%b7%e5%8f%96%e4%bb%93%e5%ba%93%e8%b4%a1%e7%8c%ae%e8%80%85
func (s *RepositoryService) GetRepoContributors(ctx context.Context, owner, repo, category string) ([]*Contributor, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/file_list", owner, repo)
	req, err := s.api.newRequest(http.MethodGet, urlStr, url.Values{"type": []string{category}}, RequestHandler{t: Query})
	if err != nil {
		return nil, false, err
	}

	var contributors []*Contributor
	resp, err := s.api.Do(ctx, req, &contributors)
	return contributors, successGetData(resp), err
}

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
	"strings"
)

// AddLabelsToPullRequest 创建 Pull Request 标签
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#12-%e5%88%9b%e5%bb%ba-pull-request-%e6%a0%87%e7%ad%be
func (s *PullRequestsService) AddLabelsToPullRequest(ctx context.Context, owner, repo, number string, labelNameList []string) ([]*Label, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/labels", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPost, urlStr, labelNameList)
	if err != nil {
		return nil, false, err
	}

	var addedLabels []*Label
	resp, err := s.api.Do(ctx, req, &addedLabels)
	return addedLabels, successCreated(resp), err
}

// RemoveLabelsFromPullRequest 删除 Pull Request 标签
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#13-%e5%88%a0%e9%99%a4-pull-request-%e6%a0%87%e7%ad%be
func (s *PullRequestsService) RemoveLabelsFromPullRequest(ctx context.Context, owner, repo, number string, labels []string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/labels/%s", owner, repo, number, strings.Join(labels, ","))
	req, err := newRequest(s.api, http.MethodDelete, urlStr, nil)
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// GetLabelsOfPullRequest 获取某个 Pull Request 的所有标签
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#16-%e8%8e%b7%e5%8f%96%e6%9f%90%e4%b8%aa-pull-request-%e7%9a%84%e6%89%80%e6%9c%89%e6%a0%87%e7%ad%be
func (s *PullRequestsService) GetLabelsOfPullRequest(ctx context.Context, owner, repo, number string) ([]*Label, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/labels", owner, repo, number)
	req, err := newRequest(s.api, http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var labels []*Label
	resp, err := s.api.Do(ctx, req, &labels)
	return labels, successGetData(resp), err
}

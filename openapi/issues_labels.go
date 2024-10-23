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

// ListRepoIssueLabels 获取仓库所有任务标签
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/labels/#2-%e8%8e%b7%e5%8f%96%e4%bb%93%e5%ba%93%e6%89%80%e6%9c%89%e4%bb%bb%e5%8a%a1%e6%a0%87%e7%ad%be
func (s *IssuesService) ListRepoIssueLabels(ctx context.Context, owner, repo string) ([]*Label, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/labels", owner, repo)
	req, err := s.api.newRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, false, err
	}

	var labels []*Label
	resp, err := s.api.Do(ctx, req, &labels)
	return labels, successCreated(resp), nil
}

// CreateRepoIssueLabel 创建仓库任务标签
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/labels/#3-%e5%88%9b%e5%bb%ba%e4%bb%93%e5%ba%93%e4%bb%bb%e5%8a%a1%e6%a0%87%e7%ad%be
func (s *IssuesService) CreateRepoIssueLabel(ctx context.Context, owner, repo string, newLabel *Label) (*Label, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/labels", owner, repo)
	req, err := s.api.newRequest(http.MethodPost, urlStr, newLabel, RequestHandler{t: Form})
	if err != nil {
		return nil, false, err
	}

	label := new(Label)
	resp, err := s.api.Do(ctx, req, label)
	return label, successCreated(resp), err
}

// UpdateRepoIssueLabel 更新一个仓库的任务标签
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/labels/#1-%e6%9b%b4%e6%96%b0%e4%b8%80%e4%b8%aa%e4%bb%93%e5%ba%93%e7%9a%84%e4%bb%bb%e5%8a%a1%e6%a0%87%e7%ad%be
func (s *IssuesService) UpdateRepoIssueLabel(ctx context.Context, owner, repo, originalName, newName, color string) (*Label, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/labels/%s", owner, repo, originalName)
	req, err := s.api.newRequest(http.MethodPatch, urlStr, &url.Values{"name": []string{newName}, "color": []string{color}}, RequestHandler{t: Form})

	editedLabel := new(Label)
	resp, err := s.api.Do(ctx, req, editedLabel)
	return editedLabel, successModified(resp), err
}

// DeleteRepoIssueLabel 删除一个仓库任务标签
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/labels/#4-%e5%88%a0%e9%99%a4%e4%b8%80%e4%b8%aa%e4%bb%93%e5%ba%93%e4%bb%bb%e5%8a%a1%e6%a0%87%e7%ad%be
func (s *IssuesService) DeleteRepoIssueLabel(ctx context.Context, owner, repo, name string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/labels/%s", owner, repo, name)
	req, err := s.api.newRequest(http.MethodDelete, urlStr, nil)
	if err != nil {
		return false, err
	}
	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

// AddLabelsToIssue 创建Issue标签
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/issues/#9%e5%88%9b%e5%bb%baissue%e6%a0%87%e7%ad%be
func (s *IssuesService) AddLabelsToIssue(ctx context.Context, owner, repo, number string, labelNameList []string) ([]*Label, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/issues/%s/labels", owner, repo, number)
	req, err := s.api.newRequest(http.MethodPost, urlStr, labelNameList)
	if err != nil {
		return nil, false, err
	}

	var havingLabels []*Label
	resp, err := s.api.Do(ctx, req, &havingLabels)
	return havingLabels, successCreated(resp), err
}

// RemoveLabelsFromIssue 删除Issue标签
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/issues/#10%e5%88%a0%e9%99%a4issue%e6%a0%87%e7%ad%be
func (s *IssuesService) RemoveLabelsFromIssue(ctx context.Context, owner, repo, number, labels string) (bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/issues/%s/labels/%s", owner, repo, number, labels)
	req, err := s.api.newRequest(http.MethodDelete, urlStr, nil)
	if err != nil {
		return false, err
	}

	resp, err := s.api.Do(ctx, req, nil)
	return successModified(resp), err
}

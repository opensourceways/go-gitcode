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
)

// CreateIssueComment 创建Issue评论
//
// API Docs: https://docs.gitcode.com/docs/openapi/repos/issues/#11%e5%88%9b%e5%bb%baissue%e8%af%84%e8%ae%ba
func (s *IssuesService) CreateIssueComment(ctx context.Context, owner, repo, number string, comment *IssueComment) (*IssueComment, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/issues/%s/comments", owner, repo, number)
	req, err := s.api.newRequest(http.MethodPost, urlStr, comment)
	if err != nil {
		return nil, false, err
	}

	addedComment := new(IssueComment)
	resp, err := s.api.Do(ctx, req, addedComment)
	return addedComment, successCreated(resp), err
}

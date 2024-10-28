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

// CreatePullRequestComment 提交 pull request 评论
//
// api Docs: https://docs.gitcode.com/docs/openapi/repos/pulls/#4-%e6%8f%90%e4%ba%a4pull-request-%e8%af%84%e8%ae%ba
func (s *PullRequestsService) CreatePullRequestComment(ctx context.Context, owner, repo, number string, comment *PullRequestCommentRequest) (*PullRequestComment, bool, error) {
	urlStr := fmt.Sprintf("repos/%s/%s/pulls/%s/comments", owner, repo, number)
	req, err := newRequest(s.api, http.MethodPost, urlStr, comment)
	if err != nil {
		return nil, false, err
	}

	addedComment := new(PullRequestComment)
	resp, err := s.api.Do(ctx, req, addedComment)
	return addedComment, successCreated(resp), err
}

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

import "time"

// Label represents a GitCode label on an Issue
type Label struct {
	ID           *int64  `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	Color        *string `json:"color,omitempty"`
	RepositoryId *int64  `json:"repository_id,omitempty"`
}

type Issue struct {
	ID               *int64            `json:"id,omitempty"`
	HTMLURL          *string           `json:"html_url,omitempty"`
	Number           *string           `json:"number,omitempty"`
	State            *string           `json:"state,omitempty"`
	IssueState       *string           `json:"issue_state,omitempty"`
	IssueStateDetail *IssueStateDetail `json:"issue_state_detail,omitempty"`
	Priority         *int              `json:"priority,omitempty"`
	Title            *string           `json:"title,omitempty"`
	Body             *string           `json:"body,omitempty"`
	User             *User             `json:"user,omitempty"`
	Assignee         *User             `json:"assignee,omitempty"`
	Repository       *Repository       `json:"repository,omitempty"`
	Labels           []*Label          `json:"labels,omitempty"`
	CreatedAt        *time.Time        `json:"created_at,omitempty"`
	UpdatedAt        *time.Time        `json:"updated_at,omitempty"`
	ClosedAt         *time.Time        `json:"finished_at,omitempty"`
	ClosedBy         *User             `json:"closed_by,omitempty"`

	PullRequestLinks *PullRequestLinks `json:"pull_request,omitempty"` // TODO
	// RepoBranchLink  // TODO

}

type IssueStateDetail struct {
	Title  *string `json:"title,omitempty"`
	Serial *int    `json:"serial,omitempty"`
}

type PullRequestLinks struct {
	URL      *string    `json:"url,omitempty"`
	HTMLURL  *string    `json:"html_url,omitempty"`
	DiffURL  *string    `json:"diff_url,omitempty"`
	PatchURL *string    `json:"patch_url,omitempty"`
	MergedAt *time.Time `json:"merged_at,omitempty"`
}

type IssueRequest struct {
	Repository    *string `json:"repo,omitempty"  required:"true"` // 仓库地址
	Title         *string `json:"title,omitempty"`
	Body          *string `json:"body,omitempty"`
	Labels        *string `json:"labels,omitempty"`   // 用逗号分开的标签
	Assignee      *string `json:"assignee,omitempty"` // Issue负责人的 username
	State         *string `json:"state,omitempty"`
	Milestone     *int64  `json:"milestone,omitempty"`
	SecurityHole  *string `json:"security_hole,omitempty"`  // 是否是私有issue(默认为false)
	IssueStage    *string `json:"issue_stage,omitempty"`    // 严重程序（Accepted,Coding,Completed,New,Rejected,Revising,Testing,Verified）
	IssueSeverity *string `json:"issue_severity,omitempty"` // 优先级 （Suggestion,Minor,Major,Fatal）
}

type IssueComment struct {
	ID        *int64     `json:"id,omitempty"`
	Body      *string    `json:"body,omitempty"`
	User      *User      `json:"user,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Target    *Issue     `json:"target,omitempty"`
}

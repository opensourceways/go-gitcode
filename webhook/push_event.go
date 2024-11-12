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
package webhook

type PushEvent struct {
	UUID         *string  `json:"uuid,omitempty"`
	EventType    *string  `json:"event_name,omitempty"`
	ObjectKind   *string  `json:"object_kind,omitempty"`
	ManualBuild  *bool    `json:"manual_build,omitempty"`
	Repository   *Project `json:"project,omitempty"`
	SourceBranch *string  `json:"git_branch,omitempty"`
	Author       *string  `json:"user_username,omitempty"`
}

func (p *PushEvent) GetAction() *string {
	return nil
}
func (p *PushEvent) GetState() *string {
	return nil
}
func (p *PushEvent) GetOrg() *string {
	if p.Repository == nil {
		return nil
	}

	return p.Repository.Namespace
}
func (p *PushEvent) GetRepo() *string {
	if p.Repository == nil {
		return nil
	}

	return p.Repository.Name
}
func (p *PushEvent) GetHtmlURL() *string {
	if p.Repository == nil {
		return nil
	}

	return p.Repository.HTMLURL
}
func (p *PushEvent) GetBase() *string {
	return p.SourceBranch
}
func (p *PushEvent) GetHead() *string {
	return nil
}
func (p *PushEvent) GetNumber() *string {
	return nil
}
func (p *PushEvent) GetAuthor() *string {
	return p.Author
}
func (p *PushEvent) GetComment() *string {
	return nil
}
func (p *PushEvent) GetCommenter() *string {
	return nil
}
func (p *PushEvent) ListLabels() []*string {
	return nil
}

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

// User represents a GitHub user.
type User struct {
	Login       *string     `json:"login,omitempty"`
	ID          *int64      `json:"id,omitempty"`
	AvatarURL   *string     `json:"avatar_url,omitempty"`
	HTMLURL     *string     `json:"html_url,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Company     *string     `json:"company,omitempty"`
	Blog        *string     `json:"blog,omitempty"`
	Email       *string     `json:"email,omitempty"`
	Bio         *string     `json:"bio,omitempty"`
	Followers   *int        `json:"followers,omitempty"`
	Following   *int        `json:"following,omitempty"`
	Type        *string     `json:"type,omitempty"`
	Permission  *string     `json:"permission,omitempty"`
	Permissions *Permission `json:"permissions,omitempty"`
}

type Permission struct {
	Admin *bool `json:"admin,omitempty"`
}

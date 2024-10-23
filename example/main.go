// Copyright 2024 Chao Feng
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
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opensourceways/go-gitcode/openapi"
	"log"
	"os"
)

func main() {
	token := os.Getenv("GITCODE_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}
	ctx := context.Background()
	client := openapi.NewAPIClientWithAuthorization([]byte(token))

	//labels, ok, err := client.Issues.ListRepoIssueLabels(ctx, "ibforuorg", "test1")
	//fmt.Printf("success: %v, error: %v \n", ok, err)
	//if ok {
	//	for _, label := range labels {
	//		d, _ := json.Marshal(label)
	//		fmt.Printf("label: %v \n", string(d))
	//	}
	//}

	//pr, ok, err := client.Issues.ListIssueLinkingPullRequests(ctx, "ibforuorg", "test1", "1")
	//fmt.Printf("success: %v, error: %v \n", ok, err)
	//if ok {
	//	for _, p := range pr {
	//		d, _ := json.Marshal(p)
	//		fmt.Printf("pr: %v \n", string(d))
	//	}
	//}

	pr, ok, err := client.PullRequests.GetPullRequest(ctx, "ibforuorg", "test1", "1")
	fmt.Printf("success: %v, error: %v \n", ok, err)
	if ok {
		d, _ := json.Marshal(pr)
		fmt.Printf("pr: %v \n", string(d))
	}
}

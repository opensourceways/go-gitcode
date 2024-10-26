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

	contributors, ok, err := client.Repository.GetRepoContributors(ctx, "ibforuorg", "test1", "")

	//issue, ok, err := client.Issues.UpdateIssue(ctx, "ibforuorg", "2", &openapi.IssueRequest{
	//	Repository: "test1",
	//	Title:      "issue1",
	//})
	fmt.Printf("success: %v, error: %v \n", ok, err)
	if ok {
		d, _ := json.Marshal(contributors)
		fmt.Printf("pr: %v \n", string(d))
	}

	//labels, ok, err := client.Issues.ListRepoIssueLabels(ctx, "ibforuorg", "test1")
	//fmt.Printf("success: %v, error: %v \n", ok, err)
	//if ok {
	//	for _, label := range labels {
	//		d, _ := json.Marshal(label)
	//		fmt.Printf("label: %v \n", string(d))
	//	}
	//}

	//comment := "fajhgdahjksghj"
	//issueComment, ok, err := client.Issues.CreateIssueComment(ctx, "ibforuorg", "test1", "1", &openapi.IssueComment{
	//	Body: &comment,
	//})
	//fmt.Printf("success: %v, error: %v \n\n", ok, err)
	//if ok {
	//	d, _ := json.Marshal(issueComment)
	//	fmt.Printf("pr: %v \n", string(d))
	//}

	//result, ok, err := client.Issues.CreateRepoIssueLabel(ctx, "ibforuorg", "test1", &openapi.Label{Name: "fasdsad", Color: "#fff"})
	//result, ok, err := client.Issues.UpdateRepoIssueLabel(ctx, "ibforuorg", "test1", "giasdlkggds", "fsaghhhhh", "#000000")
	//ok, err := client.Issues.DeleteRepoIssueLabel(ctx, "ibforuorg", "test1", "fgagasdasda")
	//result, ok, err := client.Issues.AddLabelsToIssue(ctx, "ibforuorg", "test1", "1", []string{"fsaghhhhh"})
	//ok, err := client.Issues.RemoveLabelsFromIssue(ctx, "ibforuorg", "test1", "1", "fsaghhhhh")
	//fmt.Printf("success: %v, error: %v \n\n", ok, err)
	//if ok {
	//	d, _ := json.Marshal(result)
	//	fmt.Printf("result: %v \n", string(d))
	//}

	//pr, ok, err := client.Issues.ListIssueLinkingPullRequests(ctx, "ibforuorg", "test1", "1")
	//fmt.Printf("success: %v, error: %v \n", ok, err)
	//if ok {
	//	for _, p := range pr {
	//		d, _ := json.Marshal(p)
	//		fmt.Printf("pr: %v \n", string(d))
	//	}
	//}

	//result, ok, err := client.PullRequests.GetPullRequest(ctx, "ibforuorg", "test1", "1")
	//result, ok, err := client.PullRequests.UpdatePullRequest(ctx, "ibforuorg", "test1", "1", &openapi.PullRequestRequest{
	//	State: "open",
	//})
	//result, ok, err := client.PullRequests.ListPullRequestLinkingIssues(ctx, "ibforuorg", "test1", "1")
	//result, ok, err := client.PullRequests.CreatePullRequestComment(ctx, "ibforuorg", "test1", "1", &openapi.PullRequestCommentRequest{
	//	Body: "fauygiahsgdbviahsd",
	//})
	//result, ok, err := client.PullRequests.AddLabelsToPullRequest(ctx, "ibforuorg", "test1", "1", []string{"bug1"})
	//ok, err := client.PullRequests.RemoveLabelsFromPullRequest(ctx, "ibforuorg", "test1", "1", "bug1")
	//result, ok, err := client.Repository.CheckUserIsRepoMember(ctx, "ibforuorg", "test1", "ibforu2nd")
	//fmt.Printf("success: %v, error: %v \n\n", result, err)
	//if ok {
	//	d, _ := json.Marshal(result)
	//	fmt.Printf("result: %v \n", string(d))
	//}

	//request, err := http.NewRequest(http.MethodGet, "https://api.gitcode.com/api/v5/repos/ibforuorg/test1/contributors", nil)
	//request.Header.Set("Authorization", "Bearer gouQ2vVmbEqMhzMxCoTTQfdN")
	//
	//do, err := http.DefaultClient.Do(request)
	//if err != nil {
	//	fmt.Printf("err: %v \n", err)
	//}
	//
	//fmt.Printf("resp: %v \n", do)
}

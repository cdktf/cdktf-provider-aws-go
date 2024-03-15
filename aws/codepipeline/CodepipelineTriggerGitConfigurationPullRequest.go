// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineTriggerGitConfigurationPullRequest struct {
	// branches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/codepipeline#branches Codepipeline#branches}
	Branches *CodepipelineTriggerGitConfigurationPullRequestBranches `field:"optional" json:"branches" yaml:"branches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/codepipeline#events Codepipeline#events}.
	Events *[]*string `field:"optional" json:"events" yaml:"events"`
	// file_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/codepipeline#file_paths Codepipeline#file_paths}
	FilePaths *CodepipelineTriggerGitConfigurationPullRequestFilePaths `field:"optional" json:"filePaths" yaml:"filePaths"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineTriggerGitConfigurationPush struct {
	// branches block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/codepipeline#branches Codepipeline#branches}
	Branches *CodepipelineTriggerGitConfigurationPushBranches `field:"optional" json:"branches" yaml:"branches"`
	// file_paths block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/codepipeline#file_paths Codepipeline#file_paths}
	FilePaths *CodepipelineTriggerGitConfigurationPushFilePaths `field:"optional" json:"filePaths" yaml:"filePaths"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/codepipeline#tags Codepipeline#tags}
	Tags *CodepipelineTriggerGitConfigurationPushTags `field:"optional" json:"tags" yaml:"tags"`
}


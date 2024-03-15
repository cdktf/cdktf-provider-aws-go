// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineTriggerGitConfigurationPullRequestFilePaths struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/codepipeline#excludes Codepipeline#excludes}.
	Excludes *[]*string `field:"optional" json:"excludes" yaml:"excludes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.41.0/docs/resources/codepipeline#includes Codepipeline#includes}.
	Includes *[]*string `field:"optional" json:"includes" yaml:"includes"`
}


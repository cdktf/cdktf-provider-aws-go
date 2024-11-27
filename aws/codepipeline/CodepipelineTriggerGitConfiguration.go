// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineTriggerGitConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/codepipeline#source_action_name Codepipeline#source_action_name}.
	SourceActionName *string `field:"required" json:"sourceActionName" yaml:"sourceActionName"`
	// pull_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/codepipeline#pull_request Codepipeline#pull_request}
	PullRequest interface{} `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// push block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/codepipeline#push Codepipeline#push}
	Push interface{} `field:"optional" json:"push" yaml:"push"`
}


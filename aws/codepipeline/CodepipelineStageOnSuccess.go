// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineStageOnSuccess struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.98.0/docs/resources/codepipeline#condition Codepipeline#condition}
	Condition *CodepipelineStageOnSuccessCondition `field:"required" json:"condition" yaml:"condition"`
}


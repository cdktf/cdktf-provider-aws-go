// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineStageBeforeEntry struct {
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline#condition Codepipeline#condition}
	Condition *CodepipelineStageBeforeEntryCondition `field:"required" json:"condition" yaml:"condition"`
}


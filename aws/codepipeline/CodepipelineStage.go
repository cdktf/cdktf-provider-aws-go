// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineStage struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/codepipeline#action Codepipeline#action}
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/codepipeline#name Codepipeline#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// before_entry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/codepipeline#before_entry Codepipeline#before_entry}
	BeforeEntry *CodepipelineStageBeforeEntry `field:"optional" json:"beforeEntry" yaml:"beforeEntry"`
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/codepipeline#on_failure Codepipeline#on_failure}
	OnFailure *CodepipelineStageOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
	// on_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/codepipeline#on_success Codepipeline#on_success}
	OnSuccess *CodepipelineStageOnSuccess `field:"optional" json:"onSuccess" yaml:"onSuccess"`
}


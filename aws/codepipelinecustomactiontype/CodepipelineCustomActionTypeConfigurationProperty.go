// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipelinecustomactiontype


type CodepipelineCustomActionTypeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/codepipeline_custom_action_type#key CodepipelineCustomActionType#key}.
	Key interface{} `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/codepipeline_custom_action_type#name CodepipelineCustomActionType#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/codepipeline_custom_action_type#required CodepipelineCustomActionType#required}.
	Required interface{} `field:"required" json:"required" yaml:"required"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/codepipeline_custom_action_type#secret CodepipelineCustomActionType#secret}.
	Secret interface{} `field:"required" json:"secret" yaml:"secret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/codepipeline_custom_action_type#description CodepipelineCustomActionType#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/codepipeline_custom_action_type#queryable CodepipelineCustomActionType#queryable}.
	Queryable interface{} `field:"optional" json:"queryable" yaml:"queryable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/codepipeline_custom_action_type#type CodepipelineCustomActionType#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


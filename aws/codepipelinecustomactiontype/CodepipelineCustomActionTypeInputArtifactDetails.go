// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipelinecustomactiontype


type CodepipelineCustomActionTypeInputArtifactDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/codepipeline_custom_action_type#maximum_count CodepipelineCustomActionType#maximum_count}.
	MaximumCount *float64 `field:"required" json:"maximumCount" yaml:"maximumCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/codepipeline_custom_action_type#minimum_count CodepipelineCustomActionType#minimum_count}.
	MinimumCount *float64 `field:"required" json:"minimumCount" yaml:"minimumCount"`
}


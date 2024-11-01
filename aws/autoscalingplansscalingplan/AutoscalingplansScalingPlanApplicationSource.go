// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan


type AutoscalingplansScalingPlanApplicationSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/autoscalingplans_scaling_plan#cloudformation_stack_arn AutoscalingplansScalingPlan#cloudformation_stack_arn}.
	CloudformationStackArn *string `field:"optional" json:"cloudformationStackArn" yaml:"cloudformationStackArn"`
	// tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/autoscalingplans_scaling_plan#tag_filter AutoscalingplansScalingPlan#tag_filter}
	TagFilter interface{} `field:"optional" json:"tagFilter" yaml:"tagFilter"`
}


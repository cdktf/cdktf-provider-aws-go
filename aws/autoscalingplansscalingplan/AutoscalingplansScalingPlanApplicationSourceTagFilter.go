// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan


type AutoscalingplansScalingPlanApplicationSourceTagFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/autoscalingplans_scaling_plan#key AutoscalingplansScalingPlan#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/autoscalingplans_scaling_plan#values AutoscalingplansScalingPlan#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}


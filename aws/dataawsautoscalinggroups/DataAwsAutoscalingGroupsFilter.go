// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsautoscalinggroups


type DataAwsAutoscalingGroupsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/data-sources/autoscaling_groups#name DataAwsAutoscalingGroups#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/data-sources/autoscaling_groups#values DataAwsAutoscalingGroups#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


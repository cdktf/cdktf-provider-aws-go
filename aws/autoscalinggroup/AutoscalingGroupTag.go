// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/autoscaling_group#key AutoscalingGroup#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/autoscaling_group#propagate_at_launch AutoscalingGroup#propagate_at_launch}.
	PropagateAtLaunch interface{} `field:"required" json:"propagateAtLaunch" yaml:"propagateAtLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/autoscaling_group#value AutoscalingGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}


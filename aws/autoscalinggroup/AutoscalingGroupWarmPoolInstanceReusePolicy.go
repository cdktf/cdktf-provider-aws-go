// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupWarmPoolInstanceReusePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/autoscaling_group#reuse_on_scale_in AutoscalingGroup#reuse_on_scale_in}.
	ReuseOnScaleIn interface{} `field:"optional" json:"reuseOnScaleIn" yaml:"reuseOnScaleIn"`
}


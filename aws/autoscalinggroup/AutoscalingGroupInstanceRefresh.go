// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupInstanceRefresh struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/autoscaling_group#strategy AutoscalingGroup#strategy}.
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/autoscaling_group#preferences AutoscalingGroup#preferences}
	Preferences *AutoscalingGroupInstanceRefreshPreferences `field:"optional" json:"preferences" yaml:"preferences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/autoscaling_group#triggers AutoscalingGroup#triggers}.
	Triggers *[]*string `field:"optional" json:"triggers" yaml:"triggers"`
}


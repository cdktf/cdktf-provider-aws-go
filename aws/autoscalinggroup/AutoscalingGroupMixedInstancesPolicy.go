// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupMixedInstancesPolicy struct {
	// launch_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/autoscaling_group#launch_template AutoscalingGroup#launch_template}
	LaunchTemplate *AutoscalingGroupMixedInstancesPolicyLaunchTemplate `field:"required" json:"launchTemplate" yaml:"launchTemplate"`
	// instances_distribution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/autoscaling_group#instances_distribution AutoscalingGroup#instances_distribution}
	InstancesDistribution *AutoscalingGroupMixedInstancesPolicyInstancesDistribution `field:"optional" json:"instancesDistribution" yaml:"instancesDistribution"`
}


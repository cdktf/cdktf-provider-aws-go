// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupAvailabilityZoneDistribution struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/autoscaling_group#capacity_distribution_strategy AutoscalingGroup#capacity_distribution_strategy}.
	CapacityDistributionStrategy *string `field:"optional" json:"capacityDistributionStrategy" yaml:"capacityDistributionStrategy"`
}


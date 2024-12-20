// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupMixedInstancesPolicyInstancesDistribution struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/autoscaling_group#on_demand_allocation_strategy AutoscalingGroup#on_demand_allocation_strategy}.
	OnDemandAllocationStrategy *string `field:"optional" json:"onDemandAllocationStrategy" yaml:"onDemandAllocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/autoscaling_group#on_demand_base_capacity AutoscalingGroup#on_demand_base_capacity}.
	OnDemandBaseCapacity *float64 `field:"optional" json:"onDemandBaseCapacity" yaml:"onDemandBaseCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/autoscaling_group#on_demand_percentage_above_base_capacity AutoscalingGroup#on_demand_percentage_above_base_capacity}.
	OnDemandPercentageAboveBaseCapacity *float64 `field:"optional" json:"onDemandPercentageAboveBaseCapacity" yaml:"onDemandPercentageAboveBaseCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/autoscaling_group#spot_allocation_strategy AutoscalingGroup#spot_allocation_strategy}.
	SpotAllocationStrategy *string `field:"optional" json:"spotAllocationStrategy" yaml:"spotAllocationStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/autoscaling_group#spot_instance_pools AutoscalingGroup#spot_instance_pools}.
	SpotInstancePools *float64 `field:"optional" json:"spotInstancePools" yaml:"spotInstancePools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/autoscaling_group#spot_max_price AutoscalingGroup#spot_max_price}.
	SpotMaxPrice *string `field:"optional" json:"spotMaxPrice" yaml:"spotMaxPrice"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalinggroup


type AutoscalingGroupWarmPool struct {
	// instance_reuse_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/autoscaling_group#instance_reuse_policy AutoscalingGroup#instance_reuse_policy}
	InstanceReusePolicy *AutoscalingGroupWarmPoolInstanceReusePolicy `field:"optional" json:"instanceReusePolicy" yaml:"instanceReusePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/autoscaling_group#max_group_prepared_capacity AutoscalingGroup#max_group_prepared_capacity}.
	MaxGroupPreparedCapacity *float64 `field:"optional" json:"maxGroupPreparedCapacity" yaml:"maxGroupPreparedCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/autoscaling_group#min_size AutoscalingGroup#min_size}.
	MinSize *float64 `field:"optional" json:"minSize" yaml:"minSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/autoscaling_group#pool_state AutoscalingGroup#pool_state}.
	PoolState *string `field:"optional" json:"poolState" yaml:"poolState"`
}


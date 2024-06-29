// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emrmanagedscalingpolicy


type EmrManagedScalingPolicyComputeLimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/emr_managed_scaling_policy#maximum_capacity_units EmrManagedScalingPolicy#maximum_capacity_units}.
	MaximumCapacityUnits *float64 `field:"required" json:"maximumCapacityUnits" yaml:"maximumCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/emr_managed_scaling_policy#minimum_capacity_units EmrManagedScalingPolicy#minimum_capacity_units}.
	MinimumCapacityUnits *float64 `field:"required" json:"minimumCapacityUnits" yaml:"minimumCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/emr_managed_scaling_policy#unit_type EmrManagedScalingPolicy#unit_type}.
	UnitType *string `field:"required" json:"unitType" yaml:"unitType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/emr_managed_scaling_policy#maximum_core_capacity_units EmrManagedScalingPolicy#maximum_core_capacity_units}.
	MaximumCoreCapacityUnits *float64 `field:"optional" json:"maximumCoreCapacityUnits" yaml:"maximumCoreCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/emr_managed_scaling_policy#maximum_ondemand_capacity_units EmrManagedScalingPolicy#maximum_ondemand_capacity_units}.
	MaximumOndemandCapacityUnits *float64 `field:"optional" json:"maximumOndemandCapacityUnits" yaml:"maximumOndemandCapacityUnits"`
}


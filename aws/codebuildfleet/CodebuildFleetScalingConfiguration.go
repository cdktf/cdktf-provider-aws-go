// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildfleet


type CodebuildFleetScalingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/codebuild_fleet#max_capacity CodebuildFleet#max_capacity}.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/codebuild_fleet#scaling_type CodebuildFleet#scaling_type}.
	ScalingType *string `field:"optional" json:"scalingType" yaml:"scalingType"`
	// target_tracking_scaling_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.2/docs/resources/codebuild_fleet#target_tracking_scaling_configs CodebuildFleet#target_tracking_scaling_configs}
	TargetTrackingScalingConfigs interface{} `field:"optional" json:"targetTrackingScalingConfigs" yaml:"targetTrackingScalingConfigs"`
}


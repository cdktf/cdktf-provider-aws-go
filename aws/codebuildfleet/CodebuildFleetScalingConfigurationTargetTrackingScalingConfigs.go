// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildfleet


type CodebuildFleetScalingConfigurationTargetTrackingScalingConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/codebuild_fleet#metric_type CodebuildFleet#metric_type}.
	MetricType *string `field:"optional" json:"metricType" yaml:"metricType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/codebuild_fleet#target_value CodebuildFleet#target_value}.
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingpolicy


type AutoscalingPolicyTargetTrackingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/autoscaling_policy#target_value AutoscalingPolicy#target_value}.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// customized_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/autoscaling_policy#customized_metric_specification AutoscalingPolicy#customized_metric_specification}
	CustomizedMetricSpecification *AutoscalingPolicyTargetTrackingConfigurationCustomizedMetricSpecification `field:"optional" json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/autoscaling_policy#disable_scale_in AutoscalingPolicy#disable_scale_in}.
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// predefined_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/autoscaling_policy#predefined_metric_specification AutoscalingPolicy#predefined_metric_specification}
	PredefinedMetricSpecification *AutoscalingPolicyTargetTrackingConfigurationPredefinedMetricSpecification `field:"optional" json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
}


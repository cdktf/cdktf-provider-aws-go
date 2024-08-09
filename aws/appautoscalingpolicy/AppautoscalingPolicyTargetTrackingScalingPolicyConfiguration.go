// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appautoscalingpolicy


type AppautoscalingPolicyTargetTrackingScalingPolicyConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/appautoscaling_policy#target_value AppautoscalingPolicy#target_value}.
	TargetValue *float64 `field:"required" json:"targetValue" yaml:"targetValue"`
	// customized_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/appautoscaling_policy#customized_metric_specification AppautoscalingPolicy#customized_metric_specification}
	CustomizedMetricSpecification *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification `field:"optional" json:"customizedMetricSpecification" yaml:"customizedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/appautoscaling_policy#disable_scale_in AppautoscalingPolicy#disable_scale_in}.
	DisableScaleIn interface{} `field:"optional" json:"disableScaleIn" yaml:"disableScaleIn"`
	// predefined_metric_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/appautoscaling_policy#predefined_metric_specification AppautoscalingPolicy#predefined_metric_specification}
	PredefinedMetricSpecification *AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification `field:"optional" json:"predefinedMetricSpecification" yaml:"predefinedMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/appautoscaling_policy#scale_in_cooldown AppautoscalingPolicy#scale_in_cooldown}.
	ScaleInCooldown *float64 `field:"optional" json:"scaleInCooldown" yaml:"scaleInCooldown"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/appautoscaling_policy#scale_out_cooldown AppautoscalingPolicy#scale_out_cooldown}.
	ScaleOutCooldown *float64 `field:"optional" json:"scaleOutCooldown" yaml:"scaleOutCooldown"`
}


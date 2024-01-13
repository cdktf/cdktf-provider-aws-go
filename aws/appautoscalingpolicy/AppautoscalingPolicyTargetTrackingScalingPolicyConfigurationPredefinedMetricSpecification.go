// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appautoscalingpolicy


type AppautoscalingPolicyTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/appautoscaling_policy#predefined_metric_type AppautoscalingPolicy#predefined_metric_type}.
	PredefinedMetricType *string `field:"required" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/appautoscaling_policy#resource_label AppautoscalingPolicy#resource_label}.
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}


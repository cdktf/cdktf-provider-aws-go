// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan


type AutoscalingplansScalingPlanScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/autoscalingplans_scaling_plan#predefined_scaling_metric_type AutoscalingplansScalingPlan#predefined_scaling_metric_type}.
	PredefinedScalingMetricType *string `field:"required" json:"predefinedScalingMetricType" yaml:"predefinedScalingMetricType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/autoscalingplans_scaling_plan#resource_label AutoscalingplansScalingPlan#resource_label}.
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}


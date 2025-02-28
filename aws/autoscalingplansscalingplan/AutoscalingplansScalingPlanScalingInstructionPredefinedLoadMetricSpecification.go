// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan


type AutoscalingplansScalingPlanScalingInstructionPredefinedLoadMetricSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/autoscalingplans_scaling_plan#predefined_load_metric_type AutoscalingplansScalingPlan#predefined_load_metric_type}.
	PredefinedLoadMetricType *string `field:"required" json:"predefinedLoadMetricType" yaml:"predefinedLoadMetricType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/autoscalingplans_scaling_plan#resource_label AutoscalingplansScalingPlan#resource_label}.
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}


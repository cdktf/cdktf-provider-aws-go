package autoscalingplans


type AutoscalingplansScalingPlanScalingInstructionPredefinedLoadMetricSpecification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/autoscalingplans_scaling_plan#predefined_load_metric_type AutoscalingplansScalingPlan#predefined_load_metric_type}.
	PredefinedLoadMetricType *string `field:"required" json:"predefinedLoadMetricType" yaml:"predefinedLoadMetricType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/autoscalingplans_scaling_plan#resource_label AutoscalingplansScalingPlan#resource_label}.
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}


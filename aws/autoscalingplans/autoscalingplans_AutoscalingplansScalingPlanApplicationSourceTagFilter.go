package autoscalingplans


type AutoscalingplansScalingPlanApplicationSourceTagFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/autoscalingplans_scaling_plan#key AutoscalingplansScalingPlan#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/autoscalingplans_scaling_plan#values AutoscalingplansScalingPlan#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}


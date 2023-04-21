package appautoscalingscheduledaction


type AppautoscalingScheduledActionScalableTargetAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/appautoscaling_scheduled_action#max_capacity AppautoscalingScheduledAction#max_capacity}.
	MaxCapacity *string `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/appautoscaling_scheduled_action#min_capacity AppautoscalingScheduledAction#min_capacity}.
	MinCapacity *string `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}


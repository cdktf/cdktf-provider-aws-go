package appautoscalingscheduledaction


type AppautoscalingScheduledActionScalableTargetAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#max_capacity AppautoscalingScheduledAction#max_capacity}.
	MaxCapacity *string `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/appautoscaling_scheduled_action#min_capacity AppautoscalingScheduledAction#min_capacity}.
	MinCapacity *string `field:"optional" json:"minCapacity" yaml:"minCapacity"`
}


package schedulerschedule


type SchedulerScheduleTargetEcsParametersPlacementStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/scheduler_schedule#type SchedulerSchedule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.63.0/docs/resources/scheduler_schedule#field SchedulerSchedule#field}.
	Field *string `field:"optional" json:"field" yaml:"field"`
}


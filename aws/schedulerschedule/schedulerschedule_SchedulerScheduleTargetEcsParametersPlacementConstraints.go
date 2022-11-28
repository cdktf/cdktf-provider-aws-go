package schedulerschedule


type SchedulerScheduleTargetEcsParametersPlacementConstraints struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/scheduler_schedule#type SchedulerSchedule#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/scheduler_schedule#expression SchedulerSchedule#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}


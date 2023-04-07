package schedulerschedule


type SchedulerScheduleTargetDeadLetterConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/scheduler_schedule#arn SchedulerSchedule#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
}


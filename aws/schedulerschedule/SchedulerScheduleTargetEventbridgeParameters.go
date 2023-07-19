package schedulerschedule


type SchedulerScheduleTargetEventbridgeParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/scheduler_schedule#detail_type SchedulerSchedule#detail_type}.
	DetailType *string `field:"required" json:"detailType" yaml:"detailType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/scheduler_schedule#source SchedulerSchedule#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
}


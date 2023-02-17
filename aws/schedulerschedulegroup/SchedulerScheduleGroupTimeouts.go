package schedulerschedulegroup


type SchedulerScheduleGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/scheduler_schedule_group#create SchedulerScheduleGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/scheduler_schedule_group#delete SchedulerScheduleGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}


package schedulerschedule


type SchedulerScheduleTargetEcsParametersCapacityProviderStrategy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/scheduler_schedule#capacity_provider SchedulerSchedule#capacity_provider}.
	CapacityProvider *string `field:"required" json:"capacityProvider" yaml:"capacityProvider"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/scheduler_schedule#base SchedulerSchedule#base}.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/scheduler_schedule#weight SchedulerSchedule#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}


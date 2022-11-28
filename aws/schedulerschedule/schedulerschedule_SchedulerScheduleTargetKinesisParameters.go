package schedulerschedule


type SchedulerScheduleTargetKinesisParameters struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/scheduler_schedule#partition_key SchedulerSchedule#partition_key}.
	PartitionKey *string `field:"required" json:"partitionKey" yaml:"partitionKey"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetKinesisParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/scheduler_schedule#partition_key SchedulerSchedule#partition_key}.
	PartitionKey *string `field:"required" json:"partitionKey" yaml:"partitionKey"`
}


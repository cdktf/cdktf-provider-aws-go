// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTargetSagemakerPipelineParameters struct {
	// pipeline_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/scheduler_schedule#pipeline_parameter SchedulerSchedule#pipeline_parameter}
	PipelineParameter interface{} `field:"optional" json:"pipelineParameter" yaml:"pipelineParameter"`
}


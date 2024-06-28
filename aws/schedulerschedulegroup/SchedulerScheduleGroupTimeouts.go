// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schedulerschedulegroup


type SchedulerScheduleGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/scheduler_schedule_group#create SchedulerScheduleGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/scheduler_schedule_group#delete SchedulerScheduleGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}


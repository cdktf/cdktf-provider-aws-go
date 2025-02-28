// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowTriggerConfigTriggerPropertiesScheduled struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/appflow_flow#schedule_expression AppflowFlow#schedule_expression}.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/appflow_flow#data_pull_mode AppflowFlow#data_pull_mode}.
	DataPullMode *string `field:"optional" json:"dataPullMode" yaml:"dataPullMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/appflow_flow#first_execution_from AppflowFlow#first_execution_from}.
	FirstExecutionFrom *string `field:"optional" json:"firstExecutionFrom" yaml:"firstExecutionFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/appflow_flow#schedule_end_time AppflowFlow#schedule_end_time}.
	ScheduleEndTime *string `field:"optional" json:"scheduleEndTime" yaml:"scheduleEndTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/appflow_flow#schedule_offset AppflowFlow#schedule_offset}.
	ScheduleOffset *float64 `field:"optional" json:"scheduleOffset" yaml:"scheduleOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/appflow_flow#schedule_start_time AppflowFlow#schedule_start_time}.
	ScheduleStartTime *string `field:"optional" json:"scheduleStartTime" yaml:"scheduleStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/appflow_flow#timezone AppflowFlow#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
}


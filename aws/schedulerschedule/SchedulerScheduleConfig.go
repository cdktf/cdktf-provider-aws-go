// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SchedulerScheduleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// flexible_time_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#flexible_time_window SchedulerSchedule#flexible_time_window}
	FlexibleTimeWindow *SchedulerScheduleFlexibleTimeWindow `field:"required" json:"flexibleTimeWindow" yaml:"flexibleTimeWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#schedule_expression SchedulerSchedule#schedule_expression}.
	ScheduleExpression *string `field:"required" json:"scheduleExpression" yaml:"scheduleExpression"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#target SchedulerSchedule#target}
	Target *SchedulerScheduleTarget `field:"required" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#action_after_completion SchedulerSchedule#action_after_completion}.
	ActionAfterCompletion *string `field:"optional" json:"actionAfterCompletion" yaml:"actionAfterCompletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#description SchedulerSchedule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#end_date SchedulerSchedule#end_date}.
	EndDate *string `field:"optional" json:"endDate" yaml:"endDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#group_name SchedulerSchedule#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#id SchedulerSchedule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#kms_key_arn SchedulerSchedule#kms_key_arn}.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#name SchedulerSchedule#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#name_prefix SchedulerSchedule#name_prefix}.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#region SchedulerSchedule#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#schedule_expression_timezone SchedulerSchedule#schedule_expression_timezone}.
	ScheduleExpressionTimezone *string `field:"optional" json:"scheduleExpressionTimezone" yaml:"scheduleExpressionTimezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#start_date SchedulerSchedule#start_date}.
	StartDate *string `field:"optional" json:"startDate" yaml:"startDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/scheduler_schedule#state SchedulerSchedule#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}


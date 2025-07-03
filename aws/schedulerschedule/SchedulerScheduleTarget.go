// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package schedulerschedule


type SchedulerScheduleTarget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/scheduler_schedule#arn SchedulerSchedule#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/scheduler_schedule#role_arn SchedulerSchedule#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/scheduler_schedule#dead_letter_config SchedulerSchedule#dead_letter_config}
	DeadLetterConfig *SchedulerScheduleTargetDeadLetterConfig `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// ecs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/scheduler_schedule#ecs_parameters SchedulerSchedule#ecs_parameters}
	EcsParameters *SchedulerScheduleTargetEcsParameters `field:"optional" json:"ecsParameters" yaml:"ecsParameters"`
	// eventbridge_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/scheduler_schedule#eventbridge_parameters SchedulerSchedule#eventbridge_parameters}
	EventbridgeParameters *SchedulerScheduleTargetEventbridgeParameters `field:"optional" json:"eventbridgeParameters" yaml:"eventbridgeParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/scheduler_schedule#input SchedulerSchedule#input}.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// kinesis_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/scheduler_schedule#kinesis_parameters SchedulerSchedule#kinesis_parameters}
	KinesisParameters *SchedulerScheduleTargetKinesisParameters `field:"optional" json:"kinesisParameters" yaml:"kinesisParameters"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/scheduler_schedule#retry_policy SchedulerSchedule#retry_policy}
	RetryPolicy *SchedulerScheduleTargetRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// sagemaker_pipeline_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/scheduler_schedule#sagemaker_pipeline_parameters SchedulerSchedule#sagemaker_pipeline_parameters}
	SagemakerPipelineParameters *SchedulerScheduleTargetSagemakerPipelineParameters `field:"optional" json:"sagemakerPipelineParameters" yaml:"sagemakerPipelineParameters"`
	// sqs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/scheduler_schedule#sqs_parameters SchedulerSchedule#sqs_parameters}
	SqsParameters *SchedulerScheduleTargetSqsParameters `field:"optional" json:"sqsParameters" yaml:"sqsParameters"`
}


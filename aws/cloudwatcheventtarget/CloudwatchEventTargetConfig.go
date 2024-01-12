// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatcheventtarget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudwatchEventTargetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#arn CloudwatchEventTarget#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#rule CloudwatchEventTarget#rule}.
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// batch_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#batch_target CloudwatchEventTarget#batch_target}
	BatchTarget *CloudwatchEventTargetBatchTarget `field:"optional" json:"batchTarget" yaml:"batchTarget"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#dead_letter_config CloudwatchEventTarget#dead_letter_config}
	DeadLetterConfig *CloudwatchEventTargetDeadLetterConfig `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// ecs_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#ecs_target CloudwatchEventTarget#ecs_target}
	EcsTarget *CloudwatchEventTargetEcsTarget `field:"optional" json:"ecsTarget" yaml:"ecsTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#event_bus_name CloudwatchEventTarget#event_bus_name}.
	EventBusName *string `field:"optional" json:"eventBusName" yaml:"eventBusName"`
	// http_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#http_target CloudwatchEventTarget#http_target}
	HttpTarget *CloudwatchEventTargetHttpTarget `field:"optional" json:"httpTarget" yaml:"httpTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#id CloudwatchEventTarget#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#input CloudwatchEventTarget#input}.
	Input *string `field:"optional" json:"input" yaml:"input"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#input_path CloudwatchEventTarget#input_path}.
	InputPath *string `field:"optional" json:"inputPath" yaml:"inputPath"`
	// input_transformer block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#input_transformer CloudwatchEventTarget#input_transformer}
	InputTransformer *CloudwatchEventTargetInputTransformer `field:"optional" json:"inputTransformer" yaml:"inputTransformer"`
	// kinesis_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#kinesis_target CloudwatchEventTarget#kinesis_target}
	KinesisTarget *CloudwatchEventTargetKinesisTarget `field:"optional" json:"kinesisTarget" yaml:"kinesisTarget"`
	// redshift_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#redshift_target CloudwatchEventTarget#redshift_target}
	RedshiftTarget *CloudwatchEventTargetRedshiftTarget `field:"optional" json:"redshiftTarget" yaml:"redshiftTarget"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#retry_policy CloudwatchEventTarget#retry_policy}
	RetryPolicy *CloudwatchEventTargetRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#role_arn CloudwatchEventTarget#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// run_command_targets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#run_command_targets CloudwatchEventTarget#run_command_targets}
	RunCommandTargets interface{} `field:"optional" json:"runCommandTargets" yaml:"runCommandTargets"`
	// sagemaker_pipeline_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#sagemaker_pipeline_target CloudwatchEventTarget#sagemaker_pipeline_target}
	SagemakerPipelineTarget *CloudwatchEventTargetSagemakerPipelineTarget `field:"optional" json:"sagemakerPipelineTarget" yaml:"sagemakerPipelineTarget"`
	// sqs_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#sqs_target CloudwatchEventTarget#sqs_target}
	SqsTarget *CloudwatchEventTargetSqsTarget `field:"optional" json:"sqsTarget" yaml:"sqsTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/cloudwatch_event_target#target_id CloudwatchEventTarget#target_id}.
	TargetId *string `field:"optional" json:"targetId" yaml:"targetId"`
}


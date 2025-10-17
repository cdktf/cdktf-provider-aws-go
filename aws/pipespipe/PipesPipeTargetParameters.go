// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParameters struct {
	// batch_job_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#batch_job_parameters PipesPipe#batch_job_parameters}
	BatchJobParameters *PipesPipeTargetParametersBatchJobParameters `field:"optional" json:"batchJobParameters" yaml:"batchJobParameters"`
	// cloudwatch_logs_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#cloudwatch_logs_parameters PipesPipe#cloudwatch_logs_parameters}
	CloudwatchLogsParameters *PipesPipeTargetParametersCloudwatchLogsParameters `field:"optional" json:"cloudwatchLogsParameters" yaml:"cloudwatchLogsParameters"`
	// ecs_task_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#ecs_task_parameters PipesPipe#ecs_task_parameters}
	EcsTaskParameters *PipesPipeTargetParametersEcsTaskParameters `field:"optional" json:"ecsTaskParameters" yaml:"ecsTaskParameters"`
	// eventbridge_event_bus_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#eventbridge_event_bus_parameters PipesPipe#eventbridge_event_bus_parameters}
	EventbridgeEventBusParameters *PipesPipeTargetParametersEventbridgeEventBusParameters `field:"optional" json:"eventbridgeEventBusParameters" yaml:"eventbridgeEventBusParameters"`
	// http_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#http_parameters PipesPipe#http_parameters}
	HttpParameters *PipesPipeTargetParametersHttpParameters `field:"optional" json:"httpParameters" yaml:"httpParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#input_template PipesPipe#input_template}.
	InputTemplate *string `field:"optional" json:"inputTemplate" yaml:"inputTemplate"`
	// kinesis_stream_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#kinesis_stream_parameters PipesPipe#kinesis_stream_parameters}
	KinesisStreamParameters *PipesPipeTargetParametersKinesisStreamParameters `field:"optional" json:"kinesisStreamParameters" yaml:"kinesisStreamParameters"`
	// lambda_function_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#lambda_function_parameters PipesPipe#lambda_function_parameters}
	LambdaFunctionParameters *PipesPipeTargetParametersLambdaFunctionParameters `field:"optional" json:"lambdaFunctionParameters" yaml:"lambdaFunctionParameters"`
	// redshift_data_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#redshift_data_parameters PipesPipe#redshift_data_parameters}
	RedshiftDataParameters *PipesPipeTargetParametersRedshiftDataParameters `field:"optional" json:"redshiftDataParameters" yaml:"redshiftDataParameters"`
	// sagemaker_pipeline_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#sagemaker_pipeline_parameters PipesPipe#sagemaker_pipeline_parameters}
	SagemakerPipelineParameters *PipesPipeTargetParametersSagemakerPipelineParameters `field:"optional" json:"sagemakerPipelineParameters" yaml:"sagemakerPipelineParameters"`
	// sqs_queue_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#sqs_queue_parameters PipesPipe#sqs_queue_parameters}
	SqsQueueParameters *PipesPipeTargetParametersSqsQueueParameters `field:"optional" json:"sqsQueueParameters" yaml:"sqsQueueParameters"`
	// step_function_state_machine_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/pipes_pipe#step_function_state_machine_parameters PipesPipe#step_function_state_machine_parameters}
	StepFunctionStateMachineParameters *PipesPipeTargetParametersStepFunctionStateMachineParameters `field:"optional" json:"stepFunctionStateMachineParameters" yaml:"stepFunctionStateMachineParameters"`
}


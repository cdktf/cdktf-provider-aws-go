// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/pipes_pipe#level PipesPipe#level}.
	Level *string `field:"required" json:"level" yaml:"level"`
	// cloudwatch_logs_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/pipes_pipe#cloudwatch_logs_log_destination PipesPipe#cloudwatch_logs_log_destination}
	CloudwatchLogsLogDestination *PipesPipeLogConfigurationCloudwatchLogsLogDestination `field:"optional" json:"cloudwatchLogsLogDestination" yaml:"cloudwatchLogsLogDestination"`
	// firehose_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/pipes_pipe#firehose_log_destination PipesPipe#firehose_log_destination}
	FirehoseLogDestination *PipesPipeLogConfigurationFirehoseLogDestination `field:"optional" json:"firehoseLogDestination" yaml:"firehoseLogDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/pipes_pipe#include_execution_data PipesPipe#include_execution_data}.
	IncludeExecutionData *[]*string `field:"optional" json:"includeExecutionData" yaml:"includeExecutionData"`
	// s3_log_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/pipes_pipe#s3_log_destination PipesPipe#s3_log_destination}
	S3LogDestination *PipesPipeLogConfigurationS3LogDestination `field:"optional" json:"s3LogDestination" yaml:"s3LogDestination"`
}


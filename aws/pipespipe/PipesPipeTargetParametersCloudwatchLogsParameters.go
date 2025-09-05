// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParametersCloudwatchLogsParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/pipes_pipe#log_stream_name PipesPipe#log_stream_name}.
	LogStreamName *string `field:"optional" json:"logStreamName" yaml:"logStreamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/pipes_pipe#timestamp PipesPipe#timestamp}.
	Timestamp *string `field:"optional" json:"timestamp" yaml:"timestamp"`
}


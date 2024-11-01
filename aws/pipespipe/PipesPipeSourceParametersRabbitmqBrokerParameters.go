// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeSourceParametersRabbitmqBrokerParameters struct {
	// credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/pipes_pipe#credentials PipesPipe#credentials}
	Credentials *PipesPipeSourceParametersRabbitmqBrokerParametersCredentials `field:"required" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/pipes_pipe#queue_name PipesPipe#queue_name}.
	QueueName *string `field:"required" json:"queueName" yaml:"queueName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/pipes_pipe#batch_size PipesPipe#batch_size}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/pipes_pipe#maximum_batching_window_in_seconds PipesPipe#maximum_batching_window_in_seconds}.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/pipes_pipe#virtual_host PipesPipe#virtual_host}.
	VirtualHost *string `field:"optional" json:"virtualHost" yaml:"virtualHost"`
}


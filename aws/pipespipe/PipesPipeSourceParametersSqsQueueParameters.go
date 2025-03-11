// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeSourceParametersSqsQueueParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/pipes_pipe#batch_size PipesPipe#batch_size}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/pipes_pipe#maximum_batching_window_in_seconds PipesPipe#maximum_batching_window_in_seconds}.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
}


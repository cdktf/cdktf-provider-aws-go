// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeSourceParametersKinesisStreamParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/pipes_pipe#starting_position PipesPipe#starting_position}.
	StartingPosition *string `field:"required" json:"startingPosition" yaml:"startingPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/pipes_pipe#batch_size PipesPipe#batch_size}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// dead_letter_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/pipes_pipe#dead_letter_config PipesPipe#dead_letter_config}
	DeadLetterConfig *PipesPipeSourceParametersKinesisStreamParametersDeadLetterConfig `field:"optional" json:"deadLetterConfig" yaml:"deadLetterConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/pipes_pipe#maximum_batching_window_in_seconds PipesPipe#maximum_batching_window_in_seconds}.
	MaximumBatchingWindowInSeconds *float64 `field:"optional" json:"maximumBatchingWindowInSeconds" yaml:"maximumBatchingWindowInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/pipes_pipe#maximum_record_age_in_seconds PipesPipe#maximum_record_age_in_seconds}.
	MaximumRecordAgeInSeconds *float64 `field:"optional" json:"maximumRecordAgeInSeconds" yaml:"maximumRecordAgeInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/pipes_pipe#maximum_retry_attempts PipesPipe#maximum_retry_attempts}.
	MaximumRetryAttempts *float64 `field:"optional" json:"maximumRetryAttempts" yaml:"maximumRetryAttempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/pipes_pipe#on_partial_batch_item_failure PipesPipe#on_partial_batch_item_failure}.
	OnPartialBatchItemFailure *string `field:"optional" json:"onPartialBatchItemFailure" yaml:"onPartialBatchItemFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/pipes_pipe#parallelization_factor PipesPipe#parallelization_factor}.
	ParallelizationFactor *float64 `field:"optional" json:"parallelizationFactor" yaml:"parallelizationFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/pipes_pipe#starting_position_timestamp PipesPipe#starting_position_timestamp}.
	StartingPositionTimestamp *string `field:"optional" json:"startingPositionTimestamp" yaml:"startingPositionTimestamp"`
}


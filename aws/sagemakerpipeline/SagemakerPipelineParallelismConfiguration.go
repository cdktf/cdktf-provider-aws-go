// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerpipeline


type SagemakerPipelineParallelismConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/sagemaker_pipeline#max_parallel_execution_steps SagemakerPipeline#max_parallel_execution_steps}.
	MaxParallelExecutionSteps *float64 `field:"required" json:"maxParallelExecutionSteps" yaml:"maxParallelExecutionSteps"`
}


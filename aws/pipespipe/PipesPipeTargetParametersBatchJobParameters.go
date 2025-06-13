// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParametersBatchJobParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/pipes_pipe#job_definition PipesPipe#job_definition}.
	JobDefinition *string `field:"required" json:"jobDefinition" yaml:"jobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/pipes_pipe#job_name PipesPipe#job_name}.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// array_properties block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/pipes_pipe#array_properties PipesPipe#array_properties}
	ArrayProperties *PipesPipeTargetParametersBatchJobParametersArrayProperties `field:"optional" json:"arrayProperties" yaml:"arrayProperties"`
	// container_overrides block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/pipes_pipe#container_overrides PipesPipe#container_overrides}
	ContainerOverrides *PipesPipeTargetParametersBatchJobParametersContainerOverrides `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// depends_on block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/pipes_pipe#depends_on PipesPipe#depends_on}
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/pipes_pipe#parameters PipesPipe#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// retry_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.100.0/docs/resources/pipes_pipe#retry_strategy PipesPipe#retry_strategy}
	RetryStrategy *PipesPipeTargetParametersBatchJobParametersRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
}


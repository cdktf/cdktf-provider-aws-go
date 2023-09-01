// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParametersEcsTaskParametersOverrides struct {
	// container_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/pipes_pipe#container_override PipesPipe#container_override}
	ContainerOverride interface{} `field:"optional" json:"containerOverride" yaml:"containerOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/pipes_pipe#cpu PipesPipe#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// ephemeral_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/pipes_pipe#ephemeral_storage PipesPipe#ephemeral_storage}
	EphemeralStorage *PipesPipeTargetParametersEcsTaskParametersOverridesEphemeralStorage `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/pipes_pipe#execution_role_arn PipesPipe#execution_role_arn}.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// inference_accelerator_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/pipes_pipe#inference_accelerator_override PipesPipe#inference_accelerator_override}
	InferenceAcceleratorOverride interface{} `field:"optional" json:"inferenceAcceleratorOverride" yaml:"inferenceAcceleratorOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/pipes_pipe#memory PipesPipe#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/pipes_pipe#task_role_arn PipesPipe#task_role_arn}.
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
}


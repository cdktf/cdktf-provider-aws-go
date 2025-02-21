// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParametersBatchJobParametersContainerOverrides struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/pipes_pipe#command PipesPipe#command}.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/pipes_pipe#environment PipesPipe#environment}
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/pipes_pipe#instance_type PipesPipe#instance_type}.
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// resource_requirement block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/pipes_pipe#resource_requirement PipesPipe#resource_requirement}
	ResourceRequirement interface{} `field:"optional" json:"resourceRequirement" yaml:"resourceRequirement"`
}


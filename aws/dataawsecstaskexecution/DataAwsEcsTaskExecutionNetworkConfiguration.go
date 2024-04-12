// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsecstaskexecution


type DataAwsEcsTaskExecutionNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/ecs_task_execution#subnets DataAwsEcsTaskExecution#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/ecs_task_execution#assign_public_ip DataAwsEcsTaskExecution#assign_public_ip}.
	AssignPublicIp interface{} `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/ecs_task_execution#security_groups DataAwsEcsTaskExecution#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}


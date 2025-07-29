// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsecstaskexecution


type DataAwsEcsTaskExecutionPlacementStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/data-sources/ecs_task_execution#type DataAwsEcsTaskExecution#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/data-sources/ecs_task_execution#field DataAwsEcsTaskExecution#field}.
	Field *string `field:"optional" json:"field" yaml:"field"`
}


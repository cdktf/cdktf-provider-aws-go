// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskset


type EcsTaskSetServiceRegistries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/ecs_task_set#registry_arn EcsTaskSet#registry_arn}.
	RegistryArn *string `field:"required" json:"registryArn" yaml:"registryArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/ecs_task_set#container_name EcsTaskSet#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/ecs_task_set#container_port EcsTaskSet#container_port}.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/ecs_task_set#port EcsTaskSet#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}


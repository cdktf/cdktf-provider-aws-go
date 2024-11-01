// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsservice


type EcsServiceServiceRegistries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/ecs_service#registry_arn EcsService#registry_arn}.
	RegistryArn *string `field:"required" json:"registryArn" yaml:"registryArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/ecs_service#container_name EcsService#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/ecs_service#container_port EcsService#container_port}.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/ecs_service#port EcsService#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}


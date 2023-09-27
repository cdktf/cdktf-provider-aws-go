// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionProxyConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/ecs_task_definition#container_name EcsTaskDefinition#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/ecs_task_definition#properties EcsTaskDefinition#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/ecs_task_definition#type EcsTaskDefinition#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}


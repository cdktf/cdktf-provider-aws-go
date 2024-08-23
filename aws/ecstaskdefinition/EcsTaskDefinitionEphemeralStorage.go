// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionEphemeralStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/ecs_task_definition#size_in_gib EcsTaskDefinition#size_in_gib}.
	SizeInGib *float64 `field:"required" json:"sizeInGib" yaml:"sizeInGib"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionRuntimePlatform struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ecs_task_definition#cpu_architecture EcsTaskDefinition#cpu_architecture}.
	CpuArchitecture *string `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ecs_task_definition#operating_system_family EcsTaskDefinition#operating_system_family}.
	OperatingSystemFamily *string `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}


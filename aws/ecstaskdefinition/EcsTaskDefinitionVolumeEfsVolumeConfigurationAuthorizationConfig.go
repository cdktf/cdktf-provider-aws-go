// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/ecs_task_definition#access_point_id EcsTaskDefinition#access_point_id}.
	AccessPointId *string `field:"optional" json:"accessPointId" yaml:"accessPointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/ecs_task_definition#iam EcsTaskDefinition#iam}.
	Iam *string `field:"optional" json:"iam" yaml:"iam"`
}


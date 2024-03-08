// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionVolumeEfsVolumeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/ecs_task_definition#file_system_id EcsTaskDefinition#file_system_id}.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// authorization_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/ecs_task_definition#authorization_config EcsTaskDefinition#authorization_config}
	AuthorizationConfig *EcsTaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/ecs_task_definition#root_directory EcsTaskDefinition#root_directory}.
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/ecs_task_definition#transit_encryption EcsTaskDefinition#transit_encryption}.
	TransitEncryption *string `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/ecs_task_definition#transit_encryption_port EcsTaskDefinition#transit_encryption_port}.
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}


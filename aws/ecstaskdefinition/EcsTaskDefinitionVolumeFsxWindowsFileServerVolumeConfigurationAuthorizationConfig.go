// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition


type EcsTaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/ecs_task_definition#credentials_parameter EcsTaskDefinition#credentials_parameter}.
	CredentialsParameter *string `field:"required" json:"credentialsParameter" yaml:"credentialsParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.58.0/docs/resources/ecs_task_definition#domain EcsTaskDefinition#domain}.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
}


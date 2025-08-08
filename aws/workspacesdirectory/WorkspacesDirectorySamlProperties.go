// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectorySamlProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/workspaces_directory#relay_state_parameter_name WorkspacesDirectory#relay_state_parameter_name}.
	RelayStateParameterName *string `field:"optional" json:"relayStateParameterName" yaml:"relayStateParameterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/workspaces_directory#status WorkspacesDirectory#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/workspaces_directory#user_access_url WorkspacesDirectory#user_access_url}.
	UserAccessUrl *string `field:"optional" json:"userAccessUrl" yaml:"userAccessUrl"`
}


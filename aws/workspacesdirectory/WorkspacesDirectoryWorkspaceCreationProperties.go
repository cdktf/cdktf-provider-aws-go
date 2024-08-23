// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryWorkspaceCreationProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/workspaces_directory#custom_security_group_id WorkspacesDirectory#custom_security_group_id}.
	CustomSecurityGroupId *string `field:"optional" json:"customSecurityGroupId" yaml:"customSecurityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/workspaces_directory#default_ou WorkspacesDirectory#default_ou}.
	DefaultOu *string `field:"optional" json:"defaultOu" yaml:"defaultOu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/workspaces_directory#enable_internet_access WorkspacesDirectory#enable_internet_access}.
	EnableInternetAccess interface{} `field:"optional" json:"enableInternetAccess" yaml:"enableInternetAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/workspaces_directory#enable_maintenance_mode WorkspacesDirectory#enable_maintenance_mode}.
	EnableMaintenanceMode interface{} `field:"optional" json:"enableMaintenanceMode" yaml:"enableMaintenanceMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/workspaces_directory#user_enabled_as_local_administrator WorkspacesDirectory#user_enabled_as_local_administrator}.
	UserEnabledAsLocalAdministrator interface{} `field:"optional" json:"userEnabledAsLocalAdministrator" yaml:"userEnabledAsLocalAdministrator"`
}


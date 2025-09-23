// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebusersettings


type WorkspaceswebUserSettingsCookieSynchronizationConfigurationAllowlistStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/workspacesweb_user_settings#domain WorkspaceswebUserSettings#domain}.
	Domain *string `field:"required" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/workspacesweb_user_settings#name WorkspaceswebUserSettings#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/workspacesweb_user_settings#path WorkspaceswebUserSettings#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}


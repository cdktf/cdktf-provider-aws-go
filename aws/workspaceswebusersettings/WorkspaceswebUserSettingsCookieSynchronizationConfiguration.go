// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebusersettings


type WorkspaceswebUserSettingsCookieSynchronizationConfiguration struct {
	// allowlist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/workspacesweb_user_settings#allowlist WorkspaceswebUserSettings#allowlist}
	Allowlist interface{} `field:"optional" json:"allowlist" yaml:"allowlist"`
	// blocklist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/workspacesweb_user_settings#blocklist WorkspaceswebUserSettings#blocklist}
	Blocklist interface{} `field:"optional" json:"blocklist" yaml:"blocklist"`
}


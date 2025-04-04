// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesipgroup


type WorkspacesIpGroupRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/workspaces_ip_group#source WorkspacesIpGroup#source}.
	Source *string `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/workspaces_ip_group#description WorkspacesIpGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}


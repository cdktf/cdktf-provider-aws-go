// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebsessionlogger


type WorkspaceswebSessionLoggerEventFilter struct {
	// all block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/workspacesweb_session_logger#all WorkspaceswebSessionLogger#all}
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/workspacesweb_session_logger#include WorkspaceswebSessionLogger#include}.
	Include *[]*string `field:"optional" json:"include" yaml:"include"`
}


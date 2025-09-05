// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebsessionlogger


type WorkspaceswebSessionLoggerLogConfiguration struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/workspacesweb_session_logger#s3 WorkspaceswebSessionLogger#s3}
	S3 interface{} `field:"optional" json:"s3" yaml:"s3"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebsessionlogger


type WorkspaceswebSessionLoggerLogConfigurationS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/workspacesweb_session_logger#bucket WorkspaceswebSessionLogger#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/workspacesweb_session_logger#folder_structure WorkspaceswebSessionLogger#folder_structure}.
	FolderStructure *string `field:"required" json:"folderStructure" yaml:"folderStructure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/workspacesweb_session_logger#log_file_format WorkspaceswebSessionLogger#log_file_format}.
	LogFileFormat *string `field:"required" json:"logFileFormat" yaml:"logFileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/workspacesweb_session_logger#bucket_owner WorkspaceswebSessionLogger#bucket_owner}.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/workspacesweb_session_logger#key_prefix WorkspaceswebSessionLogger#key_prefix}.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}


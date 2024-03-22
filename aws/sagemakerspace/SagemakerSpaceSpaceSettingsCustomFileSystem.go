// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace


type SagemakerSpaceSpaceSettingsCustomFileSystem struct {
	// efs_file_system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/sagemaker_space#efs_file_system SagemakerSpace#efs_file_system}
	EfsFileSystem *SagemakerSpaceSpaceSettingsCustomFileSystemEfsFileSystem `field:"required" json:"efsFileSystem" yaml:"efsFileSystem"`
}


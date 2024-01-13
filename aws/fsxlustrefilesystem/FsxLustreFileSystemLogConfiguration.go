// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxlustrefilesystem


type FsxLustreFileSystemLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/fsx_lustre_file_system#destination FsxLustreFileSystem#destination}.
	Destination *string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/fsx_lustre_file_system#level FsxLustreFileSystem#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
}


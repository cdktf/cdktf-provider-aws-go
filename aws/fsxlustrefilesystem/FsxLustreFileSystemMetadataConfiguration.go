// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxlustrefilesystem


type FsxLustreFileSystemMetadataConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/fsx_lustre_file_system#iops FsxLustreFileSystem#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/fsx_lustre_file_system#mode FsxLustreFileSystem#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}


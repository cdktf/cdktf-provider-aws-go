// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxlustrefilesystem


type FsxLustreFileSystemDataReadCacheConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/fsx_lustre_file_system#sizing_mode FsxLustreFileSystem#sizing_mode}.
	SizingMode *string `field:"required" json:"sizingMode" yaml:"sizingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/fsx_lustre_file_system#size FsxLustreFileSystem#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}


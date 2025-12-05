// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxopenzfsfilesystem


type FsxOpenzfsFileSystemReadCacheConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/fsx_openzfs_file_system#size FsxOpenzfsFileSystem#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/fsx_openzfs_file_system#sizing_mode FsxOpenzfsFileSystem#sizing_mode}.
	SizingMode *string `field:"optional" json:"sizingMode" yaml:"sizingMode"`
}


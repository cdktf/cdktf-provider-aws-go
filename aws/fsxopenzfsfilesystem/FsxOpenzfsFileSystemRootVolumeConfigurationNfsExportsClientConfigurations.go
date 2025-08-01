// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxopenzfsfilesystem


type FsxOpenzfsFileSystemRootVolumeConfigurationNfsExportsClientConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#clients FsxOpenzfsFileSystem#clients}.
	Clients *string `field:"required" json:"clients" yaml:"clients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_openzfs_file_system#options FsxOpenzfsFileSystem#options}.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
}


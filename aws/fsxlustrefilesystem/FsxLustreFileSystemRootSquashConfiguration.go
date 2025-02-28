// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxlustrefilesystem


type FsxLustreFileSystemRootSquashConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/fsx_lustre_file_system#no_squash_nids FsxLustreFileSystem#no_squash_nids}.
	NoSquashNids *[]*string `field:"optional" json:"noSquashNids" yaml:"noSquashNids"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/fsx_lustre_file_system#root_squash FsxLustreFileSystem#root_squash}.
	RootSquash *string `field:"optional" json:"rootSquash" yaml:"rootSquash"`
}


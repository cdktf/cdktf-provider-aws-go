// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynclocationfsxopenzfsfilesystem


type DatasyncLocationFsxOpenzfsFileSystemProtocol struct {
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/datasync_location_fsx_openzfs_file_system#nfs DatasyncLocationFsxOpenzfsFileSystem#nfs}
	Nfs *DatasyncLocationFsxOpenzfsFileSystemProtocolNfs `field:"required" json:"nfs" yaml:"nfs"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynclocationfsxontapfilesystem


type DatasyncLocationFsxOntapFileSystemProtocolNfs struct {
	// mount_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/datasync_location_fsx_ontap_file_system#mount_options DatasyncLocationFsxOntapFileSystem#mount_options}
	MountOptions *DatasyncLocationFsxOntapFileSystemProtocolNfsMountOptions `field:"required" json:"mountOptions" yaml:"mountOptions"`
}


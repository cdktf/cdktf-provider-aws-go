// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxopenzfsfilesystem


type FsxOpenzfsFileSystemRootVolumeConfigurationNfsExports struct {
	// client_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/fsx_openzfs_file_system#client_configurations FsxOpenzfsFileSystem#client_configurations}
	ClientConfigurations interface{} `field:"required" json:"clientConfigurations" yaml:"clientConfigurations"`
}


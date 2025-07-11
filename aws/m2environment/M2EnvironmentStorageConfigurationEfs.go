// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package m2environment


type M2EnvironmentStorageConfigurationEfs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/m2_environment#file_system_id M2Environment#file_system_id}.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/m2_environment#mount_point M2Environment#mount_point}.
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
}


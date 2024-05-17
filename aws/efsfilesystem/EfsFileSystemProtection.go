// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package efsfilesystem


type EfsFileSystemProtection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.50.0/docs/resources/efs_file_system#replication_overwrite EfsFileSystem#replication_overwrite}.
	ReplicationOverwrite *string `field:"optional" json:"replicationOverwrite" yaml:"replicationOverwrite"`
}


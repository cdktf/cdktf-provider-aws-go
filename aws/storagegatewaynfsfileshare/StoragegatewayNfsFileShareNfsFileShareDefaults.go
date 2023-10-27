// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewaynfsfileshare


type StoragegatewayNfsFileShareNfsFileShareDefaults struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/storagegateway_nfs_file_share#directory_mode StoragegatewayNfsFileShare#directory_mode}.
	DirectoryMode *string `field:"optional" json:"directoryMode" yaml:"directoryMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/storagegateway_nfs_file_share#file_mode StoragegatewayNfsFileShare#file_mode}.
	FileMode *string `field:"optional" json:"fileMode" yaml:"fileMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/storagegateway_nfs_file_share#group_id StoragegatewayNfsFileShare#group_id}.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/storagegateway_nfs_file_share#owner_id StoragegatewayNfsFileShare#owner_id}.
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
}


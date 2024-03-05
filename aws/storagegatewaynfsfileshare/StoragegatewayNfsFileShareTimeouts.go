// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewaynfsfileshare


type StoragegatewayNfsFileShareTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/storagegateway_nfs_file_share#create StoragegatewayNfsFileShare#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/storagegateway_nfs_file_share#delete StoragegatewayNfsFileShare#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/storagegateway_nfs_file_share#update StoragegatewayNfsFileShare#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


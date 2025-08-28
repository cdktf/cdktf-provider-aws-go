// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewayfilesystemassociation


type StoragegatewayFileSystemAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/storagegateway_file_system_association#create StoragegatewayFileSystemAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/storagegateway_file_system_association#delete StoragegatewayFileSystemAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/storagegateway_file_system_association#update StoragegatewayFileSystemAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


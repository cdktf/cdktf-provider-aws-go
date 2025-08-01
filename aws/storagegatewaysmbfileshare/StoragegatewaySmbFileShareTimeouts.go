// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewaysmbfileshare


type StoragegatewaySmbFileShareTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/storagegateway_smb_file_share#create StoragegatewaySmbFileShare#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/storagegateway_smb_file_share#delete StoragegatewaySmbFileShare#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/storagegateway_smb_file_share#update StoragegatewaySmbFileShare#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


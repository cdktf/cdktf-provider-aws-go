// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynclocationazureblob


type DatasyncLocationAzureBlobSasConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/datasync_location_azure_blob#token DatasyncLocationAzureBlob#token}.
	Token *string `field:"required" json:"token" yaml:"token"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datasynclocationazureblob


type DatasyncLocationAzureBlobSasConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/datasync_location_azure_blob#token DatasyncLocationAzureBlob#token}.
	Token *string `field:"required" json:"token" yaml:"token"`
}


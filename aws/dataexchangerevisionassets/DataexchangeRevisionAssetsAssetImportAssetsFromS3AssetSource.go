// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataexchangerevisionassets


type DataexchangeRevisionAssetsAssetImportAssetsFromS3AssetSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/dataexchange_revision_assets#bucket DataexchangeRevisionAssets#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/dataexchange_revision_assets#key DataexchangeRevisionAssets#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataexchangerevisionassets


type DataexchangeRevisionAssetsAsset struct {
	// create_s3_data_access_from_s3_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/dataexchange_revision_assets#create_s3_data_access_from_s3_bucket DataexchangeRevisionAssets#create_s3_data_access_from_s3_bucket}
	CreateS3DataAccessFromS3Bucket interface{} `field:"optional" json:"createS3DataAccessFromS3Bucket" yaml:"createS3DataAccessFromS3Bucket"`
	// import_assets_from_s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/dataexchange_revision_assets#import_assets_from_s3 DataexchangeRevisionAssets#import_assets_from_s3}
	ImportAssetsFromS3 interface{} `field:"optional" json:"importAssetsFromS3" yaml:"importAssetsFromS3"`
	// import_assets_from_signed_url block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/dataexchange_revision_assets#import_assets_from_signed_url DataexchangeRevisionAssets#import_assets_from_signed_url}
	ImportAssetsFromSignedUrl interface{} `field:"optional" json:"importAssetsFromSignedUrl" yaml:"importAssetsFromSignedUrl"`
}


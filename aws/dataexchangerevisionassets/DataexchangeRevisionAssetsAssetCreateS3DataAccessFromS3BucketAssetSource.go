// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataexchangerevisionassets


type DataexchangeRevisionAssetsAssetCreateS3DataAccessFromS3BucketAssetSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/dataexchange_revision_assets#bucket DataexchangeRevisionAssets#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/dataexchange_revision_assets#key_prefixes DataexchangeRevisionAssets#key_prefixes}.
	KeyPrefixes *[]*string `field:"optional" json:"keyPrefixes" yaml:"keyPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/dataexchange_revision_assets#keys DataexchangeRevisionAssets#keys}.
	Keys *[]*string `field:"optional" json:"keys" yaml:"keys"`
	// kms_keys_to_grant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/dataexchange_revision_assets#kms_keys_to_grant DataexchangeRevisionAssets#kms_keys_to_grant}
	KmsKeysToGrant interface{} `field:"optional" json:"kmsKeysToGrant" yaml:"kmsKeysToGrant"`
}


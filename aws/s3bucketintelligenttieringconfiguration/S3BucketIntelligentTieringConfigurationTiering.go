// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketintelligenttieringconfiguration


type S3BucketIntelligentTieringConfigurationTiering struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/s3_bucket_intelligent_tiering_configuration#access_tier S3BucketIntelligentTieringConfiguration#access_tier}.
	AccessTier *string `field:"required" json:"accessTier" yaml:"accessTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/s3_bucket_intelligent_tiering_configuration#days S3BucketIntelligentTieringConfiguration#days}.
	Days *float64 `field:"required" json:"days" yaml:"days"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opensearchpackage


type OpensearchPackagePackageSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/opensearch_package#s3_bucket_name OpensearchPackage#s3_bucket_name}.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/opensearch_package#s3_key OpensearchPackage#s3_key}.
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
}


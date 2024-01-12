// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketpublicaccessblock

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketPublicAccessBlockConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/s3_bucket_public_access_block#bucket S3BucketPublicAccessBlock#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/s3_bucket_public_access_block#block_public_acls S3BucketPublicAccessBlock#block_public_acls}.
	BlockPublicAcls interface{} `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/s3_bucket_public_access_block#block_public_policy S3BucketPublicAccessBlock#block_public_policy}.
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/s3_bucket_public_access_block#id S3BucketPublicAccessBlock#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/s3_bucket_public_access_block#ignore_public_acls S3BucketPublicAccessBlock#ignore_public_acls}.
	IgnorePublicAcls interface{} `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/s3_bucket_public_access_block#restrict_public_buckets S3BucketPublicAccessBlock#restrict_public_buckets}.
	RestrictPublicBuckets interface{} `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}


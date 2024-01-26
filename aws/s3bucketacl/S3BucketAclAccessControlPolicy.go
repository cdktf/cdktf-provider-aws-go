// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketacl


type S3BucketAclAccessControlPolicy struct {
	// owner block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/s3_bucket_acl#owner S3BucketAcl#owner}
	Owner *S3BucketAclAccessControlPolicyOwner `field:"required" json:"owner" yaml:"owner"`
	// grant block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/s3_bucket_acl#grant S3BucketAcl#grant}
	Grant interface{} `field:"optional" json:"grant" yaml:"grant"`
}


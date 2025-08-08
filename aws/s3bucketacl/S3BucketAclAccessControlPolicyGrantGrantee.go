// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketacl


type S3BucketAclAccessControlPolicyGrantGrantee struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/s3_bucket_acl#type S3BucketAcl#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/s3_bucket_acl#email_address S3BucketAcl#email_address}.
	EmailAddress *string `field:"optional" json:"emailAddress" yaml:"emailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/s3_bucket_acl#id S3BucketAcl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/s3_bucket_acl#uri S3BucketAcl#uri}.
	Uri *string `field:"optional" json:"uri" yaml:"uri"`
}


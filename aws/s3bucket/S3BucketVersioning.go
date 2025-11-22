// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketVersioning struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/s3_bucket#enabled S3Bucket#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/s3_bucket#mfa_delete S3Bucket#mfa_delete}.
	MfaDelete interface{} `field:"optional" json:"mfaDelete" yaml:"mfaDelete"`
}


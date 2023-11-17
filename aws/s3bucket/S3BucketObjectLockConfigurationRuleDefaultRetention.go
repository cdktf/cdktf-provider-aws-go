// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketObjectLockConfigurationRuleDefaultRetention struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/s3_bucket#mode S3Bucket#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/s3_bucket#days S3Bucket#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/s3_bucket#years S3Bucket#years}.
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}


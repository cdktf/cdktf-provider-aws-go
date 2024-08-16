// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketwebsiteconfiguration


type S3BucketWebsiteConfigurationRoutingRuleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/s3_bucket_website_configuration#http_error_code_returned_equals S3BucketWebsiteConfiguration#http_error_code_returned_equals}.
	HttpErrorCodeReturnedEquals *string `field:"optional" json:"httpErrorCodeReturnedEquals" yaml:"httpErrorCodeReturnedEquals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/s3_bucket_website_configuration#key_prefix_equals S3BucketWebsiteConfiguration#key_prefix_equals}.
	KeyPrefixEquals *string `field:"optional" json:"keyPrefixEquals" yaml:"keyPrefixEquals"`
}


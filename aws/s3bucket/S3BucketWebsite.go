// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucket


type S3BucketWebsite struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/s3_bucket#error_document S3Bucket#error_document}.
	ErrorDocument *string `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/s3_bucket#index_document S3Bucket#index_document}.
	IndexDocument *string `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/s3_bucket#redirect_all_requests_to S3Bucket#redirect_all_requests_to}.
	RedirectAllRequestsTo *string `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/s3_bucket#routing_rules S3Bucket#routing_rules}.
	RoutingRules *string `field:"optional" json:"routingRules" yaml:"routingRules"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketwebsiteconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type S3BucketWebsiteConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_website_configuration#bucket S3BucketWebsiteConfiguration#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// error_document block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_website_configuration#error_document S3BucketWebsiteConfiguration#error_document}
	ErrorDocument *S3BucketWebsiteConfigurationErrorDocument `field:"optional" json:"errorDocument" yaml:"errorDocument"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_website_configuration#expected_bucket_owner S3BucketWebsiteConfiguration#expected_bucket_owner}.
	ExpectedBucketOwner *string `field:"optional" json:"expectedBucketOwner" yaml:"expectedBucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_website_configuration#id S3BucketWebsiteConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// index_document block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_website_configuration#index_document S3BucketWebsiteConfiguration#index_document}
	IndexDocument *S3BucketWebsiteConfigurationIndexDocument `field:"optional" json:"indexDocument" yaml:"indexDocument"`
	// redirect_all_requests_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_website_configuration#redirect_all_requests_to S3BucketWebsiteConfiguration#redirect_all_requests_to}
	RedirectAllRequestsTo *S3BucketWebsiteConfigurationRedirectAllRequestsTo `field:"optional" json:"redirectAllRequestsTo" yaml:"redirectAllRequestsTo"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_website_configuration#region S3BucketWebsiteConfiguration#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// routing_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_website_configuration#routing_rule S3BucketWebsiteConfiguration#routing_rule}
	RoutingRule interface{} `field:"optional" json:"routingRule" yaml:"routingRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/s3_bucket_website_configuration#routing_rules S3BucketWebsiteConfiguration#routing_rules}.
	RoutingRules *string `field:"optional" json:"routingRules" yaml:"routingRules"`
}


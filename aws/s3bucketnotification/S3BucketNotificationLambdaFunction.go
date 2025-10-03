// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketnotification


type S3BucketNotificationLambdaFunction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/s3_bucket_notification#events S3BucketNotification#events}.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/s3_bucket_notification#filter_prefix S3BucketNotification#filter_prefix}.
	FilterPrefix *string `field:"optional" json:"filterPrefix" yaml:"filterPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/s3_bucket_notification#filter_suffix S3BucketNotification#filter_suffix}.
	FilterSuffix *string `field:"optional" json:"filterSuffix" yaml:"filterSuffix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/s3_bucket_notification#id S3BucketNotification#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/s3_bucket_notification#lambda_function_arn S3BucketNotification#lambda_function_arn}.
	LambdaFunctionArn *string `field:"optional" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
}


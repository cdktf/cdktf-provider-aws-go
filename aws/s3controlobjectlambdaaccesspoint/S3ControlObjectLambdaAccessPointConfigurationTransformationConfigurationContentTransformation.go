// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlobjectlambdaaccesspoint


type S3ControlObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformation struct {
	// aws_lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/s3control_object_lambda_access_point#aws_lambda S3ControlObjectLambdaAccessPoint#aws_lambda}
	AwsLambda *S3ControlObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationAwsLambda `field:"required" json:"awsLambda" yaml:"awsLambda"`
}


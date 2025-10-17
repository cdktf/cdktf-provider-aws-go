// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3controlobjectlambdaaccesspoint


type S3ControlObjectLambdaAccessPointConfigurationTransformationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/s3control_object_lambda_access_point#actions S3ControlObjectLambdaAccessPoint#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// content_transformation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/s3control_object_lambda_access_point#content_transformation S3ControlObjectLambdaAccessPoint#content_transformation}
	ContentTransformation *S3ControlObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformation `field:"required" json:"contentTransformation" yaml:"contentTransformation"`
}


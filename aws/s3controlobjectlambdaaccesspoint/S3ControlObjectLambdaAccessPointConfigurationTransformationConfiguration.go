package s3controlobjectlambdaaccesspoint


type S3ControlObjectLambdaAccessPointConfigurationTransformationConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_object_lambda_access_point#actions S3ControlObjectLambdaAccessPoint#actions}.
	Actions *[]*string `field:"required" json:"actions" yaml:"actions"`
	// content_transformation block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_object_lambda_access_point#content_transformation S3ControlObjectLambdaAccessPoint#content_transformation}
	ContentTransformation *S3ControlObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformation `field:"required" json:"contentTransformation" yaml:"contentTransformation"`
}


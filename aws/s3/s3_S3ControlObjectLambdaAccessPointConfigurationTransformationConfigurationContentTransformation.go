package s3


type S3ControlObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformation struct {
	// aws_lambda block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/s3control_object_lambda_access_point#aws_lambda S3ControlObjectLambdaAccessPoint#aws_lambda}
	AwsLambda *S3ControlObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationAwsLambda `field:"required" json:"awsLambda" yaml:"awsLambda"`
}


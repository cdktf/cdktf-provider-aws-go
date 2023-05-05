package s3controlobjectlambdaaccesspoint


type S3ControlObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationAwsLambda struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/s3control_object_lambda_access_point#function_arn S3ControlObjectLambdaAccessPoint#function_arn}.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/s3control_object_lambda_access_point#function_payload S3ControlObjectLambdaAccessPoint#function_payload}.
	FunctionPayload *string `field:"optional" json:"functionPayload" yaml:"functionPayload"`
}


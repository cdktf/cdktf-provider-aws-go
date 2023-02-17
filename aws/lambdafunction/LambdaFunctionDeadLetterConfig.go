package lambdafunction


type LambdaFunctionDeadLetterConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#target_arn LambdaFunction#target_arn}.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
}


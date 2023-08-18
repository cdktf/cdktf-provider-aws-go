package lambdafunction


type LambdaFunctionDeadLetterConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.0/docs/resources/lambda_function#target_arn LambdaFunction#target_arn}.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
}


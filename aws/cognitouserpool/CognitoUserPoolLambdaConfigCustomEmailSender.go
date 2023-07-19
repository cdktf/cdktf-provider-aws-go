package cognitouserpool


type CognitoUserPoolLambdaConfigCustomEmailSender struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/cognito_user_pool#lambda_arn CognitoUserPool#lambda_arn}.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/cognito_user_pool#lambda_version CognitoUserPool#lambda_version}.
	LambdaVersion *string `field:"required" json:"lambdaVersion" yaml:"lambdaVersion"`
}


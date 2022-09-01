package cognito


type CognitoUserPoolLambdaConfigCustomEmailSender struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#lambda_arn CognitoUserPool#lambda_arn}.
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/cognito_user_pool#lambda_version CognitoUserPool#lambda_version}.
	LambdaVersion *string `field:"required" json:"lambdaVersion" yaml:"lambdaVersion"`
}


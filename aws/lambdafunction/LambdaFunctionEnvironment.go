package lambdafunction


type LambdaFunctionEnvironment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#variables LambdaFunction#variables}.
	Variables *map[string]*string `field:"optional" json:"variables" yaml:"variables"`
}


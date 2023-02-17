package lambdafunction


type LambdaFunctionTracingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#mode LambdaFunction#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}


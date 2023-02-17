package lambdafunction


type LambdaFunctionEphemeralStorage struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#size LambdaFunction#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}


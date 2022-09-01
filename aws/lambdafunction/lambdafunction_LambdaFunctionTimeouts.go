package lambdafunction


type LambdaFunctionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#create LambdaFunction#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}


package lambdafunction


type LambdaFunctionSnapStart struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_function#apply_on LambdaFunction#apply_on}.
	ApplyOn *string `field:"required" json:"applyOn" yaml:"applyOn"`
}


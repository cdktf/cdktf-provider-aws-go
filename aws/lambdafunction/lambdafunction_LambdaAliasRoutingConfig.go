package lambdafunction


type LambdaAliasRoutingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_alias#additional_version_weights LambdaAlias#additional_version_weights}.
	AdditionalVersionWeights *map[string]*float64 `field:"optional" json:"additionalVersionWeights" yaml:"additionalVersionWeights"`
}


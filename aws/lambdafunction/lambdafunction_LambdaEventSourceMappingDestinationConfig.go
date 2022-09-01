package lambdafunction


type LambdaEventSourceMappingDestinationConfig struct {
	// on_failure block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#on_failure LambdaEventSourceMapping#on_failure}
	OnFailure *LambdaEventSourceMappingDestinationConfigOnFailure `field:"optional" json:"onFailure" yaml:"onFailure"`
}


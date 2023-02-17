package lambdaeventsourcemapping


type LambdaEventSourceMappingScalingConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#maximum_concurrency LambdaEventSourceMapping#maximum_concurrency}.
	MaximumConcurrency *float64 `field:"optional" json:"maximumConcurrency" yaml:"maximumConcurrency"`
}


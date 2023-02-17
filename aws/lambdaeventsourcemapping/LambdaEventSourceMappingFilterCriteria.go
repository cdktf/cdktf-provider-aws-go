package lambdaeventsourcemapping


type LambdaEventSourceMappingFilterCriteria struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/lambda_event_source_mapping#filter LambdaEventSourceMapping#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}


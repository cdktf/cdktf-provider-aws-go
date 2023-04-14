package pipespipe


type PipesPipeSourceParametersFilterCriteria struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pipes_pipe#filter PipesPipe#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
}


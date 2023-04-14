package pipespipe


type PipesPipeSourceParameters struct {
	// filter_criteria block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/pipes_pipe#filter_criteria PipesPipe#filter_criteria}
	FilterCriteria *PipesPipeSourceParametersFilterCriteria `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
}


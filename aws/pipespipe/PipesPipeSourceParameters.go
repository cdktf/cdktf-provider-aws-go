package pipespipe


type PipesPipeSourceParameters struct {
	// filter_criteria block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.66.0/docs/resources/pipes_pipe#filter_criteria PipesPipe#filter_criteria}
	FilterCriteria *PipesPipeSourceParametersFilterCriteria `field:"optional" json:"filterCriteria" yaml:"filterCriteria"`
}


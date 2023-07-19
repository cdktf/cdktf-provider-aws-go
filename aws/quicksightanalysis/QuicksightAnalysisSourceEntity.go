package quicksightanalysis


type QuicksightAnalysisSourceEntity struct {
	// source_template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/quicksight_analysis#source_template QuicksightAnalysis#source_template}
	SourceTemplate *QuicksightAnalysisSourceEntitySourceTemplate `field:"optional" json:"sourceTemplate" yaml:"sourceTemplate"`
}


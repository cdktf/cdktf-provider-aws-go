package quicksightanalysis


type QuicksightAnalysisTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/quicksight_analysis#create QuicksightAnalysis#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/quicksight_analysis#delete QuicksightAnalysis#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/quicksight_analysis#update QuicksightAnalysis#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


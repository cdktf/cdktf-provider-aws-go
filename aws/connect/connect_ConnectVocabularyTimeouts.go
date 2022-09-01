package connect


type ConnectVocabularyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/connect_vocabulary#create ConnectVocabulary#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/connect_vocabulary#delete ConnectVocabulary#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}


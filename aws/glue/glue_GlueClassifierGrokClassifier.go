package glue


type GlueClassifierGrokClassifier struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier#classification GlueClassifier#classification}.
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier#grok_pattern GlueClassifier#grok_pattern}.
	GrokPattern *string `field:"required" json:"grokPattern" yaml:"grokPattern"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/glue_classifier#custom_patterns GlueClassifier#custom_patterns}.
	CustomPatterns *string `field:"optional" json:"customPatterns" yaml:"customPatterns"`
}


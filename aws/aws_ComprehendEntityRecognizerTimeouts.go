// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type ComprehendEntityRecognizerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#create ComprehendEntityRecognizer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#delete ComprehendEntityRecognizer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/comprehend_entity_recognizer#update ComprehendEntityRecognizer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


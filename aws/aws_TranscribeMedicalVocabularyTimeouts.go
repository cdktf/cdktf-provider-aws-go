// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws


type TranscribeMedicalVocabularyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transcribe_medical_vocabulary#create TranscribeMedicalVocabulary#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transcribe_medical_vocabulary#delete TranscribeMedicalVocabulary#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/transcribe_medical_vocabulary#update TranscribeMedicalVocabulary#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


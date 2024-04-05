// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcribemedicalvocabulary


type TranscribeMedicalVocabularyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/transcribe_medical_vocabulary#create TranscribeMedicalVocabulary#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/transcribe_medical_vocabulary#delete TranscribeMedicalVocabulary#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/transcribe_medical_vocabulary#update TranscribeMedicalVocabulary#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


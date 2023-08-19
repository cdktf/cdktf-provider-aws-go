package comprehendentityrecognizer


type ComprehendEntityRecognizerInputDataConfigAugmentedManifests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/comprehend_entity_recognizer#attribute_names ComprehendEntityRecognizer#attribute_names}.
	AttributeNames *[]*string `field:"required" json:"attributeNames" yaml:"attributeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/comprehend_entity_recognizer#s3_uri ComprehendEntityRecognizer#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/comprehend_entity_recognizer#annotation_data_s3_uri ComprehendEntityRecognizer#annotation_data_s3_uri}.
	AnnotationDataS3Uri *string `field:"optional" json:"annotationDataS3Uri" yaml:"annotationDataS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/comprehend_entity_recognizer#document_type ComprehendEntityRecognizer#document_type}.
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/comprehend_entity_recognizer#source_documents_s3_uri ComprehendEntityRecognizer#source_documents_s3_uri}.
	SourceDocumentsS3Uri *string `field:"optional" json:"sourceDocumentsS3Uri" yaml:"sourceDocumentsS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.13.1/docs/resources/comprehend_entity_recognizer#split ComprehendEntityRecognizer#split}.
	Split *string `field:"optional" json:"split" yaml:"split"`
}


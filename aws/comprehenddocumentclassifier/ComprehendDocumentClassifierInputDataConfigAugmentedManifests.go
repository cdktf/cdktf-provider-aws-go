// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package comprehenddocumentclassifier


type ComprehendDocumentClassifierInputDataConfigAugmentedManifests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/comprehend_document_classifier#attribute_names ComprehendDocumentClassifier#attribute_names}.
	AttributeNames *[]*string `field:"required" json:"attributeNames" yaml:"attributeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/comprehend_document_classifier#s3_uri ComprehendDocumentClassifier#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/comprehend_document_classifier#annotation_data_s3_uri ComprehendDocumentClassifier#annotation_data_s3_uri}.
	AnnotationDataS3Uri *string `field:"optional" json:"annotationDataS3Uri" yaml:"annotationDataS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/comprehend_document_classifier#document_type ComprehendDocumentClassifier#document_type}.
	DocumentType *string `field:"optional" json:"documentType" yaml:"documentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/comprehend_document_classifier#source_documents_s3_uri ComprehendDocumentClassifier#source_documents_s3_uri}.
	SourceDocumentsS3Uri *string `field:"optional" json:"sourceDocumentsS3Uri" yaml:"sourceDocumentsS3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/comprehend_document_classifier#split ComprehendDocumentClassifier#split}.
	Split *string `field:"optional" json:"split" yaml:"split"`
}


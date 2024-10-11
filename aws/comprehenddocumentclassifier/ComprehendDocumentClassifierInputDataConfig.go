// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package comprehenddocumentclassifier


type ComprehendDocumentClassifierInputDataConfig struct {
	// augmented_manifests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/comprehend_document_classifier#augmented_manifests ComprehendDocumentClassifier#augmented_manifests}
	AugmentedManifests interface{} `field:"optional" json:"augmentedManifests" yaml:"augmentedManifests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/comprehend_document_classifier#data_format ComprehendDocumentClassifier#data_format}.
	DataFormat *string `field:"optional" json:"dataFormat" yaml:"dataFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/comprehend_document_classifier#label_delimiter ComprehendDocumentClassifier#label_delimiter}.
	LabelDelimiter *string `field:"optional" json:"labelDelimiter" yaml:"labelDelimiter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/comprehend_document_classifier#s3_uri ComprehendDocumentClassifier#s3_uri}.
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.71.0/docs/resources/comprehend_document_classifier#test_s3_uri ComprehendDocumentClassifier#test_s3_uri}.
	TestS3Uri *string `field:"optional" json:"testS3Uri" yaml:"testS3Uri"`
}


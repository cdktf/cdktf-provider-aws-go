// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package comprehenddocumentclassifier


type ComprehendDocumentClassifierOutputDataConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/comprehend_document_classifier#s3_uri ComprehendDocumentClassifier#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/comprehend_document_classifier#kms_key_id ComprehendDocumentClassifier#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}


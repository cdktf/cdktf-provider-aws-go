// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package comprehendentityrecognizer


type ComprehendEntityRecognizerInputDataConfigAnnotations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/comprehend_entity_recognizer#s3_uri ComprehendEntityRecognizer#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/comprehend_entity_recognizer#test_s3_uri ComprehendEntityRecognizer#test_s3_uri}.
	TestS3Uri *string `field:"optional" json:"testS3Uri" yaml:"testS3Uri"`
}


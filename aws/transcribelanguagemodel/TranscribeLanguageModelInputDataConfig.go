// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package transcribelanguagemodel


type TranscribeLanguageModelInputDataConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/transcribe_language_model#data_access_role_arn TranscribeLanguageModel#data_access_role_arn}.
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/transcribe_language_model#s3_uri TranscribeLanguageModel#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/transcribe_language_model#tuning_data_s3_uri TranscribeLanguageModel#tuning_data_s3_uri}.
	TuningDataS3Uri *string `field:"optional" json:"tuningDataS3Uri" yaml:"tuningDataS3Uri"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakeruserprofile


type SagemakerUserProfileUserSettingsSharingSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/sagemaker_user_profile#notebook_output_option SagemakerUserProfile#notebook_output_option}.
	NotebookOutputOption *string `field:"optional" json:"notebookOutputOption" yaml:"notebookOutputOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/sagemaker_user_profile#s3_kms_key_id SagemakerUserProfile#s3_kms_key_id}.
	S3KmsKeyId *string `field:"optional" json:"s3KmsKeyId" yaml:"s3KmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/sagemaker_user_profile#s3_output_path SagemakerUserProfile#s3_output_path}.
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}


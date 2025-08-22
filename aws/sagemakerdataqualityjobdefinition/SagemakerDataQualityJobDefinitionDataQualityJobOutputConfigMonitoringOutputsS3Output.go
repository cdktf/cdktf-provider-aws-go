// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsS3Output struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_data_quality_job_definition#s3_uri SagemakerDataQualityJobDefinition#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_data_quality_job_definition#local_path SagemakerDataQualityJobDefinition#local_path}.
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/sagemaker_data_quality_job_definition#s3_upload_mode SagemakerDataQualityJobDefinition#s3_upload_mode}.
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
}


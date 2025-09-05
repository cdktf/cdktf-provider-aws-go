// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimesdkmediapipelinesmediainsightspipelineconfiguration


type ChimesdkmediapipelinesMediaInsightsPipelineConfigurationElementsAmazonTranscribeCallAnalyticsProcessorConfigurationPostCallAnalyticsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#data_access_role_arn ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#data_access_role_arn}.
	DataAccessRoleArn *string `field:"required" json:"dataAccessRoleArn" yaml:"dataAccessRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#output_location ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#output_location}.
	OutputLocation *string `field:"required" json:"outputLocation" yaml:"outputLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#content_redaction_output ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#content_redaction_output}.
	ContentRedactionOutput *string `field:"optional" json:"contentRedactionOutput" yaml:"contentRedactionOutput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/chimesdkmediapipelines_media_insights_pipeline_configuration#output_encryption_kms_key_id ChimesdkmediapipelinesMediaInsightsPipelineConfiguration#output_encryption_kms_key_id}.
	OutputEncryptionKmsKeyId *string `field:"optional" json:"outputEncryptionKmsKeyId" yaml:"outputEncryptionKmsKeyId"`
}


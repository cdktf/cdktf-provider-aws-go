// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputs struct {
	// s3_output block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/sagemaker_data_quality_job_definition#s3_output SagemakerDataQualityJobDefinition#s3_output}
	S3Output *SagemakerDataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsS3Output `field:"required" json:"s3Output" yaml:"s3Output"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionDataQualityBaselineConfig struct {
	// constraints_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_data_quality_job_definition#constraints_resource SagemakerDataQualityJobDefinition#constraints_resource}
	ConstraintsResource *SagemakerDataQualityJobDefinitionDataQualityBaselineConfigConstraintsResource `field:"optional" json:"constraintsResource" yaml:"constraintsResource"`
	// statistics_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/sagemaker_data_quality_job_definition#statistics_resource SagemakerDataQualityJobDefinition#statistics_resource}
	StatisticsResource *SagemakerDataQualityJobDefinitionDataQualityBaselineConfigStatisticsResource `field:"optional" json:"statisticsResource" yaml:"statisticsResource"`
}


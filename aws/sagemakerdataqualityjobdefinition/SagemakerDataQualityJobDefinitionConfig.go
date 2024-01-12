// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdataqualityjobdefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerDataQualityJobDefinitionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// data_quality_app_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#data_quality_app_specification SagemakerDataQualityJobDefinition#data_quality_app_specification}
	DataQualityAppSpecification *SagemakerDataQualityJobDefinitionDataQualityAppSpecification `field:"required" json:"dataQualityAppSpecification" yaml:"dataQualityAppSpecification"`
	// data_quality_job_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#data_quality_job_input SagemakerDataQualityJobDefinition#data_quality_job_input}
	DataQualityJobInput *SagemakerDataQualityJobDefinitionDataQualityJobInput `field:"required" json:"dataQualityJobInput" yaml:"dataQualityJobInput"`
	// data_quality_job_output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#data_quality_job_output_config SagemakerDataQualityJobDefinition#data_quality_job_output_config}
	DataQualityJobOutputConfig *SagemakerDataQualityJobDefinitionDataQualityJobOutputConfig `field:"required" json:"dataQualityJobOutputConfig" yaml:"dataQualityJobOutputConfig"`
	// job_resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#job_resources SagemakerDataQualityJobDefinition#job_resources}
	JobResources *SagemakerDataQualityJobDefinitionJobResources `field:"required" json:"jobResources" yaml:"jobResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#role_arn SagemakerDataQualityJobDefinition#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// data_quality_baseline_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#data_quality_baseline_config SagemakerDataQualityJobDefinition#data_quality_baseline_config}
	DataQualityBaselineConfig *SagemakerDataQualityJobDefinitionDataQualityBaselineConfig `field:"optional" json:"dataQualityBaselineConfig" yaml:"dataQualityBaselineConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#id SagemakerDataQualityJobDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#name SagemakerDataQualityJobDefinition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// network_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#network_config SagemakerDataQualityJobDefinition#network_config}
	NetworkConfig *SagemakerDataQualityJobDefinitionNetworkConfig `field:"optional" json:"networkConfig" yaml:"networkConfig"`
	// stopping_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#stopping_condition SagemakerDataQualityJobDefinition#stopping_condition}
	StoppingCondition *SagemakerDataQualityJobDefinitionStoppingCondition `field:"optional" json:"stoppingCondition" yaml:"stoppingCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#tags SagemakerDataQualityJobDefinition#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/sagemaker_data_quality_job_definition#tags_all SagemakerDataQualityJobDefinition#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}


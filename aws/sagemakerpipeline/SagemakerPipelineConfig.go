// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerpipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerPipelineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_pipeline#pipeline_display_name SagemakerPipeline#pipeline_display_name}.
	PipelineDisplayName *string `field:"required" json:"pipelineDisplayName" yaml:"pipelineDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_pipeline#pipeline_name SagemakerPipeline#pipeline_name}.
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_pipeline#id SagemakerPipeline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// parallelism_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_pipeline#parallelism_configuration SagemakerPipeline#parallelism_configuration}
	ParallelismConfiguration *SagemakerPipelineParallelismConfiguration `field:"optional" json:"parallelismConfiguration" yaml:"parallelismConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_pipeline#pipeline_definition SagemakerPipeline#pipeline_definition}.
	PipelineDefinition *string `field:"optional" json:"pipelineDefinition" yaml:"pipelineDefinition"`
	// pipeline_definition_s3_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_pipeline#pipeline_definition_s3_location SagemakerPipeline#pipeline_definition_s3_location}
	PipelineDefinitionS3Location *SagemakerPipelinePipelineDefinitionS3Location `field:"optional" json:"pipelineDefinitionS3Location" yaml:"pipelineDefinitionS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_pipeline#pipeline_description SagemakerPipeline#pipeline_description}.
	PipelineDescription *string `field:"optional" json:"pipelineDescription" yaml:"pipelineDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_pipeline#role_arn SagemakerPipeline#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_pipeline#tags SagemakerPipeline#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/sagemaker_pipeline#tags_all SagemakerPipeline#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}


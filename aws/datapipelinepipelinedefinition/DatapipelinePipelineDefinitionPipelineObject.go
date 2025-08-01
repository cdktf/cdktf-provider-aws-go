// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipelinedefinition


type DatapipelinePipelineDefinitionPipelineObject struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/datapipeline_pipeline_definition#id DatapipelinePipelineDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/datapipeline_pipeline_definition#name DatapipelinePipelineDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/datapipeline_pipeline_definition#field DatapipelinePipelineDefinition#field}
	Field interface{} `field:"optional" json:"field" yaml:"field"`
}


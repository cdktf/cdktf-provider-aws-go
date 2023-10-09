// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipelinedefinition


type DatapipelinePipelineDefinitionParameterObjectAttribute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/datapipeline_pipeline_definition#key DatapipelinePipelineDefinition#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/datapipeline_pipeline_definition#string_value DatapipelinePipelineDefinition#string_value}.
	StringValue *string `field:"required" json:"stringValue" yaml:"stringValue"`
}


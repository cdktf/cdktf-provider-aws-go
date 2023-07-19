package datapipelinepipelinedefinition


type DatapipelinePipelineDefinitionPipelineObjectField struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/datapipeline_pipeline_definition#key DatapipelinePipelineDefinition#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/datapipeline_pipeline_definition#ref_value DatapipelinePipelineDefinition#ref_value}.
	RefValue *string `field:"optional" json:"refValue" yaml:"refValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.8.0/docs/resources/datapipeline_pipeline_definition#string_value DatapipelinePipelineDefinition#string_value}.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}


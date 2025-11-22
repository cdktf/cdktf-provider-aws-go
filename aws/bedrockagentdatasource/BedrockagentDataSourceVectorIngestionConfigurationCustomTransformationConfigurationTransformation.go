// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentdatasource


type BedrockagentDataSourceVectorIngestionConfigurationCustomTransformationConfigurationTransformation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/bedrockagent_data_source#step_to_apply BedrockagentDataSource#step_to_apply}.
	StepToApply *string `field:"required" json:"stepToApply" yaml:"stepToApply"`
	// transformation_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/bedrockagent_data_source#transformation_function BedrockagentDataSource#transformation_function}
	TransformationFunction interface{} `field:"optional" json:"transformationFunction" yaml:"transformationFunction"`
}


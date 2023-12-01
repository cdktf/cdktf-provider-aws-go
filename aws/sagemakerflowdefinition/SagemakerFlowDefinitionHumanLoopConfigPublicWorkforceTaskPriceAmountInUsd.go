// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerflowdefinition


type SagemakerFlowDefinitionHumanLoopConfigPublicWorkforceTaskPriceAmountInUsd struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/sagemaker_flow_definition#cents SagemakerFlowDefinition#cents}.
	Cents *float64 `field:"optional" json:"cents" yaml:"cents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/sagemaker_flow_definition#dollars SagemakerFlowDefinition#dollars}.
	Dollars *float64 `field:"optional" json:"dollars" yaml:"dollars"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/sagemaker_flow_definition#tenth_fractions_of_a_cent SagemakerFlowDefinition#tenth_fractions_of_a_cent}.
	TenthFractionsOfACent *float64 `field:"optional" json:"tenthFractionsOfACent" yaml:"tenthFractionsOfACent"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerflowdefinition


type SagemakerFlowDefinitionHumanLoopActivationConfig struct {
	// human_loop_activation_conditions_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/sagemaker_flow_definition#human_loop_activation_conditions_config SagemakerFlowDefinition#human_loop_activation_conditions_config}
	HumanLoopActivationConditionsConfig *SagemakerFlowDefinitionHumanLoopActivationConfigHumanLoopActivationConditionsConfig `field:"optional" json:"humanLoopActivationConditionsConfig" yaml:"humanLoopActivationConditionsConfig"`
}


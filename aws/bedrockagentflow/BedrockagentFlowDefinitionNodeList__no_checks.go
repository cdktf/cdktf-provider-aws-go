// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package bedrockagentflow

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BedrockagentFlowDefinitionNodeList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BedrockagentFlowDefinitionNodeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBedrockagentFlowDefinitionNodeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


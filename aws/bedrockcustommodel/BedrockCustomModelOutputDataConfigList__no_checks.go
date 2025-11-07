// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package bedrockcustommodel

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockCustomModelOutputDataConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BedrockCustomModelOutputDataConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BedrockCustomModelOutputDataConfigList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BedrockCustomModelOutputDataConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BedrockCustomModelOutputDataConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BedrockCustomModelOutputDataConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BedrockCustomModelOutputDataConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBedrockCustomModelOutputDataConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


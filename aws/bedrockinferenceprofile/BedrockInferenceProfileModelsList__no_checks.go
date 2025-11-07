// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package bedrockinferenceprofile

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BedrockInferenceProfileModelsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BedrockInferenceProfileModelsList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BedrockInferenceProfileModelsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BedrockInferenceProfileModelsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BedrockInferenceProfileModelsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BedrockInferenceProfileModelsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBedrockInferenceProfileModelsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


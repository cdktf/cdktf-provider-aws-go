// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lambdacapacityprovider

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaCapacityProviderVpcConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (l *jsiiProxy_LambdaCapacityProviderVpcConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LambdaCapacityProviderVpcConfigList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LambdaCapacityProviderVpcConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LambdaCapacityProviderVpcConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LambdaCapacityProviderVpcConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LambdaCapacityProviderVpcConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLambdaCapacityProviderVpcConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


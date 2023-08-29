// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appconfigextension

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppconfigExtensionActionPointActionList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppconfigExtensionActionPointActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionActionPointActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionActionPointActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionActionPointActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppconfigExtensionActionPointActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppconfigExtensionActionPointActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


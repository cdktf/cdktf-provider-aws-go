// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package appsyncapi

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppsyncApiEventConfigList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AppsyncApiEventConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppsyncApiEventConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppsyncApiEventConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppsyncApiEventConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppsyncApiEventConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppsyncApiEventConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppsyncApiEventConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


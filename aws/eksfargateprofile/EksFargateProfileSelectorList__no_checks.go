// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eksfargateprofile

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksFargateProfileSelectorList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksFargateProfileSelectorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksFargateProfileSelectorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksFargateProfileSelectorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksFargateProfileSelectorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksFargateProfileSelectorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksFargateProfileSelectorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


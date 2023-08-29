// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package codebuildproject

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodebuildProjectSecondarySourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodebuildProjectSecondarySourcesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodebuildProjectSecondarySourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodebuildProjectSecondarySourcesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


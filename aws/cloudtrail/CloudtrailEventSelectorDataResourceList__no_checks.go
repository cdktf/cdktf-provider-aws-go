// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudtrail

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudtrailEventSelectorDataResourceList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudtrailEventSelectorDataResourceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudtrailEventSelectorDataResourceList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudtrailEventSelectorDataResourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudtrailEventSelectorDataResourceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudtrailEventSelectorDataResourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudtrailEventSelectorDataResourceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


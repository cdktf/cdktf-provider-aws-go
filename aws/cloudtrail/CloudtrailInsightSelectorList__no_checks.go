// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package cloudtrail

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudtrailInsightSelectorList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (c *jsiiProxy_CloudtrailInsightSelectorList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudtrailInsightSelectorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudtrailInsightSelectorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudtrailInsightSelectorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudtrailInsightSelectorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudtrailInsightSelectorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudtrailInsightSelectorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


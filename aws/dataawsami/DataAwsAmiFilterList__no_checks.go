// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsami

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsAmiFilterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsAmiFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsAmiFilterList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsAmiFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsAmiFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsAmiFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsAmiFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsAmiFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


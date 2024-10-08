// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawslb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsLbSubnetMappingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsLbSubnetMappingList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsLbSubnetMappingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbSubnetMappingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbSubnetMappingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbSubnetMappingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsLbSubnetMappingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


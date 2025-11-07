// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsinstance

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsInstanceRootBlockDeviceList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsInstanceRootBlockDeviceList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsInstanceRootBlockDeviceList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsInstanceRootBlockDeviceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsInstanceRootBlockDeviceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsInstanceRootBlockDeviceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsInstanceRootBlockDeviceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


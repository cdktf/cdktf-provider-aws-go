// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsdmsendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsDmsEndpointS3SettingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsDmsEndpointS3SettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsDmsEndpointS3SettingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


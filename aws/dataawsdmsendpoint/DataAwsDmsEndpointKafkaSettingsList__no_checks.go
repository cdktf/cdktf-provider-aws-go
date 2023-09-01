// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsdmsendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsDmsEndpointKafkaSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsDmsEndpointKafkaSettingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsDmsEndpointKafkaSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsDmsEndpointKafkaSettingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsDmsEndpointKafkaSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsDmsEndpointKafkaSettingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


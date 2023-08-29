// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsdmsendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsDmsEndpointMongodbSettingsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsDmsEndpointMongodbSettingsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsDmsEndpointMongodbSettingsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsDmsEndpointMongodbSettingsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsDmsEndpointMongodbSettingsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsDmsEndpointMongodbSettingsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsDmsEndpointMongodbSettingsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


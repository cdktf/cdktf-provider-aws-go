// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsmqbroker

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsMqBrokerUserList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsMqBrokerUserList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsMqBrokerUserList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerUserList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerUserList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerUserList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsMqBrokerUserListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


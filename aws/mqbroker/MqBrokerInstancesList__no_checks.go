// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mqbroker

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MqBrokerInstancesList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MqBrokerInstancesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MqBrokerInstancesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MqBrokerInstancesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MqBrokerInstancesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMqBrokerInstancesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


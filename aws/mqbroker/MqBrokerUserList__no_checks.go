//go:build no_runtime_type_checking

package mqbroker

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MqBrokerUserList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MqBrokerUserList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MqBrokerUserList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MqBrokerUserList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MqBrokerUserList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MqBrokerUserList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMqBrokerUserListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


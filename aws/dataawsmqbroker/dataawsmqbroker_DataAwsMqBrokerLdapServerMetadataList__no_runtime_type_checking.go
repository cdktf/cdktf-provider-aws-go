//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dataawsmqbroker

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsMqBrokerLdapServerMetadataList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsMqBrokerLdapServerMetadataList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadataList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadataList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsMqBrokerLdapServerMetadataList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsMqBrokerLdapServerMetadataListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


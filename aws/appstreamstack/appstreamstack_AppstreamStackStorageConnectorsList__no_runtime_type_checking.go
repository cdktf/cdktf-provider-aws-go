//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package appstreamstack

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppstreamStackStorageConnectorsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppstreamStackStorageConnectorsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackStorageConnectorsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackStorageConnectorsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackStorageConnectorsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppstreamStackStorageConnectorsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppstreamStackStorageConnectorsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


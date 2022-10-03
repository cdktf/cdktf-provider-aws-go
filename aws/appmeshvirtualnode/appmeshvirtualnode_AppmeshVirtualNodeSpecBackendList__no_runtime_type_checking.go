//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package appmeshvirtualnode

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AppmeshVirtualNodeSpecBackendList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AppmeshVirtualNodeSpecBackendList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAppmeshVirtualNodeSpecBackendListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package msk

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MskServerlessClusterVpcConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MskServerlessClusterVpcConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MskServerlessClusterVpcConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MskServerlessClusterVpcConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MskServerlessClusterVpcConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MskServerlessClusterVpcConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMskServerlessClusterVpcConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


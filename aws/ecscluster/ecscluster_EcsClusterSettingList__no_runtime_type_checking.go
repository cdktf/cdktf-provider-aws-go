//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ecscluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsClusterSettingList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsClusterSettingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsClusterSettingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsClusterSettingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsClusterSettingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsClusterSettingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsClusterSettingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


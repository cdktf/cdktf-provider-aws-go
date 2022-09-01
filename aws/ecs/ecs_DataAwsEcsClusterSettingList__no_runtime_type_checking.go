//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ecs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEcsClusterSettingList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEcsClusterSettingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsClusterSettingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsClusterSettingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsClusterSettingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEcsClusterSettingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


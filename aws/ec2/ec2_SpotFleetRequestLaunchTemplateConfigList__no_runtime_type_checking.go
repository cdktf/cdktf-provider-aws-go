//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSpotFleetRequestLaunchTemplateConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


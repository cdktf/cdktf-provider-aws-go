//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SpotFleetRequestLaunchTemplateConfigOverridesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSpotFleetRequestLaunchTemplateConfigOverridesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


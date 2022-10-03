//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2fleet

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2FleetLaunchTemplateConfigOverrideList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2FleetLaunchTemplateConfigOverrideListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


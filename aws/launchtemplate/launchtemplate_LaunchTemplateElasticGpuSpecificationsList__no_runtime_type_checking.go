//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package launchtemplate

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LaunchTemplateElasticGpuSpecificationsList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LaunchTemplateElasticGpuSpecificationsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateElasticGpuSpecificationsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateElasticGpuSpecificationsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateElasticGpuSpecificationsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateElasticGpuSpecificationsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLaunchTemplateElasticGpuSpecificationsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


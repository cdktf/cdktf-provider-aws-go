//go:build no_runtime_type_checking

package launchtemplate

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LaunchTemplateLicenseSpecificationList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LaunchTemplateLicenseSpecificationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateLicenseSpecificationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateLicenseSpecificationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateLicenseSpecificationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LaunchTemplateLicenseSpecificationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLaunchTemplateLicenseSpecificationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


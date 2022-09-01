//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsLaunchTemplateCreditSpecificationList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsLaunchTemplateCreditSpecificationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsLaunchTemplateCreditSpecificationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsLaunchTemplateCreditSpecificationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsLaunchTemplateCreditSpecificationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsLaunchTemplateCreditSpecificationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


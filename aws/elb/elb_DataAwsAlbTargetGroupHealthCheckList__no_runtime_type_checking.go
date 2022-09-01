//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package elb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsAlbTargetGroupHealthCheckList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsAlbTargetGroupHealthCheckList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbTargetGroupHealthCheckList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbTargetGroupHealthCheckList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbTargetGroupHealthCheckList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsAlbTargetGroupHealthCheckListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


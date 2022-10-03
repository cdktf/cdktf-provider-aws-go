//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dataawselb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsElbListenerList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsElbListenerList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsElbListenerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsElbListenerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsElbListenerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsElbListenerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


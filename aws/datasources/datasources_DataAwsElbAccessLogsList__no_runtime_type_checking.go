//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package datasources

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsElbAccessLogsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsElbAccessLogsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsElbAccessLogsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsElbAccessLogsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsElbAccessLogsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsElbAccessLogsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

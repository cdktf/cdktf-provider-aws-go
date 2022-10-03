//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package quicksightdatasource

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QuicksightDataSourcePermissionList) validateGetParameters(index *float64) error {
	return nil
}

func (q *jsiiProxy_QuicksightDataSourcePermissionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_QuicksightDataSourcePermissionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_QuicksightDataSourcePermissionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_QuicksightDataSourcePermissionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_QuicksightDataSourcePermissionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewQuicksightDataSourcePermissionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


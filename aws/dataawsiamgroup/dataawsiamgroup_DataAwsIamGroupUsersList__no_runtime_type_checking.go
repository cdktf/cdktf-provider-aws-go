//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dataawsiamgroup

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsIamGroupUsersList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsIamGroupUsersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsIamGroupUsersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsIamGroupUsersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsIamGroupUsersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsIamGroupUsersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


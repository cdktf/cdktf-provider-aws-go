//go:build no_runtime_type_checking

package iamrole

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IamRoleRoleLastUsedList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IamRoleRoleLastUsedList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IamRoleRoleLastUsedList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IamRoleRoleLastUsedList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IamRoleRoleLastUsedList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIamRoleRoleLastUsedListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


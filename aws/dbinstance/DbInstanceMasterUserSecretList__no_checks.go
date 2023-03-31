//go:build no_runtime_type_checking

package dbinstance

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DbInstanceMasterUserSecretList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DbInstanceMasterUserSecretList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DbInstanceMasterUserSecretList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DbInstanceMasterUserSecretList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DbInstanceMasterUserSecretList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDbInstanceMasterUserSecretListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


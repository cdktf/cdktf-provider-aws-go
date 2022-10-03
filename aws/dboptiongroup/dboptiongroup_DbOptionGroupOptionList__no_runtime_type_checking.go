//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dboptiongroup

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DbOptionGroupOptionList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DbOptionGroupOptionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DbOptionGroupOptionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DbOptionGroupOptionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DbOptionGroupOptionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DbOptionGroupOptionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDbOptionGroupOptionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


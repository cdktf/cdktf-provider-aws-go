//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package rds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DbParameterGroupParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DbParameterGroupParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DbParameterGroupParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DbParameterGroupParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DbParameterGroupParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DbParameterGroupParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDbParameterGroupParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


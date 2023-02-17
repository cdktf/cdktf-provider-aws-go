//go:build no_runtime_type_checking

package eksnodegroup

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksNodeGroupTaintList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksNodeGroupTaintList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksNodeGroupTaintList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksNodeGroupTaintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksNodeGroupTaintList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksNodeGroupTaintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksNodeGroupTaintListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


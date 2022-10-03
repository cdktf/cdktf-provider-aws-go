//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package rdsclusterparametergroup

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsClusterParameterGroupParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsClusterParameterGroupParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsClusterParameterGroupParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RdsClusterParameterGroupParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsClusterParameterGroupParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsClusterParameterGroupParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsClusterParameterGroupParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


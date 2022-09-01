//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package redshift

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedshiftdataStatementParametersList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedshiftdataStatementParametersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedshiftdataStatementParametersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RedshiftdataStatementParametersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedshiftdataStatementParametersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedshiftdataStatementParametersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedshiftdataStatementParametersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


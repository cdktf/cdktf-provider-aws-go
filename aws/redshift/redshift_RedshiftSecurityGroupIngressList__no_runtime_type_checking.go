//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package redshift

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RedshiftSecurityGroupIngressList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RedshiftSecurityGroupIngressList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RedshiftSecurityGroupIngressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RedshiftSecurityGroupIngressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RedshiftSecurityGroupIngressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RedshiftSecurityGroupIngressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRedshiftSecurityGroupIngressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


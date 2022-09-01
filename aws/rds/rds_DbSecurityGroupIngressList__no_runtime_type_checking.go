//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package rds

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DbSecurityGroupIngressList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DbSecurityGroupIngressList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DbSecurityGroupIngressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DbSecurityGroupIngressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DbSecurityGroupIngressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DbSecurityGroupIngressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDbSecurityGroupIngressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


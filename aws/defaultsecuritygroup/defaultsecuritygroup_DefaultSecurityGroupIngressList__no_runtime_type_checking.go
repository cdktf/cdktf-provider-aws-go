//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package defaultsecuritygroup

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DefaultSecurityGroupIngressList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DefaultSecurityGroupIngressList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DefaultSecurityGroupIngressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DefaultSecurityGroupIngressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DefaultSecurityGroupIngressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DefaultSecurityGroupIngressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDefaultSecurityGroupIngressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


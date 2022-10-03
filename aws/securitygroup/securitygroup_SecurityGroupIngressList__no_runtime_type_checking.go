//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package securitygroup

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityGroupIngressList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecurityGroupIngressList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecurityGroupIngressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecurityGroupIngressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecurityGroupIngressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecurityGroupIngressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecurityGroupIngressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


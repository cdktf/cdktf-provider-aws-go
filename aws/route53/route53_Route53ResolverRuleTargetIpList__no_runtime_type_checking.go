//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package route53

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route53ResolverRuleTargetIpList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_Route53ResolverRuleTargetIpList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Route53ResolverRuleTargetIpList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Route53ResolverRuleTargetIpList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Route53ResolverRuleTargetIpList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Route53ResolverRuleTargetIpList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoute53ResolverRuleTargetIpListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


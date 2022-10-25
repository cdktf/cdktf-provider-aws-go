//go:build no_runtime_type_checking

package route53resolverendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route53ResolverEndpointIpAddressList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_Route53ResolverEndpointIpAddressList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Route53ResolverEndpointIpAddressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Route53ResolverEndpointIpAddressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Route53ResolverEndpointIpAddressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Route53ResolverEndpointIpAddressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRoute53ResolverEndpointIpAddressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


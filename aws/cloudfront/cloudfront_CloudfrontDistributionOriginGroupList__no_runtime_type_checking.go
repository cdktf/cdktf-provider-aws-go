//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudfrontDistributionOriginGroupList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudfrontDistributionOriginGroupList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudfrontDistributionOriginGroupList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudfrontDistributionOriginGroupListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


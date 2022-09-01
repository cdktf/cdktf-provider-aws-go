//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cloudfront

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudfrontDistributionTrustedSignersList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CloudfrontDistributionTrustedSignersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSignersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSignersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CloudfrontDistributionTrustedSignersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCloudfrontDistributionTrustedSignersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package securityhub

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityhubInsightFiltersProductArnList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersProductArnList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersProductArnList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersProductArnList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersProductArnList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersProductArnList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecurityhubInsightFiltersProductArnListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


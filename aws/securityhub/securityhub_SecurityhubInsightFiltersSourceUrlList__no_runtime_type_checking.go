//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package securityhub

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityhubInsightFiltersSourceUrlList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersSourceUrlList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersSourceUrlList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersSourceUrlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersSourceUrlList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersSourceUrlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecurityhubInsightFiltersSourceUrlListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


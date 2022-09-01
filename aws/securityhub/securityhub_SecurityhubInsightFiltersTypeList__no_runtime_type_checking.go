//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package securityhub

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityhubInsightFiltersTypeList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersTypeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersTypeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersTypeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersTypeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersTypeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecurityhubInsightFiltersTypeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


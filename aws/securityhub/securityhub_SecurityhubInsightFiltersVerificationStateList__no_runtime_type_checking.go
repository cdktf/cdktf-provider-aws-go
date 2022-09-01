//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package securityhub

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityhubInsightFiltersVerificationStateList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersVerificationStateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersVerificationStateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersVerificationStateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersVerificationStateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersVerificationStateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecurityhubInsightFiltersVerificationStateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package securityhub

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecurityhubInsightFiltersRecordStateList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SecurityhubInsightFiltersRecordStateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersRecordStateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersRecordStateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersRecordStateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SecurityhubInsightFiltersRecordStateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSecurityhubInsightFiltersRecordStateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


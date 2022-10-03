//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kendraindex

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KendraIndexIndexStatisticsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KendraIndexIndexStatisticsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KendraIndexIndexStatisticsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KendraIndexIndexStatisticsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KendraIndexIndexStatisticsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKendraIndexIndexStatisticsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


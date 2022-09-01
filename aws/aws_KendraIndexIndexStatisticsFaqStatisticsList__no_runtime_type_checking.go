//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KendraIndexIndexStatisticsFaqStatisticsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KendraIndexIndexStatisticsFaqStatisticsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KendraIndexIndexStatisticsFaqStatisticsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KendraIndexIndexStatisticsFaqStatisticsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KendraIndexIndexStatisticsFaqStatisticsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKendraIndexIndexStatisticsFaqStatisticsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


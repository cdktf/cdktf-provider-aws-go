//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package accessanalyzer

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessanalyzerArchiveRuleFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AccessanalyzerArchiveRuleFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AccessanalyzerArchiveRuleFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AccessanalyzerArchiveRuleFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AccessanalyzerArchiveRuleFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AccessanalyzerArchiveRuleFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAccessanalyzerArchiveRuleFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


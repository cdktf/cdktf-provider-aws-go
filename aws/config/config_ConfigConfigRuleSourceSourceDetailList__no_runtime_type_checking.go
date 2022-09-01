//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package config

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigConfigRuleSourceSourceDetailList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigConfigRuleSourceSourceDetailList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigConfigRuleSourceSourceDetailList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigConfigRuleSourceSourceDetailList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigConfigRuleSourceSourceDetailList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigConfigRuleSourceSourceDetailList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigConfigRuleSourceSourceDetailListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


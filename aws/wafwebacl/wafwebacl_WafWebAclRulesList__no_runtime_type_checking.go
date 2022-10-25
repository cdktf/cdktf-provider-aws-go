//go:build no_runtime_type_checking

package wafwebacl

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WafWebAclRulesList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WafWebAclRulesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WafWebAclRulesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WafWebAclRulesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WafWebAclRulesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WafWebAclRulesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWafWebAclRulesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package wafregional

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WafregionalWebAclRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WafregionalWebAclRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WafregionalWebAclRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WafregionalWebAclRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WafregionalWebAclRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WafregionalWebAclRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWafregionalWebAclRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


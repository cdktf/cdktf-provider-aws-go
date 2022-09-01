//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package wafregional

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WafregionalRulePredicateList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WafregionalRulePredicateList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WafregionalRulePredicateList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WafregionalRulePredicateList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WafregionalRulePredicateList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WafregionalRulePredicateList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWafregionalRulePredicateListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


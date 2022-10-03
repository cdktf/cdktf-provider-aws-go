//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package wafrule

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WafRulePredicatesList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WafRulePredicatesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WafRulePredicatesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WafRulePredicatesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WafRulePredicatesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WafRulePredicatesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWafRulePredicatesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


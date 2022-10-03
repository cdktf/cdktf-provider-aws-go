//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package wafrulegroup

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WafRuleGroupActivatedRuleList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WafRuleGroupActivatedRuleList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WafRuleGroupActivatedRuleList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WafRuleGroupActivatedRuleList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WafRuleGroupActivatedRuleList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WafRuleGroupActivatedRuleList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWafRuleGroupActivatedRuleListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


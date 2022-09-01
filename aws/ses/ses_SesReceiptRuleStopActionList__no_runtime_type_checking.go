//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ses

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SesReceiptRuleStopActionList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SesReceiptRuleStopActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleStopActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleStopActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleStopActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleStopActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSesReceiptRuleStopActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


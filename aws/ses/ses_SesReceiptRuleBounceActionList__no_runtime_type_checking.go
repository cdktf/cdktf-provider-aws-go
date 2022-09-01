//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ses

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SesReceiptRuleBounceActionList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SesReceiptRuleBounceActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleBounceActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleBounceActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleBounceActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleBounceActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSesReceiptRuleBounceActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


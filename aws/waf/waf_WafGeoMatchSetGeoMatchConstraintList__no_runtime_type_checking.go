//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package waf

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WafGeoMatchSetGeoMatchConstraintList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WafGeoMatchSetGeoMatchConstraintList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WafGeoMatchSetGeoMatchConstraintList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WafGeoMatchSetGeoMatchConstraintList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WafGeoMatchSetGeoMatchConstraintList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WafGeoMatchSetGeoMatchConstraintList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWafGeoMatchSetGeoMatchConstraintListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


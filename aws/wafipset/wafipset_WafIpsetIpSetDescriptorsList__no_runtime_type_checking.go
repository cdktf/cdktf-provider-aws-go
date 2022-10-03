//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package wafipset

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WafIpsetIpSetDescriptorsList) validateGetParameters(index *float64) error {
	return nil
}

func (w *jsiiProxy_WafIpsetIpSetDescriptorsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_WafIpsetIpSetDescriptorsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_WafIpsetIpSetDescriptorsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_WafIpsetIpSetDescriptorsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_WafIpsetIpSetDescriptorsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewWafIpsetIpSetDescriptorsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


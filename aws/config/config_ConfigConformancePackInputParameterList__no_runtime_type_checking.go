//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package config

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigConformancePackInputParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_ConfigConformancePackInputParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackInputParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackInputParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackInputParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ConfigConformancePackInputParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewConfigConformancePackInputParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


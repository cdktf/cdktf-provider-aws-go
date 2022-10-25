//go:build no_runtime_type_checking

package medialiveinput

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MedialiveInputSourcesList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MedialiveInputSourcesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MedialiveInputSourcesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MedialiveInputSourcesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MedialiveInputSourcesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MedialiveInputSourcesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMedialiveInputSourcesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


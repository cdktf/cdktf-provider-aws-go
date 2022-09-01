//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ecs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsTaskDefinitionInferenceAcceleratorList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsTaskDefinitionInferenceAcceleratorList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsTaskDefinitionInferenceAcceleratorList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsTaskDefinitionInferenceAcceleratorList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsTaskDefinitionInferenceAcceleratorList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsTaskDefinitionInferenceAcceleratorList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsTaskDefinitionInferenceAcceleratorListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


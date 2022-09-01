//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ecs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsTaskDefinitionVolumeList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsTaskDefinitionVolumeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsTaskDefinitionVolumeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsTaskDefinitionVolumeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


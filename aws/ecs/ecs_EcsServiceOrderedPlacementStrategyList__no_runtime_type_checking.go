//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ecs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsServiceOrderedPlacementStrategyList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsServiceOrderedPlacementStrategyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsServiceOrderedPlacementStrategyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsServiceOrderedPlacementStrategyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsServiceOrderedPlacementStrategyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsServiceOrderedPlacementStrategyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsServiceOrderedPlacementStrategyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


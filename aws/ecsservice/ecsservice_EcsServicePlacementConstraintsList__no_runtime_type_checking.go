//go:build no_runtime_type_checking

package ecsservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsServicePlacementConstraintsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsServicePlacementConstraintsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsServicePlacementConstraintsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsServicePlacementConstraintsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsServicePlacementConstraintsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsServicePlacementConstraintsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsServicePlacementConstraintsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


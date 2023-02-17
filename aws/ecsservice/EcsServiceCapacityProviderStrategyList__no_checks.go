//go:build no_runtime_type_checking

package ecsservice

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsServiceCapacityProviderStrategyList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EcsServiceCapacityProviderStrategyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EcsServiceCapacityProviderStrategyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EcsServiceCapacityProviderStrategyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EcsServiceCapacityProviderStrategyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EcsServiceCapacityProviderStrategyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEcsServiceCapacityProviderStrategyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


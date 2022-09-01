//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package autoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AutoscalingGroupInitialLifecycleHookList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AutoscalingGroupInitialLifecycleHookList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAutoscalingGroupInitialLifecycleHookListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


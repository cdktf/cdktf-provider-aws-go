//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package autoscaling

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AutoscalingPolicyStepAdjustmentList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AutoscalingPolicyStepAdjustmentList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AutoscalingPolicyStepAdjustmentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AutoscalingPolicyStepAdjustmentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AutoscalingPolicyStepAdjustmentList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AutoscalingPolicyStepAdjustmentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAutoscalingPolicyStepAdjustmentListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking

package transferworkflow

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransferWorkflowStepsList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TransferWorkflowStepsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TransferWorkflowStepsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TransferWorkflowStepsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TransferWorkflowStepsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TransferWorkflowStepsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTransferWorkflowStepsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


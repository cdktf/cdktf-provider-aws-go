//go:build no_runtime_type_checking

package kmsgrant

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KmsGrantConstraintsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KmsGrantConstraintsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KmsGrantConstraintsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KmsGrantConstraintsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KmsGrantConstraintsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KmsGrantConstraintsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKmsGrantConstraintsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


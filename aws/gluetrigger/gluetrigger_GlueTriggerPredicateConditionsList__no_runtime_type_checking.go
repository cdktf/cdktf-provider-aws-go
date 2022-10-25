//go:build no_runtime_type_checking

package gluetrigger

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueTriggerPredicateConditionsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueTriggerPredicateConditionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueTriggerPredicateConditionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueTriggerPredicateConditionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueTriggerPredicateConditionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueTriggerPredicateConditionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueTriggerPredicateConditionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package lex

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LexIntentSlotList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LexIntentSlotList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LexIntentSlotList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LexIntentSlotList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LexIntentSlotList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LexIntentSlotList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLexIntentSlotListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

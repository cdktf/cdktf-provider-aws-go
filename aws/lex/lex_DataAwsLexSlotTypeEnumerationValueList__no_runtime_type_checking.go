//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package lex

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValueList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsLexSlotTypeEnumerationValueList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValueList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValueList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsLexSlotTypeEnumerationValueList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsLexSlotTypeEnumerationValueListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package cognitouserpool

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CognitoUserPoolSchemaList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CognitoUserPoolSchemaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CognitoUserPoolSchemaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CognitoUserPoolSchemaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CognitoUserPoolSchemaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CognitoUserPoolSchemaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCognitoUserPoolSchemaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


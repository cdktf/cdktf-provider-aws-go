//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package iam

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsIamPolicyDocumentStatementList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsIamPolicyDocumentStatementList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsIamPolicyDocumentStatementListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


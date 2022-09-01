//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableGlobalSecondaryIndexList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbTableGlobalSecondaryIndexListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


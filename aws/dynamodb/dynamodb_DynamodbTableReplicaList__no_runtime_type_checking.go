//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbTableReplicaList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbTableReplicaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableReplicaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableReplicaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableReplicaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbTableReplicaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbTableReplicaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


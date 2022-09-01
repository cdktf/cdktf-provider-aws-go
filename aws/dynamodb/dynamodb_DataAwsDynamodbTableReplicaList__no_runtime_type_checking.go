//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsDynamodbTableReplicaList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsDynamodbTableReplicaList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsDynamodbTableReplicaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsDynamodbTableReplicaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsDynamodbTableReplicaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsDynamodbTableReplicaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


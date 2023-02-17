//go:build no_runtime_type_checking

package keyspacestable

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionPartitionKeyList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionPartitionKeyList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionPartitionKeyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionPartitionKeyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionPartitionKeyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionPartitionKeyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKeyspacesTableSchemaDefinitionPartitionKeyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


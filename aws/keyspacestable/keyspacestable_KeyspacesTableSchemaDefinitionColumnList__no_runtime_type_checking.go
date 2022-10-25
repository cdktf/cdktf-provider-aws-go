//go:build no_runtime_type_checking

package keyspacestable

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionColumnList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionColumnList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionColumnList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionColumnList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionColumnList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionColumnList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKeyspacesTableSchemaDefinitionColumnListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


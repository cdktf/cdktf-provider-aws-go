//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionStaticColumnList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KeyspacesTableSchemaDefinitionStaticColumnList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionStaticColumnList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionStaticColumnList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionStaticColumnList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KeyspacesTableSchemaDefinitionStaticColumnList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKeyspacesTableSchemaDefinitionStaticColumnListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


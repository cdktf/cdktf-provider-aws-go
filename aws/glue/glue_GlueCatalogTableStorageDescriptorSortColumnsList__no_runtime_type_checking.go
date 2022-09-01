//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSortColumnsList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueCatalogTableStorageDescriptorSortColumnsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSortColumnsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSortColumnsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSortColumnsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueCatalogTableStorageDescriptorSortColumnsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueCatalogTableStorageDescriptorSortColumnsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


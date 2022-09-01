//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueCatalogTablePartitionIndexList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueCatalogTablePartitionIndexList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueCatalogTablePartitionIndexList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GlueCatalogTablePartitionIndexList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueCatalogTablePartitionIndexList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueCatalogTablePartitionIndexList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueCatalogTablePartitionIndexListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


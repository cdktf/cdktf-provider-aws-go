//go:build no_runtime_type_checking

package alb

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbSubnetMappingList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlbSubnetMappingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlbSubnetMappingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlbSubnetMappingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlbSubnetMappingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlbSubnetMappingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlbSubnetMappingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


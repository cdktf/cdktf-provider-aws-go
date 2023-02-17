//go:build no_runtime_type_checking

package lb

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LbSubnetMappingList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LbSubnetMappingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LbSubnetMappingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LbSubnetMappingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LbSubnetMappingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LbSubnetMappingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLbSubnetMappingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


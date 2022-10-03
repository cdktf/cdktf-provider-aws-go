//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package dataawsalb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsAlbSubnetMappingList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsAlbSubnetMappingList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbSubnetMappingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbSubnetMappingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbSubnetMappingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsAlbSubnetMappingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


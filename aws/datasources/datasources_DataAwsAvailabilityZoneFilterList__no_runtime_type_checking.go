//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package datasources

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsAvailabilityZoneFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsAvailabilityZoneFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsAvailabilityZoneFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsAvailabilityZoneFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsAvailabilityZoneFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsAvailabilityZoneFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsAvailabilityZoneFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


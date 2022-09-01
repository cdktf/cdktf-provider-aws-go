//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package servicecatalog

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsServicecatalogLaunchPathsSummariesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsServicecatalogLaunchPathsSummariesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


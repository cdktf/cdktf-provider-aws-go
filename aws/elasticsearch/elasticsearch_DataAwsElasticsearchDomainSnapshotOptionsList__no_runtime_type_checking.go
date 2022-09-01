//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package elasticsearch

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainSnapshotOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsElasticsearchDomainSnapshotOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


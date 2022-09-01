//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package elasticsearch

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRestList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRestList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRestList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRestList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainEncryptionAtRestList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsElasticsearchDomainEncryptionAtRestListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package elasticsearch

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsElasticsearchDomainVpcOptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainVpcOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsElasticsearchDomainVpcOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


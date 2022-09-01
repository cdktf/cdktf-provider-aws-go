//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package elasticsearch

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsElasticsearchDomainAdvancedSecurityOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsElasticsearchDomainAdvancedSecurityOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


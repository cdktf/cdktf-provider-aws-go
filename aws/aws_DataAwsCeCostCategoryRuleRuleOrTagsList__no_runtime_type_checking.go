//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsCeCostCategoryRuleRuleOrTagsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsCeCostCategoryRuleRuleOrTagsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsCeCostCategoryRuleRuleOrTagsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsCeCostCategoryRuleRuleOrTagsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsCeCostCategoryRuleRuleOrTagsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsCeCostCategoryRuleRuleOrTagsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


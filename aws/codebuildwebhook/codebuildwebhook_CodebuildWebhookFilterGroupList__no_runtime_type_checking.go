//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package codebuildwebhook

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodebuildWebhookFilterGroupList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodebuildWebhookFilterGroupList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodebuildWebhookFilterGroupList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodebuildWebhookFilterGroupList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodebuildWebhookFilterGroupList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodebuildWebhookFilterGroupList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodebuildWebhookFilterGroupListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


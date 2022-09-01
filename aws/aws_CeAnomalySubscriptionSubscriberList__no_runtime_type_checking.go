//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CeAnomalySubscriptionSubscriberList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CeAnomalySubscriptionSubscriberList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CeAnomalySubscriptionSubscriberList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CeAnomalySubscriptionSubscriberList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CeAnomalySubscriptionSubscriberList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CeAnomalySubscriptionSubscriberList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCeAnomalySubscriptionSubscriberListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


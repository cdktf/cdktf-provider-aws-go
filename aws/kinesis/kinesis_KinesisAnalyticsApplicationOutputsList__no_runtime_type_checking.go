//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package kinesis

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsList) validateGetParameters(index *float64) error {
	return nil
}

func (k *jsiiProxy_KinesisAnalyticsApplicationOutputsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_KinesisAnalyticsApplicationOutputsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewKinesisAnalyticsApplicationOutputsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


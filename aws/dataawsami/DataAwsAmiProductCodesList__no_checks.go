//go:build no_runtime_type_checking

package dataawsami

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsAmiProductCodesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsAmiProductCodesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsAmiProductCodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsAmiProductCodesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsAmiProductCodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsAmiProductCodesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


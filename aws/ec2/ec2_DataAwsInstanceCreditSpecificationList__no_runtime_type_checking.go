//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsInstanceCreditSpecificationList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsInstanceCreditSpecificationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsInstanceCreditSpecificationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsInstanceCreditSpecificationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsInstanceCreditSpecificationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsInstanceCreditSpecificationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


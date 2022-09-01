//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package vpc

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsInternetGatewayAttachmentsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsInternetGatewayAttachmentsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsInternetGatewayAttachmentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsInternetGatewayAttachmentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsInternetGatewayAttachmentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsInternetGatewayAttachmentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package apigateway

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfigurationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsApiGatewayDomainNameEndpointConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsApiGatewayDomainNameEndpointConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


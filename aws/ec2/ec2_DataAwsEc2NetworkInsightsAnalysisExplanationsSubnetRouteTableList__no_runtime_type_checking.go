//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEc2NetworkInsightsAnalysisExplanationsSubnetRouteTableListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


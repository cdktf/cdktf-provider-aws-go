//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEc2NetworkInsightsAnalysisExplanationsRouteTableListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsVpcList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsVpcList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsVpcList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsVpcList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEc2NetworkInsightsAnalysisExplanationsVpcList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEc2NetworkInsightsAnalysisExplanationsVpcListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


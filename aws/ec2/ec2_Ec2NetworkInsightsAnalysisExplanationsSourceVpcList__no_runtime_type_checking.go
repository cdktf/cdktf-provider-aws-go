//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsSourceVpcList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsSourceVpcList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsSourceVpcList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsSourceVpcList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsSourceVpcList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2NetworkInsightsAnalysisExplanationsSourceVpcListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ec2

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsIngressRouteTableList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsIngressRouteTableList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsIngressRouteTableList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsIngressRouteTableList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_Ec2NetworkInsightsAnalysisExplanationsIngressRouteTableList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEc2NetworkInsightsAnalysisExplanationsIngressRouteTableListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package codedeploy

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodedeployDeploymentGroupEc2TagFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (c *jsiiProxy_CodedeployDeploymentGroupEc2TagFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_CodedeployDeploymentGroupEc2TagFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_CodedeployDeploymentGroupEc2TagFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_CodedeployDeploymentGroupEc2TagFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_CodedeployDeploymentGroupEc2TagFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewCodedeployDeploymentGroupEc2TagFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


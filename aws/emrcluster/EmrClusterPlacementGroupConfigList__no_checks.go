//go:build no_runtime_type_checking

package emrcluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmrClusterPlacementGroupConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EmrClusterPlacementGroupConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EmrClusterPlacementGroupConfigList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EmrClusterPlacementGroupConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EmrClusterPlacementGroupConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EmrClusterPlacementGroupConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEmrClusterPlacementGroupConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


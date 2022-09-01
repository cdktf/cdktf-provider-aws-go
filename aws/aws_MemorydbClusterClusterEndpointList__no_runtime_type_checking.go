//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt aws Provider for Terraform CDK (cdktf)
package aws

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MemorydbClusterClusterEndpointList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MemorydbClusterClusterEndpointList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MemorydbClusterClusterEndpointList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MemorydbClusterClusterEndpointList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MemorydbClusterClusterEndpointList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMemorydbClusterClusterEndpointListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


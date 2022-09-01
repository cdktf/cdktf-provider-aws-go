//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package workspaces

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspacePropertiesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspacePropertiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspacePropertiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspacePropertiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsWorkspacesWorkspaceWorkspacePropertiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsWorkspacesWorkspaceWorkspacePropertiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


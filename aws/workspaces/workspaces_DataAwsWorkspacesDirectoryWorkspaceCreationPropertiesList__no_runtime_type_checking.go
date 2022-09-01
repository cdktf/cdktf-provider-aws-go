//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package workspaces

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationPropertiesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationPropertiesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationPropertiesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationPropertiesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsWorkspacesDirectoryWorkspaceCreationPropertiesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsWorkspacesDirectoryWorkspaceCreationPropertiesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


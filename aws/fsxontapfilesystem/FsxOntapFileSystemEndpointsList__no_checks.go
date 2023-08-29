// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package fsxontapfilesystem

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FsxOntapFileSystemEndpointsList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FsxOntapFileSystemEndpointsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FsxOntapFileSystemEndpointsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFsxOntapFileSystemEndpointsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


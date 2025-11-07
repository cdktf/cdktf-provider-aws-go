// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package efsfilesystem

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EfsFileSystemSizeInBytesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EfsFileSystemSizeInBytesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EfsFileSystemSizeInBytesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EfsFileSystemSizeInBytesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EfsFileSystemSizeInBytesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EfsFileSystemSizeInBytesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEfsFileSystemSizeInBytesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


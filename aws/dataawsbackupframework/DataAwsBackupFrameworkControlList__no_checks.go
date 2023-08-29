// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsbackupframework

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsBackupFrameworkControlList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsBackupFrameworkControlList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsBackupFrameworkControlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsBackupFrameworkControlList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsBackupFrameworkControlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsBackupFrameworkControlListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package backupframework

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupFrameworkControlList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupFrameworkControlList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupFrameworkControlList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupFrameworkControlList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupFrameworkControlList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupFrameworkControlList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupFrameworkControlListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


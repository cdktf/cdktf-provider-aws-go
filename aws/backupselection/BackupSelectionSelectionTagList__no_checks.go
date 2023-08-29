// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package backupselection

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupSelectionSelectionTagList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupSelectionSelectionTagList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionSelectionTagList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionSelectionTagList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionSelectionTagList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionSelectionTagList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupSelectionSelectionTagListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


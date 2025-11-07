// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package backupselection

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupSelectionConditionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BackupSelectionConditionList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupSelectionConditionList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupSelectionConditionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


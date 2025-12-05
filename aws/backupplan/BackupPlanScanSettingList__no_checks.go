// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package backupplan

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupPlanScanSettingList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BackupPlanScanSettingList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupPlanScanSettingList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupPlanScanSettingList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupPlanScanSettingList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupPlanScanSettingList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupPlanScanSettingList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupPlanScanSettingListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


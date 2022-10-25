//go:build no_runtime_type_checking

package backupselection

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupSelectionConditionStringLikeList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupSelectionConditionStringLikeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionStringLikeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionStringLikeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionStringLikeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionStringLikeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupSelectionConditionStringLikeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


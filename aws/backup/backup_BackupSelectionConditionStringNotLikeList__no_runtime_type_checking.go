//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package backup

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BackupSelectionConditionStringNotLikeList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BackupSelectionConditionStringNotLikeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionStringNotLikeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionStringNotLikeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionStringNotLikeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BackupSelectionConditionStringNotLikeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBackupSelectionConditionStringNotLikeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


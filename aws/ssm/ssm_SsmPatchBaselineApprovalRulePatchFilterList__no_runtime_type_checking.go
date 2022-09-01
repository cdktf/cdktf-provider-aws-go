//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package ssm

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmPatchBaselineApprovalRulePatchFilterList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmPatchBaselineApprovalRulePatchFilterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRulePatchFilterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRulePatchFilterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRulePatchFilterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmPatchBaselineApprovalRulePatchFilterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmPatchBaselineApprovalRulePatchFilterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


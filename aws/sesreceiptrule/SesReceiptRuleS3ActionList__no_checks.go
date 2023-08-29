// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sesreceiptrule

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SesReceiptRuleS3ActionList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SesReceiptRuleS3ActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleS3ActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleS3ActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleS3ActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SesReceiptRuleS3ActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSesReceiptRuleS3ActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


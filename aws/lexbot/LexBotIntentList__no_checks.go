// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package lexbot

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LexBotIntentList) validateGetParameters(index *float64) error {
	return nil
}

func (l *jsiiProxy_LexBotIntentList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_LexBotIntentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_LexBotIntentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_LexBotIntentList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_LexBotIntentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewLexBotIntentListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


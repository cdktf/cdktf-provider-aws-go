// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package alblistenerrule

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlbListenerRuleTransformList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AlbListenerRuleTransformList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AlbListenerRuleTransformList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleTransformList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleTransformList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleTransformList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AlbListenerRuleTransformList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAlbListenerRuleTransformListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


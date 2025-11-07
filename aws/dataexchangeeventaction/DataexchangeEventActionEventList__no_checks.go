// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataexchangeeventaction

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataexchangeEventActionEventList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataexchangeEventActionEventList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataexchangeEventActionEventList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataexchangeEventActionEventList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataexchangeEventActionEventList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataexchangeEventActionEventList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataexchangeEventActionEventList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataexchangeEventActionEventListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


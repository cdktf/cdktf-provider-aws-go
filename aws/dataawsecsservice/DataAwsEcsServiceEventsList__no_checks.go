// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsecsservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEcsServiceEventsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsEcsServiceEventsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEcsServiceEventsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceEventsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceEventsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceEventsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEcsServiceEventsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


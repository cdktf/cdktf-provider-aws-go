// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawselb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsElbHealthCheckList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsElbHealthCheckList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsElbHealthCheckList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsElbHealthCheckList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsElbHealthCheckList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsElbHealthCheckList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsElbHealthCheckListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawslb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsLbAccessLogsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsLbAccessLogsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsLbAccessLogsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbAccessLogsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbAccessLogsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbAccessLogsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsLbAccessLogsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


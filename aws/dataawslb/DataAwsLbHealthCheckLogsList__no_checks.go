// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawslb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsLbHealthCheckLogsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsLbHealthCheckLogsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsLbHealthCheckLogsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbHealthCheckLogsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbHealthCheckLogsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsLbHealthCheckLogsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsLbHealthCheckLogsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


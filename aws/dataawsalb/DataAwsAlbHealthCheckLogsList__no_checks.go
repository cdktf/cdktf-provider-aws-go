// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsalb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsAlbHealthCheckLogsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsAlbHealthCheckLogsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsAlbHealthCheckLogsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbHealthCheckLogsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbHealthCheckLogsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsAlbHealthCheckLogsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsAlbHealthCheckLogsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawslambdafunction

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfigList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsLambdaFunctionVpcConfigList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfigList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfigList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsLambdaFunctionVpcConfigList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsLambdaFunctionVpcConfigListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


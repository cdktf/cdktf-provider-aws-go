// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsgluescript

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsGlueScriptDagNodeArgsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsGlueScriptDagNodeArgsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsGlueScriptDagNodeArgsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsGlueScriptDagNodeArgsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsGlueScriptDagNodeArgsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsGlueScriptDagNodeArgsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsGlueScriptDagNodeArgsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


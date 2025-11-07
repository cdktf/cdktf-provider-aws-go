// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datazoneenvironment

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatazoneEnvironmentLastDeploymentList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatazoneEnvironmentLastDeploymentList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatazoneEnvironmentLastDeploymentList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatazoneEnvironmentLastDeploymentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatazoneEnvironmentLastDeploymentList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatazoneEnvironmentLastDeploymentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatazoneEnvironmentLastDeploymentListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


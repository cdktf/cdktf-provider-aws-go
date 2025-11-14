// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsecsservice

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsEcsServiceDeploymentsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsEcsServiceDeploymentsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsEcsServiceDeploymentsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceDeploymentsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceDeploymentsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsEcsServiceDeploymentsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsEcsServiceDeploymentsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


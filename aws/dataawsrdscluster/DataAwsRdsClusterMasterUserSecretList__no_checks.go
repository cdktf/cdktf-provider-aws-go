// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsrdscluster

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsRdsClusterMasterUserSecretList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsRdsClusterMasterUserSecretList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsRdsClusterMasterUserSecretList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsRdsClusterMasterUserSecretList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsRdsClusterMasterUserSecretList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsRdsClusterMasterUserSecretListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package finspacekxcluster

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FinspaceKxClusterDatabaseList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FinspaceKxClusterDatabaseList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FinspaceKxClusterDatabaseList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_FinspaceKxClusterDatabaseList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FinspaceKxClusterDatabaseList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FinspaceKxClusterDatabaseList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFinspaceKxClusterDatabaseListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


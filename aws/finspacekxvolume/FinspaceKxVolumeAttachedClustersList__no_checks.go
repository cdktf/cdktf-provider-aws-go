// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package finspacekxvolume

// Building without runtime type checking enabled, so all the below just return nil

func (f *jsiiProxy_FinspaceKxVolumeAttachedClustersList) validateGetParameters(index *float64) error {
	return nil
}

func (f *jsiiProxy_FinspaceKxVolumeAttachedClustersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_FinspaceKxVolumeAttachedClustersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_FinspaceKxVolumeAttachedClustersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_FinspaceKxVolumeAttachedClustersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewFinspaceKxVolumeAttachedClustersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


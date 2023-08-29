// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package emrcluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EmrClusterCoreInstanceFleetInstanceTypeConfigsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEmrClusterCoreInstanceFleetInstanceTypeConfigsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


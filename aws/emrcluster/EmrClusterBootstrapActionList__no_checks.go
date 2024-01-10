// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package emrcluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EmrClusterBootstrapActionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EmrClusterBootstrapActionList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EmrClusterBootstrapActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EmrClusterBootstrapActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EmrClusterBootstrapActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EmrClusterBootstrapActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EmrClusterBootstrapActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEmrClusterBootstrapActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


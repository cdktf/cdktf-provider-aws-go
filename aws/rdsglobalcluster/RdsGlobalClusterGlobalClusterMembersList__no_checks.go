// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package rdsglobalcluster

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RdsGlobalClusterGlobalClusterMembersList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RdsGlobalClusterGlobalClusterMembersList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RdsGlobalClusterGlobalClusterMembersList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RdsGlobalClusterGlobalClusterMembersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RdsGlobalClusterGlobalClusterMembersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RdsGlobalClusterGlobalClusterMembersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRdsGlobalClusterGlobalClusterMembersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package mskreplicator

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MskReplicatorKafkaClusterList) validateGetParameters(index *float64) error {
	return nil
}

func (m *jsiiProxy_MskReplicatorKafkaClusterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorKafkaClusterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorKafkaClusterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorKafkaClusterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_MskReplicatorKafkaClusterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewMskReplicatorKafkaClusterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


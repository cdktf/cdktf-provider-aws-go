// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package elasticachecluster

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ElasticacheClusterCacheNodesList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_ElasticacheClusterCacheNodesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ElasticacheClusterCacheNodesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ElasticacheClusterCacheNodesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ElasticacheClusterCacheNodesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewElasticacheClusterCacheNodesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


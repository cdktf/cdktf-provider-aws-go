// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dynamodbglobaltable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamodbGlobalTableReplicaList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DynamodbGlobalTableReplicaList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DynamodbGlobalTableReplicaList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableReplicaList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableReplicaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableReplicaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DynamodbGlobalTableReplicaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDynamodbGlobalTableReplicaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package batchjobqueue

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchJobQueueComputeEnvironmentOrderList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BatchJobQueueComputeEnvironmentOrderList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchJobQueueComputeEnvironmentOrderList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueComputeEnvironmentOrderList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueComputeEnvironmentOrderList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueComputeEnvironmentOrderList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueComputeEnvironmentOrderList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchJobQueueComputeEnvironmentOrderListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


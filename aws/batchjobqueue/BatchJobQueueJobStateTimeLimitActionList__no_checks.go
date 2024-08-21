// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package batchjobqueue

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BatchJobQueueJobStateTimeLimitActionList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BatchJobQueueJobStateTimeLimitActionList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BatchJobQueueJobStateTimeLimitActionList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueJobStateTimeLimitActionList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueJobStateTimeLimitActionList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueJobStateTimeLimitActionList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BatchJobQueueJobStateTimeLimitActionList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBatchJobQueueJobStateTimeLimitActionListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


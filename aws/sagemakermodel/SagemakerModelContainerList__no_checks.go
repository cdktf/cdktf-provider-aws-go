// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package sagemakermodel

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SagemakerModelContainerList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SagemakerModelContainerList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SagemakerModelContainerList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SagemakerModelContainerList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_SagemakerModelContainerList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SagemakerModelContainerList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SagemakerModelContainerList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSagemakerModelContainerListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


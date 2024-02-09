// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package osispipeline

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OsisPipelineVpcOptionsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (o *jsiiProxy_OsisPipelineVpcOptionsList) validateGetParameters(index *float64) error {
	return nil
}

func (o *jsiiProxy_OsisPipelineVpcOptionsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineVpcOptionsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineVpcOptionsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineVpcOptionsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_OsisPipelineVpcOptionsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewOsisPipelineVpcOptionsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package imagebuilderimage

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ImagebuilderImageWorkflowParameterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_ImagebuilderImageWorkflowParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_ImagebuilderImageWorkflowParameterList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageWorkflowParameterList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageWorkflowParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageWorkflowParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ImagebuilderImageWorkflowParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewImagebuilderImageWorkflowParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


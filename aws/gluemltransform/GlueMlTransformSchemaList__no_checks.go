// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package gluemltransform

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GlueMlTransformSchemaList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (g *jsiiProxy_GlueMlTransformSchemaList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GlueMlTransformSchemaList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GlueMlTransformSchemaList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GlueMlTransformSchemaList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GlueMlTransformSchemaList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGlueMlTransformSchemaListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


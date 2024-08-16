// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datazoneformtype

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatazoneFormTypeImportsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatazoneFormTypeImportsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatazoneFormTypeImportsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatazoneFormTypeImportsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatazoneFormTypeImportsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatazoneFormTypeImportsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatazoneFormTypeImportsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


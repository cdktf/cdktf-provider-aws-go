// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datazoneuserprofile

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatazoneUserProfileDetailsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DatazoneUserProfileDetailsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DatazoneUserProfileDetailsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DatazoneUserProfileDetailsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DatazoneUserProfileDetailsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DatazoneUserProfileDetailsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDatazoneUserProfileDetailsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


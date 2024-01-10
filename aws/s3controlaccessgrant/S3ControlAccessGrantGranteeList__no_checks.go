// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3controlaccessgrant

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3ControlAccessGrantGranteeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_S3ControlAccessGrantGranteeList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3ControlAccessGrantGranteeList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3ControlAccessGrantGranteeList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_S3ControlAccessGrantGranteeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3ControlAccessGrantGranteeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3ControlAccessGrantGranteeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3ControlAccessGrantGranteeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package ssmdocument

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SsmDocumentParameterList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (s *jsiiProxy_SsmDocumentParameterList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_SsmDocumentParameterList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentParameterList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentParameterList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_SsmDocumentParameterList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewSsmDocumentParameterListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


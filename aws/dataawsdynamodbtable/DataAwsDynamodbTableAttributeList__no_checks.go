// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsdynamodbtable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsDynamodbTableAttributeList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsDynamodbTableAttributeList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsDynamodbTableAttributeList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsDynamodbTableAttributeList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsDynamodbTableAttributeList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsDynamodbTableAttributeList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsDynamodbTableAttributeListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


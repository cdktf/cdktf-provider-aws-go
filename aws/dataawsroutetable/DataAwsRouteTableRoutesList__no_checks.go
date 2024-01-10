// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsroutetable

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsRouteTableRoutesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsRouteTableRoutesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsRouteTableRoutesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsRouteTableRoutesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsRouteTableRoutesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsRouteTableRoutesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsRouteTableRoutesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


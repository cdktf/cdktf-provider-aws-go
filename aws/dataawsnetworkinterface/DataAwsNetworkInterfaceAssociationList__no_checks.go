// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawsnetworkinterface

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsNetworkInterfaceAssociationList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsNetworkInterfaceAssociationList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsNetworkInterfaceAssociationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsNetworkInterfaceAssociationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsNetworkInterfaceAssociationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsNetworkInterfaceAssociationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


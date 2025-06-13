// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dsqlcluster

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DsqlClusterEncryptionDetailsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DsqlClusterEncryptionDetailsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DsqlClusterEncryptionDetailsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DsqlClusterEncryptionDetailsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DsqlClusterEncryptionDetailsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DsqlClusterEncryptionDetailsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDsqlClusterEncryptionDetailsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


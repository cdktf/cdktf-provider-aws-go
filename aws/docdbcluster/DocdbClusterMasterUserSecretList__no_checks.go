// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package docdbcluster

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DocdbClusterMasterUserSecretList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DocdbClusterMasterUserSecretList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DocdbClusterMasterUserSecretList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DocdbClusterMasterUserSecretList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DocdbClusterMasterUserSecretList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DocdbClusterMasterUserSecretList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDocdbClusterMasterUserSecretListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


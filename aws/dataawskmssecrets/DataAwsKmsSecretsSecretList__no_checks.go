// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package dataawskmssecrets

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsKmsSecretsSecretList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsKmsSecretsSecretList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsKmsSecretsSecretList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsKmsSecretsSecretList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataAwsKmsSecretsSecretList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsKmsSecretsSecretList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsKmsSecretsSecretList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsKmsSecretsSecretListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


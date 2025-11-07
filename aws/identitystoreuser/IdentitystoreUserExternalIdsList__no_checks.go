// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package identitystoreuser

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IdentitystoreUserExternalIdsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IdentitystoreUserExternalIdsList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IdentitystoreUserExternalIdsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IdentitystoreUserExternalIdsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IdentitystoreUserExternalIdsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IdentitystoreUserExternalIdsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIdentitystoreUserExternalIdsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


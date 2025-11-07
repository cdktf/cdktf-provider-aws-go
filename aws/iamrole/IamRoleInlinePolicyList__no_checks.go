// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package iamrole

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IamRoleInlinePolicyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IamRoleInlinePolicyList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IamRoleInlinePolicyList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IamRoleInlinePolicyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IamRoleInlinePolicyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IamRoleInlinePolicyList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IamRoleInlinePolicyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIamRoleInlinePolicyListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


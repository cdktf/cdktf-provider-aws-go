// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package eksaddon

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EksAddonPodIdentityAssociationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EksAddonPodIdentityAssociationList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EksAddonPodIdentityAssociationList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EksAddonPodIdentityAssociationList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EksAddonPodIdentityAssociationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EksAddonPodIdentityAssociationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EksAddonPodIdentityAssociationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEksAddonPodIdentityAssociationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


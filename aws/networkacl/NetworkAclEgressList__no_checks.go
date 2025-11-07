// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkacl

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkAclEgressList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkAclEgressList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkAclEgressList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkAclEgressList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkAclEgressList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkAclEgressList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkAclEgressList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkAclEgressListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


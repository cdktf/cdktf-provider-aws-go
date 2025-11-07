// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkmanagerconnectpeer

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkmanagerConnectPeerConfigurationList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (n *jsiiProxy_NetworkmanagerConnectPeerConfigurationList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkmanagerConnectPeerConfigurationList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerConnectPeerConfigurationList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerConnectPeerConfigurationList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkmanagerConnectPeerConfigurationList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkmanagerConnectPeerConfigurationListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


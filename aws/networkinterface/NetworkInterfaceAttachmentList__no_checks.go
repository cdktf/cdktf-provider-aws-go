// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package networkinterface

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkInterfaceAttachmentList) validateGetParameters(index *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkInterfaceAttachmentList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceAttachmentList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceAttachmentList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceAttachmentList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_NetworkInterfaceAttachmentList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewNetworkInterfaceAttachmentListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


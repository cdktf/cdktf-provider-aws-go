// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package s3outpostsendpoint

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3OutpostsEndpointNetworkInterfacesList) validateGetParameters(index *float64) error {
	return nil
}

func (s *jsiiProxy_S3OutpostsEndpointNetworkInterfacesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_S3OutpostsEndpointNetworkInterfacesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_S3OutpostsEndpointNetworkInterfacesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_S3OutpostsEndpointNetworkInterfacesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewS3OutpostsEndpointNetworkInterfacesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}


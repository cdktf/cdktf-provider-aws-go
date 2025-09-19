// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotinstancerequest


type SpotInstanceRequestEphemeralBlockDevice struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/spot_instance_request#device_name SpotInstanceRequest#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/spot_instance_request#no_device SpotInstanceRequest#no_device}.
	NoDevice interface{} `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/spot_instance_request#virtual_name SpotInstanceRequest#virtual_name}.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}


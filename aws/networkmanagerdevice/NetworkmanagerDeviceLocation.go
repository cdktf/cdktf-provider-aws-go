// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerdevice


type NetworkmanagerDeviceLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/networkmanager_device#address NetworkmanagerDevice#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/networkmanager_device#latitude NetworkmanagerDevice#latitude}.
	Latitude *string `field:"optional" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/networkmanager_device#longitude NetworkmanagerDevice#longitude}.
	Longitude *string `field:"optional" json:"longitude" yaml:"longitude"`
}


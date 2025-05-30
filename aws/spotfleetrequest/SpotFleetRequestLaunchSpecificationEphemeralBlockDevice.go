// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotfleetrequest


type SpotFleetRequestLaunchSpecificationEphemeralBlockDevice struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/spot_fleet_request#device_name SpotFleetRequest#device_name}.
	DeviceName *string `field:"required" json:"deviceName" yaml:"deviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/spot_fleet_request#virtual_name SpotFleetRequest#virtual_name}.
	VirtualName *string `field:"required" json:"virtualName" yaml:"virtualName"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package instance


type InstancePrimaryNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/instance#network_interface_id Instance#network_interface_id}.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
}


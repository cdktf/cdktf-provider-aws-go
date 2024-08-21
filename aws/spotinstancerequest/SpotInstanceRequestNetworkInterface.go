// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotinstancerequest


type SpotInstanceRequestNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/spot_instance_request#device_index SpotInstanceRequest#device_index}.
	DeviceIndex *float64 `field:"required" json:"deviceIndex" yaml:"deviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/spot_instance_request#network_interface_id SpotInstanceRequest#network_interface_id}.
	NetworkInterfaceId *string `field:"required" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/spot_instance_request#delete_on_termination SpotInstanceRequest#delete_on_termination}.
	DeleteOnTermination interface{} `field:"optional" json:"deleteOnTermination" yaml:"deleteOnTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/spot_instance_request#network_card_index SpotInstanceRequest#network_card_index}.
	NetworkCardIndex *float64 `field:"optional" json:"networkCardIndex" yaml:"networkCardIndex"`
}


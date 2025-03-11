// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package launchtemplate


type LaunchTemplateNetworkInterfacesEnaSrdSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/launch_template#ena_srd_enabled LaunchTemplate#ena_srd_enabled}.
	EnaSrdEnabled interface{} `field:"optional" json:"enaSrdEnabled" yaml:"enaSrdEnabled"`
	// ena_srd_udp_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/launch_template#ena_srd_udp_specification LaunchTemplate#ena_srd_udp_specification}
	EnaSrdUdpSpecification *LaunchTemplateNetworkInterfacesEnaSrdSpecificationEnaSrdUdpSpecification `field:"optional" json:"enaSrdUdpSpecification" yaml:"enaSrdUdpSpecification"`
}


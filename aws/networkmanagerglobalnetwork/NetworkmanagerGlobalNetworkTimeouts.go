// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerglobalnetwork


type NetworkmanagerGlobalNetworkTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/networkmanager_global_network#create NetworkmanagerGlobalNetwork#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/networkmanager_global_network#delete NetworkmanagerGlobalNetwork#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/networkmanager_global_network#update NetworkmanagerGlobalNetwork#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


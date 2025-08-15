// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument


type DataAwsNetworkmanagerCoreNetworkPolicyDocumentCoreNetworkConfigurationEdgeLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#location DataAwsNetworkmanagerCoreNetworkPolicyDocument#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#asn DataAwsNetworkmanagerCoreNetworkPolicyDocument#asn}.
	Asn *string `field:"optional" json:"asn" yaml:"asn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#inside_cidr_blocks DataAwsNetworkmanagerCoreNetworkPolicyDocument#inside_cidr_blocks}.
	InsideCidrBlocks *[]*string `field:"optional" json:"insideCidrBlocks" yaml:"insideCidrBlocks"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument


type DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsViaWithEdgeOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/data-sources/networkmanager_core_network_policy_document#edge_sets DataAwsNetworkmanagerCoreNetworkPolicyDocument#edge_sets}.
	EdgeSets interface{} `field:"optional" json:"edgeSets" yaml:"edgeSets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/data-sources/networkmanager_core_network_policy_document#use_edge DataAwsNetworkmanagerCoreNetworkPolicyDocument#use_edge}.
	UseEdge *string `field:"optional" json:"useEdge" yaml:"useEdge"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/data-sources/networkmanager_core_network_policy_document#use_edge_location DataAwsNetworkmanagerCoreNetworkPolicyDocument#use_edge_location}.
	UseEdgeLocation *string `field:"optional" json:"useEdgeLocation" yaml:"useEdgeLocation"`
}


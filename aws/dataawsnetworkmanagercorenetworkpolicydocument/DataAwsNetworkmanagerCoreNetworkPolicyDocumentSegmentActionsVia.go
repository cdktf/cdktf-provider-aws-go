// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument


type DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsVia struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/data-sources/networkmanager_core_network_policy_document#network_function_groups DataAwsNetworkmanagerCoreNetworkPolicyDocument#network_function_groups}.
	NetworkFunctionGroups *[]*string `field:"optional" json:"networkFunctionGroups" yaml:"networkFunctionGroups"`
	// with_edge_override block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/data-sources/networkmanager_core_network_policy_document#with_edge_override DataAwsNetworkmanagerCoreNetworkPolicyDocument#with_edge_override}
	WithEdgeOverride interface{} `field:"optional" json:"withEdgeOverride" yaml:"withEdgeOverride"`
}


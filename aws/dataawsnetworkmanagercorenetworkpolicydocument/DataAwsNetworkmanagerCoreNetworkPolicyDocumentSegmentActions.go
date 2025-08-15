// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument


type DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#action DataAwsNetworkmanagerCoreNetworkPolicyDocument#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#segment DataAwsNetworkmanagerCoreNetworkPolicyDocument#segment}.
	Segment *string `field:"required" json:"segment" yaml:"segment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#description DataAwsNetworkmanagerCoreNetworkPolicyDocument#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#destination_cidr_blocks DataAwsNetworkmanagerCoreNetworkPolicyDocument#destination_cidr_blocks}.
	DestinationCidrBlocks *[]*string `field:"optional" json:"destinationCidrBlocks" yaml:"destinationCidrBlocks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#destinations DataAwsNetworkmanagerCoreNetworkPolicyDocument#destinations}.
	Destinations *[]*string `field:"optional" json:"destinations" yaml:"destinations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#mode DataAwsNetworkmanagerCoreNetworkPolicyDocument#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#share_with DataAwsNetworkmanagerCoreNetworkPolicyDocument#share_with}.
	ShareWith *[]*string `field:"optional" json:"shareWith" yaml:"shareWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#share_with_except DataAwsNetworkmanagerCoreNetworkPolicyDocument#share_with_except}.
	ShareWithExcept *[]*string `field:"optional" json:"shareWithExcept" yaml:"shareWithExcept"`
	// via block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#via DataAwsNetworkmanagerCoreNetworkPolicyDocument#via}
	Via *DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsVia `field:"optional" json:"via" yaml:"via"`
	// when_sent_to block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/networkmanager_core_network_policy_document#when_sent_to DataAwsNetworkmanagerCoreNetworkPolicyDocument#when_sent_to}
	WhenSentTo *DataAwsNetworkmanagerCoreNetworkPolicyDocumentSegmentActionsWhenSentTo `field:"optional" json:"whenSentTo" yaml:"whenSentTo"`
}


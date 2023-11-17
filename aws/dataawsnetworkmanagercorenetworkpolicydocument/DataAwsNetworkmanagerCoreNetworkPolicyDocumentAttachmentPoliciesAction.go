// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument


type DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentPoliciesAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/networkmanager_core_network_policy_document#association_method DataAwsNetworkmanagerCoreNetworkPolicyDocument#association_method}.
	AssociationMethod *string `field:"required" json:"associationMethod" yaml:"associationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/networkmanager_core_network_policy_document#require_acceptance DataAwsNetworkmanagerCoreNetworkPolicyDocument#require_acceptance}.
	RequireAcceptance interface{} `field:"optional" json:"requireAcceptance" yaml:"requireAcceptance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/networkmanager_core_network_policy_document#segment DataAwsNetworkmanagerCoreNetworkPolicyDocument#segment}.
	Segment *string `field:"optional" json:"segment" yaml:"segment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/networkmanager_core_network_policy_document#tag_value_of_key DataAwsNetworkmanagerCoreNetworkPolicyDocument#tag_value_of_key}.
	TagValueOfKey *string `field:"optional" json:"tagValueOfKey" yaml:"tagValueOfKey"`
}


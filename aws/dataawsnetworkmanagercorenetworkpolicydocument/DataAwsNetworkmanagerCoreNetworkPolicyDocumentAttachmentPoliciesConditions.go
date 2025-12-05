// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagercorenetworkpolicydocument


type DataAwsNetworkmanagerCoreNetworkPolicyDocumentAttachmentPoliciesConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/data-sources/networkmanager_core_network_policy_document#type DataAwsNetworkmanagerCoreNetworkPolicyDocument#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/data-sources/networkmanager_core_network_policy_document#key DataAwsNetworkmanagerCoreNetworkPolicyDocument#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/data-sources/networkmanager_core_network_policy_document#operator DataAwsNetworkmanagerCoreNetworkPolicyDocument#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/data-sources/networkmanager_core_network_policy_document#value DataAwsNetworkmanagerCoreNetworkPolicyDocument#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}


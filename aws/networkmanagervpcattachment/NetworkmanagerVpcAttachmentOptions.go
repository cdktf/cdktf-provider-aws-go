// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagervpcattachment


type NetworkmanagerVpcAttachmentOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/networkmanager_vpc_attachment#appliance_mode_support NetworkmanagerVpcAttachment#appliance_mode_support}.
	ApplianceModeSupport interface{} `field:"optional" json:"applianceModeSupport" yaml:"applianceModeSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/networkmanager_vpc_attachment#dns_support NetworkmanagerVpcAttachment#dns_support}.
	DnsSupport interface{} `field:"optional" json:"dnsSupport" yaml:"dnsSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/networkmanager_vpc_attachment#ipv6_support NetworkmanagerVpcAttachment#ipv6_support}.
	Ipv6Support interface{} `field:"optional" json:"ipv6Support" yaml:"ipv6Support"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/networkmanager_vpc_attachment#security_group_referencing_support NetworkmanagerVpcAttachment#security_group_referencing_support}.
	SecurityGroupReferencingSupport interface{} `field:"optional" json:"securityGroupReferencingSupport" yaml:"securityGroupReferencingSupport"`
}


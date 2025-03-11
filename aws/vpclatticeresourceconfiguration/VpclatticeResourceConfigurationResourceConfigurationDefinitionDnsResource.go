// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpclatticeresourceconfiguration


type VpclatticeResourceConfigurationResourceConfigurationDefinitionDnsResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/vpclattice_resource_configuration#domain_name VpclatticeResourceConfiguration#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/vpclattice_resource_configuration#ip_address_type VpclatticeResourceConfiguration#ip_address_type}.
	IpAddressType *string `field:"required" json:"ipAddressType" yaml:"ipAddressType"`
}


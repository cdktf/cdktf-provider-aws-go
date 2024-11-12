// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package spotinstancerequest


type SpotInstanceRequestPrivateDnsNameOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/spot_instance_request#enable_resource_name_dns_aaaa_record SpotInstanceRequest#enable_resource_name_dns_aaaa_record}.
	EnableResourceNameDnsAaaaRecord interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecord" yaml:"enableResourceNameDnsAaaaRecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/spot_instance_request#enable_resource_name_dns_a_record SpotInstanceRequest#enable_resource_name_dns_a_record}.
	EnableResourceNameDnsARecord interface{} `field:"optional" json:"enableResourceNameDnsARecord" yaml:"enableResourceNameDnsARecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/spot_instance_request#hostname_type SpotInstanceRequest#hostname_type}.
	HostnameType *string `field:"optional" json:"hostnameType" yaml:"hostnameType"`
}


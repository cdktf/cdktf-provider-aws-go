// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcdhcpoptions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcDhcpOptionsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_dhcp_options#domain_name VpcDhcpOptions#domain_name}.
	DomainName *string `field:"optional" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_dhcp_options#domain_name_servers VpcDhcpOptions#domain_name_servers}.
	DomainNameServers *[]*string `field:"optional" json:"domainNameServers" yaml:"domainNameServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_dhcp_options#id VpcDhcpOptions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_dhcp_options#ipv6_address_preferred_lease_time VpcDhcpOptions#ipv6_address_preferred_lease_time}.
	Ipv6AddressPreferredLeaseTime *string `field:"optional" json:"ipv6AddressPreferredLeaseTime" yaml:"ipv6AddressPreferredLeaseTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_dhcp_options#netbios_name_servers VpcDhcpOptions#netbios_name_servers}.
	NetbiosNameServers *[]*string `field:"optional" json:"netbiosNameServers" yaml:"netbiosNameServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_dhcp_options#netbios_node_type VpcDhcpOptions#netbios_node_type}.
	NetbiosNodeType *string `field:"optional" json:"netbiosNodeType" yaml:"netbiosNodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_dhcp_options#ntp_servers VpcDhcpOptions#ntp_servers}.
	NtpServers *[]*string `field:"optional" json:"ntpServers" yaml:"ntpServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_dhcp_options#tags VpcDhcpOptions#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/vpc_dhcp_options#tags_all VpcDhcpOptions#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}


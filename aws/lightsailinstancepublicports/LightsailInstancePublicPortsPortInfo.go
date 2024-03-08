// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lightsailinstancepublicports


type LightsailInstancePublicPortsPortInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/lightsail_instance_public_ports#from_port LightsailInstancePublicPorts#from_port}.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/lightsail_instance_public_ports#protocol LightsailInstancePublicPorts#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/lightsail_instance_public_ports#to_port LightsailInstancePublicPorts#to_port}.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/lightsail_instance_public_ports#cidr_list_aliases LightsailInstancePublicPorts#cidr_list_aliases}.
	CidrListAliases *[]*string `field:"optional" json:"cidrListAliases" yaml:"cidrListAliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/lightsail_instance_public_ports#cidrs LightsailInstancePublicPorts#cidrs}.
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.40.0/docs/resources/lightsail_instance_public_ports#ipv6_cidrs LightsailInstancePublicPorts#ipv6_cidrs}.
	Ipv6Cidrs *[]*string `field:"optional" json:"ipv6Cidrs" yaml:"ipv6Cidrs"`
}


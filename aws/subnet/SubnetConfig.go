// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subnet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SubnetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#vpc_id Subnet#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#assign_ipv6_address_on_creation Subnet#assign_ipv6_address_on_creation}.
	AssignIpv6AddressOnCreation interface{} `field:"optional" json:"assignIpv6AddressOnCreation" yaml:"assignIpv6AddressOnCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#availability_zone Subnet#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#availability_zone_id Subnet#availability_zone_id}.
	AvailabilityZoneId *string `field:"optional" json:"availabilityZoneId" yaml:"availabilityZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#cidr_block Subnet#cidr_block}.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#customer_owned_ipv4_pool Subnet#customer_owned_ipv4_pool}.
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#enable_dns64 Subnet#enable_dns64}.
	EnableDns64 interface{} `field:"optional" json:"enableDns64" yaml:"enableDns64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#enable_lni_at_device_index Subnet#enable_lni_at_device_index}.
	EnableLniAtDeviceIndex *float64 `field:"optional" json:"enableLniAtDeviceIndex" yaml:"enableLniAtDeviceIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#enable_resource_name_dns_aaaa_record_on_launch Subnet#enable_resource_name_dns_aaaa_record_on_launch}.
	EnableResourceNameDnsAaaaRecordOnLaunch interface{} `field:"optional" json:"enableResourceNameDnsAaaaRecordOnLaunch" yaml:"enableResourceNameDnsAaaaRecordOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#enable_resource_name_dns_a_record_on_launch Subnet#enable_resource_name_dns_a_record_on_launch}.
	EnableResourceNameDnsARecordOnLaunch interface{} `field:"optional" json:"enableResourceNameDnsARecordOnLaunch" yaml:"enableResourceNameDnsARecordOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#id Subnet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#ipv6_cidr_block Subnet#ipv6_cidr_block}.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#ipv6_native Subnet#ipv6_native}.
	Ipv6Native interface{} `field:"optional" json:"ipv6Native" yaml:"ipv6Native"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#map_customer_owned_ip_on_launch Subnet#map_customer_owned_ip_on_launch}.
	MapCustomerOwnedIpOnLaunch interface{} `field:"optional" json:"mapCustomerOwnedIpOnLaunch" yaml:"mapCustomerOwnedIpOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#map_public_ip_on_launch Subnet#map_public_ip_on_launch}.
	MapPublicIpOnLaunch interface{} `field:"optional" json:"mapPublicIpOnLaunch" yaml:"mapPublicIpOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#outpost_arn Subnet#outpost_arn}.
	OutpostArn *string `field:"optional" json:"outpostArn" yaml:"outpostArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#private_dns_hostname_type_on_launch Subnet#private_dns_hostname_type_on_launch}.
	PrivateDnsHostnameTypeOnLaunch *string `field:"optional" json:"privateDnsHostnameTypeOnLaunch" yaml:"privateDnsHostnameTypeOnLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#tags Subnet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#tags_all Subnet#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/subnet#timeouts Subnet#timeouts}
	Timeouts *SubnetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsvpngateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsVpnGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/data-sources/vpn_gateway#amazon_side_asn DataAwsVpnGateway#amazon_side_asn}.
	AmazonSideAsn *string `field:"optional" json:"amazonSideAsn" yaml:"amazonSideAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/data-sources/vpn_gateway#attached_vpc_id DataAwsVpnGateway#attached_vpc_id}.
	AttachedVpcId *string `field:"optional" json:"attachedVpcId" yaml:"attachedVpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/data-sources/vpn_gateway#availability_zone DataAwsVpnGateway#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/data-sources/vpn_gateway#filter DataAwsVpnGateway#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/data-sources/vpn_gateway#id DataAwsVpnGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/data-sources/vpn_gateway#state DataAwsVpnGateway#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/data-sources/vpn_gateway#tags DataAwsVpnGateway#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/data-sources/vpn_gateway#timeouts DataAwsVpnGateway#timeouts}
	Timeouts *DataAwsVpnGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


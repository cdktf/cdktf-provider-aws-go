// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsvpcpeeringconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsVpcPeeringConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#cidr_block DataAwsVpcPeeringConnection#cidr_block}.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#filter DataAwsVpcPeeringConnection#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#id DataAwsVpcPeeringConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#owner_id DataAwsVpcPeeringConnection#owner_id}.
	OwnerId *string `field:"optional" json:"ownerId" yaml:"ownerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#peer_cidr_block DataAwsVpcPeeringConnection#peer_cidr_block}.
	PeerCidrBlock *string `field:"optional" json:"peerCidrBlock" yaml:"peerCidrBlock"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#peer_owner_id DataAwsVpcPeeringConnection#peer_owner_id}.
	PeerOwnerId *string `field:"optional" json:"peerOwnerId" yaml:"peerOwnerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#peer_region DataAwsVpcPeeringConnection#peer_region}.
	PeerRegion *string `field:"optional" json:"peerRegion" yaml:"peerRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#peer_vpc_id DataAwsVpcPeeringConnection#peer_vpc_id}.
	PeerVpcId *string `field:"optional" json:"peerVpcId" yaml:"peerVpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#region DataAwsVpcPeeringConnection#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#status DataAwsVpcPeeringConnection#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#tags DataAwsVpcPeeringConnection#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#timeouts DataAwsVpcPeeringConnection#timeouts}
	Timeouts *DataAwsVpcPeeringConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.48.0/docs/data-sources/vpc_peering_connection#vpc_id DataAwsVpcPeeringConnection#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}


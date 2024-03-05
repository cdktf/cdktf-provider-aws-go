// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcpeeringconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcPeeringConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/vpc_peering_connection#peer_vpc_id VpcPeeringConnection#peer_vpc_id}.
	PeerVpcId *string `field:"required" json:"peerVpcId" yaml:"peerVpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/vpc_peering_connection#vpc_id VpcPeeringConnection#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// accepter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/vpc_peering_connection#accepter VpcPeeringConnection#accepter}
	Accepter *VpcPeeringConnectionAccepter `field:"optional" json:"accepter" yaml:"accepter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/vpc_peering_connection#auto_accept VpcPeeringConnection#auto_accept}.
	AutoAccept interface{} `field:"optional" json:"autoAccept" yaml:"autoAccept"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/vpc_peering_connection#id VpcPeeringConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/vpc_peering_connection#peer_owner_id VpcPeeringConnection#peer_owner_id}.
	PeerOwnerId *string `field:"optional" json:"peerOwnerId" yaml:"peerOwnerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/vpc_peering_connection#peer_region VpcPeeringConnection#peer_region}.
	PeerRegion *string `field:"optional" json:"peerRegion" yaml:"peerRegion"`
	// requester block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/vpc_peering_connection#requester VpcPeeringConnection#requester}
	Requester *VpcPeeringConnectionRequester `field:"optional" json:"requester" yaml:"requester"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/vpc_peering_connection#tags VpcPeeringConnection#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/vpc_peering_connection#tags_all VpcPeeringConnection#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/vpc_peering_connection#timeouts VpcPeeringConnection#timeouts}
	Timeouts *VpcPeeringConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


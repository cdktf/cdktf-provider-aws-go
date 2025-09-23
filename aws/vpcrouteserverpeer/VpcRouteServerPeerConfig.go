// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcrouteserverpeer

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcRouteServerPeerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/vpc_route_server_peer#peer_address VpcRouteServerPeer#peer_address}.
	PeerAddress *string `field:"required" json:"peerAddress" yaml:"peerAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/vpc_route_server_peer#route_server_endpoint_id VpcRouteServerPeer#route_server_endpoint_id}.
	RouteServerEndpointId *string `field:"required" json:"routeServerEndpointId" yaml:"routeServerEndpointId"`
	// bgp_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/vpc_route_server_peer#bgp_options VpcRouteServerPeer#bgp_options}
	BgpOptions interface{} `field:"optional" json:"bgpOptions" yaml:"bgpOptions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/vpc_route_server_peer#region VpcRouteServerPeer#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/vpc_route_server_peer#tags VpcRouteServerPeer#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/vpc_route_server_peer#timeouts VpcRouteServerPeer#timeouts}
	Timeouts *VpcRouteServerPeerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


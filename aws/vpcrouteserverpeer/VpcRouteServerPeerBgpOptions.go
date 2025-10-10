// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcrouteserverpeer


type VpcRouteServerPeerBgpOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/vpc_route_server_peer#peer_asn VpcRouteServerPeer#peer_asn}.
	PeerAsn *float64 `field:"required" json:"peerAsn" yaml:"peerAsn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/vpc_route_server_peer#peer_liveness_detection VpcRouteServerPeer#peer_liveness_detection}.
	PeerLivenessDetection *string `field:"optional" json:"peerLivenessDetection" yaml:"peerLivenessDetection"`
}


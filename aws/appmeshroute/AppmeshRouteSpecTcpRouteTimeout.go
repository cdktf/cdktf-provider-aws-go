// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecTcpRouteTimeout struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/appmesh_route#idle AppmeshRoute#idle}
	Idle *AppmeshRouteSpecTcpRouteTimeoutIdle `field:"optional" json:"idle" yaml:"idle"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecHttpRouteTimeout struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/appmesh_route#idle AppmeshRoute#idle}
	Idle *AppmeshRouteSpecHttpRouteTimeoutIdle `field:"optional" json:"idle" yaml:"idle"`
	// per_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/appmesh_route#per_request AppmeshRoute#per_request}
	PerRequest *AppmeshRouteSpecHttpRouteTimeoutPerRequest `field:"optional" json:"perRequest" yaml:"perRequest"`
}


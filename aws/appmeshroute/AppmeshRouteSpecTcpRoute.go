// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecTcpRoute struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/appmesh_route#action AppmeshRoute#action}
	Action *AppmeshRouteSpecTcpRouteAction `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/appmesh_route#match AppmeshRoute#match}
	Match *AppmeshRouteSpecTcpRouteMatch `field:"optional" json:"match" yaml:"match"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.68.0/docs/resources/appmesh_route#timeout AppmeshRoute#timeout}
	Timeout *AppmeshRouteSpecTcpRouteTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}


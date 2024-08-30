// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecHttp2Route struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/appmesh_route#action AppmeshRoute#action}
	Action *AppmeshRouteSpecHttp2RouteAction `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/appmesh_route#match AppmeshRoute#match}
	Match *AppmeshRouteSpecHttp2RouteMatch `field:"required" json:"match" yaml:"match"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/appmesh_route#retry_policy AppmeshRoute#retry_policy}
	RetryPolicy *AppmeshRouteSpecHttp2RouteRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/appmesh_route#timeout AppmeshRoute#timeout}
	Timeout *AppmeshRouteSpecHttp2RouteTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}


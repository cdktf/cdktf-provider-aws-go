// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecHttpRoute struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/appmesh_route#action AppmeshRoute#action}
	Action *AppmeshRouteSpecHttpRouteAction `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/appmesh_route#match AppmeshRoute#match}
	Match *AppmeshRouteSpecHttpRouteMatch `field:"required" json:"match" yaml:"match"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/appmesh_route#retry_policy AppmeshRoute#retry_policy}
	RetryPolicy *AppmeshRouteSpecHttpRouteRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/appmesh_route#timeout AppmeshRoute#timeout}
	Timeout *AppmeshRouteSpecHttpRouteTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}


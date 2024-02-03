// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshroute


type AppmeshRouteSpecGrpcRoute struct {
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/appmesh_route#action AppmeshRoute#action}
	Action *AppmeshRouteSpecGrpcRouteAction `field:"required" json:"action" yaml:"action"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/appmesh_route#match AppmeshRoute#match}
	Match *AppmeshRouteSpecGrpcRouteMatch `field:"optional" json:"match" yaml:"match"`
	// retry_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/appmesh_route#retry_policy AppmeshRoute#retry_policy}
	RetryPolicy *AppmeshRouteSpecGrpcRouteRetryPolicy `field:"optional" json:"retryPolicy" yaml:"retryPolicy"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/appmesh_route#timeout AppmeshRoute#timeout}
	Timeout *AppmeshRouteSpecGrpcRouteTimeout `field:"optional" json:"timeout" yaml:"timeout"`
}


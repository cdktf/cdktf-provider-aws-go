// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudfrontvpcorigin


type CloudfrontVpcOriginVpcOriginEndpointConfigOriginSslProtocols struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_vpc_origin#items CloudfrontVpcOrigin#items}.
	Items *[]*string `field:"required" json:"items" yaml:"items"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/cloudfront_vpc_origin#quantity CloudfrontVpcOrigin#quantity}.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
}


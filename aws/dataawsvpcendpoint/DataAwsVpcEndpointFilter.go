// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsvpcendpoint


type DataAwsVpcEndpointFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/data-sources/vpc_endpoint#name DataAwsVpcEndpoint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/data-sources/vpc_endpoint#values DataAwsVpcEndpoint#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


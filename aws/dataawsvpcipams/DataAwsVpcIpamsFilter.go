// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsvpcipams


type DataAwsVpcIpamsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/data-sources/vpc_ipams#name DataAwsVpcIpams#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.1/docs/data-sources/vpc_ipams#values DataAwsVpcIpams#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


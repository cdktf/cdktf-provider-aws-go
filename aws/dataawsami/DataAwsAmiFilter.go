// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsami


type DataAwsAmiFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/data-sources/ami#name DataAwsAmi#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/data-sources/ami#values DataAwsAmi#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


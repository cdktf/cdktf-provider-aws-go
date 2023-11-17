// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsregions


type DataAwsRegionsFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/regions#name DataAwsRegions#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/regions#values DataAwsRegions#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


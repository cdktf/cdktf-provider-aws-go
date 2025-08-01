// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsavailabilityzones


type DataAwsAvailabilityZonesFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/availability_zones#name DataAwsAvailabilityZones#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/data-sources/availability_zones#values DataAwsAvailabilityZones#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsrdsengineversion


type DataAwsRdsEngineVersionFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/rds_engine_version#name DataAwsRdsEngineVersion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/rds_engine_version#values DataAwsRdsEngineVersion#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


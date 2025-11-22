// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdscustomdbengineversion


type RdsCustomDbEngineVersionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/rds_custom_db_engine_version#create RdsCustomDbEngineVersion#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/rds_custom_db_engine_version#delete RdsCustomDbEngineVersion#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/rds_custom_db_engine_version#update RdsCustomDbEngineVersion#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


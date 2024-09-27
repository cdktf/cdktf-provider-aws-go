// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxdataview


type FinspaceKxDataviewSegmentConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/finspace_kx_dataview#db_paths FinspaceKxDataview#db_paths}.
	DbPaths *[]*string `field:"required" json:"dbPaths" yaml:"dbPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/finspace_kx_dataview#volume_name FinspaceKxDataview#volume_name}.
	VolumeName *string `field:"required" json:"volumeName" yaml:"volumeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/finspace_kx_dataview#on_demand FinspaceKxDataview#on_demand}.
	OnDemand interface{} `field:"optional" json:"onDemand" yaml:"onDemand"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetColumnGroupsGeoSpatialColumnGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/quicksight_data_set#columns QuicksightDataSet#columns}.
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/quicksight_data_set#country_code QuicksightDataSet#country_code}.
	CountryCode *string `field:"required" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/quicksight_data_set#name QuicksightDataSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}


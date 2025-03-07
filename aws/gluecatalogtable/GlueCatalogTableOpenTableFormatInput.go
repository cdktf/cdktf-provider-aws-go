// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable


type GlueCatalogTableOpenTableFormatInput struct {
	// iceberg_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/glue_catalog_table#iceberg_input GlueCatalogTable#iceberg_input}
	IcebergInput *GlueCatalogTableOpenTableFormatInputIcebergInput `field:"required" json:"icebergInput" yaml:"icebergInput"`
}


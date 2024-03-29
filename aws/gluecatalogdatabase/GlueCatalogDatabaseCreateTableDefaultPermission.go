// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogdatabase


type GlueCatalogDatabaseCreateTableDefaultPermission struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/glue_catalog_database#permissions GlueCatalogDatabase#permissions}.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
	// principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/glue_catalog_database#principal GlueCatalogDatabase#principal}
	Principal *GlueCatalogDatabaseCreateTableDefaultPermissionPrincipal `field:"optional" json:"principal" yaml:"principal"`
}


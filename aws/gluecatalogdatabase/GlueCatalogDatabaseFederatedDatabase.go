// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogdatabase


type GlueCatalogDatabaseFederatedDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/glue_catalog_database#connection_name GlueCatalogDatabase#connection_name}.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.44.0/docs/resources/glue_catalog_database#identifier GlueCatalogDatabase#identifier}.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
}


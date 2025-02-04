// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appfabricingestiondestination


type AppfabricIngestionDestinationProcessingConfigurationAuditLog struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/appfabric_ingestion_destination#format AppfabricIngestionDestination#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.85.0/docs/resources/appfabric_ingestion_destination#schema AppfabricIngestionDestination#schema}.
	Schema *string `field:"required" json:"schema" yaml:"schema"`
}


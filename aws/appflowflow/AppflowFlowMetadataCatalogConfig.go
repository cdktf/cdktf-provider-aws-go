// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appflowflow


type AppflowFlowMetadataCatalogConfig struct {
	// glue_data_catalog block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/appflow_flow#glue_data_catalog AppflowFlow#glue_data_catalog}
	GlueDataCatalog *AppflowFlowMetadataCatalogConfigGlueDataCatalog `field:"optional" json:"glueDataCatalog" yaml:"glueDataCatalog"`
}


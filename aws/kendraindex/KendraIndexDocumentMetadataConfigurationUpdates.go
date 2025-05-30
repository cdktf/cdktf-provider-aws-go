// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendraindex


type KendraIndexDocumentMetadataConfigurationUpdates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/kendra_index#name KendraIndex#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/kendra_index#type KendraIndex#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// relevance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/kendra_index#relevance KendraIndex#relevance}
	Relevance *KendraIndexDocumentMetadataConfigurationUpdatesRelevance `field:"optional" json:"relevance" yaml:"relevance"`
	// search block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/kendra_index#search KendraIndex#search}
	Search *KendraIndexDocumentMetadataConfigurationUpdatesSearch `field:"optional" json:"search" yaml:"search"`
}


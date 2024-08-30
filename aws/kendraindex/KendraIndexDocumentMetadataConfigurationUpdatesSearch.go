// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kendraindex


type KendraIndexDocumentMetadataConfigurationUpdatesSearch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/kendra_index#displayable KendraIndex#displayable}.
	Displayable interface{} `field:"optional" json:"displayable" yaml:"displayable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/kendra_index#facetable KendraIndex#facetable}.
	Facetable interface{} `field:"optional" json:"facetable" yaml:"facetable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/kendra_index#searchable KendraIndex#searchable}.
	Searchable interface{} `field:"optional" json:"searchable" yaml:"searchable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/kendra_index#sortable KendraIndex#sortable}.
	Sortable interface{} `field:"optional" json:"sortable" yaml:"sortable"`
}


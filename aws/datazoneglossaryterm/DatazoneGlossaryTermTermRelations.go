// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datazoneglossaryterm


type DatazoneGlossaryTermTermRelations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/datazone_glossary_term#classifies DatazoneGlossaryTerm#classifies}.
	Classifies *[]*string `field:"optional" json:"classifies" yaml:"classifies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/datazone_glossary_term#is_a DatazoneGlossaryTerm#is_a}.
	IsA *[]*string `field:"optional" json:"isA" yaml:"isA"`
}


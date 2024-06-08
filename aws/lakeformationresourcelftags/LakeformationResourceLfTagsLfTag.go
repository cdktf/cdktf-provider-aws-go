// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationresourcelftags


type LakeformationResourceLfTagsLfTag struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/lakeformation_resource_lf_tags#key LakeformationResourceLfTags#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/lakeformation_resource_lf_tags#value LakeformationResourceLfTags#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/lakeformation_resource_lf_tags#catalog_id LakeformationResourceLfTags#catalog_id}.
	CatalogId *string `field:"optional" json:"catalogId" yaml:"catalogId"`
}


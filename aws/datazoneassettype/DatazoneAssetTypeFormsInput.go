// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datazoneassettype


type DatazoneAssetTypeFormsInput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/datazone_asset_type#map_block_key DatazoneAssetType#map_block_key}.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/datazone_asset_type#type_identifier DatazoneAssetType#type_identifier}.
	TypeIdentifier *string `field:"required" json:"typeIdentifier" yaml:"typeIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/datazone_asset_type#type_revision DatazoneAssetType#type_revision}.
	TypeRevision *string `field:"required" json:"typeRevision" yaml:"typeRevision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/datazone_asset_type#required DatazoneAssetType#required}.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
}


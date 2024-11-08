// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxdatarepositoryassociation


type FsxDataRepositoryAssociationS3 struct {
	// auto_export_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/fsx_data_repository_association#auto_export_policy FsxDataRepositoryAssociation#auto_export_policy}
	AutoExportPolicy *FsxDataRepositoryAssociationS3AutoExportPolicy `field:"optional" json:"autoExportPolicy" yaml:"autoExportPolicy"`
	// auto_import_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/fsx_data_repository_association#auto_import_policy FsxDataRepositoryAssociation#auto_import_policy}
	AutoImportPolicy *FsxDataRepositoryAssociationS3AutoImportPolicy `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
}


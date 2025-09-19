// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrpublicrepository


type EcrpublicRepositoryCatalogData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ecrpublic_repository#about_text EcrpublicRepository#about_text}.
	AboutText *string `field:"optional" json:"aboutText" yaml:"aboutText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ecrpublic_repository#architectures EcrpublicRepository#architectures}.
	Architectures *[]*string `field:"optional" json:"architectures" yaml:"architectures"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ecrpublic_repository#description EcrpublicRepository#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ecrpublic_repository#logo_image_blob EcrpublicRepository#logo_image_blob}.
	LogoImageBlob *string `field:"optional" json:"logoImageBlob" yaml:"logoImageBlob"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ecrpublic_repository#operating_systems EcrpublicRepository#operating_systems}.
	OperatingSystems *[]*string `field:"optional" json:"operatingSystems" yaml:"operatingSystems"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/ecrpublic_repository#usage_text EcrpublicRepository#usage_text}.
	UsageText *string `field:"optional" json:"usageText" yaml:"usageText"`
}


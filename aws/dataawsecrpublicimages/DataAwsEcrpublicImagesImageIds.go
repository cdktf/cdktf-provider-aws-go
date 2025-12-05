// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsecrpublicimages


type DataAwsEcrpublicImagesImageIds struct {
	// Image digest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/data-sources/ecrpublic_images#image_digest DataAwsEcrpublicImages#image_digest}
	ImageDigest *string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// Image tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/data-sources/ecrpublic_images#image_tag DataAwsEcrpublicImages#image_tag}
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}


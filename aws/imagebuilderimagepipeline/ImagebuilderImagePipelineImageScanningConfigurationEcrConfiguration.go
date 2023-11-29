// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilderimagepipeline


type ImagebuilderImagePipelineImageScanningConfigurationEcrConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/imagebuilder_image_pipeline#container_tags ImagebuilderImagePipeline#container_tags}.
	ContainerTags *[]*string `field:"optional" json:"containerTags" yaml:"containerTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/imagebuilder_image_pipeline#repository_name ImagebuilderImagePipeline#repository_name}.
	RepositoryName *string `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}


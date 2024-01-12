// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceSourceConfigurationImageRepository struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/apprunner_service#image_identifier ApprunnerService#image_identifier}.
	ImageIdentifier *string `field:"required" json:"imageIdentifier" yaml:"imageIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/apprunner_service#image_repository_type ApprunnerService#image_repository_type}.
	ImageRepositoryType *string `field:"required" json:"imageRepositoryType" yaml:"imageRepositoryType"`
	// image_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/apprunner_service#image_configuration ApprunnerService#image_configuration}
	ImageConfiguration *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}


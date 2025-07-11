// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerappimageconfig


type SagemakerAppImageConfigKernelGatewayImageConfigKernelSpec struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/sagemaker_app_image_config#name SagemakerAppImageConfig#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/sagemaker_app_image_config#display_name SagemakerAppImageConfig#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}


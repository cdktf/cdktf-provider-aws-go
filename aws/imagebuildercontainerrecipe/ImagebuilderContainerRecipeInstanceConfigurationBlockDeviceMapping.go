// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuildercontainerrecipe


type ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/imagebuilder_container_recipe#device_name ImagebuilderContainerRecipe#device_name}.
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// ebs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/imagebuilder_container_recipe#ebs ImagebuilderContainerRecipe#ebs}
	Ebs *ImagebuilderContainerRecipeInstanceConfigurationBlockDeviceMappingEbs `field:"optional" json:"ebs" yaml:"ebs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/imagebuilder_container_recipe#no_device ImagebuilderContainerRecipe#no_device}.
	NoDevice interface{} `field:"optional" json:"noDevice" yaml:"noDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.43.0/docs/resources/imagebuilder_container_recipe#virtual_name ImagebuilderContainerRecipe#virtual_name}.
	VirtualName *string `field:"optional" json:"virtualName" yaml:"virtualName"`
}


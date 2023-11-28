// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksstaticweblayer


type OpsworksStaticWebLayerLoadBasedAutoScaling struct {
	// downscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/opsworks_static_web_layer#downscaling OpsworksStaticWebLayer#downscaling}
	Downscaling *OpsworksStaticWebLayerLoadBasedAutoScalingDownscaling `field:"optional" json:"downscaling" yaml:"downscaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/opsworks_static_web_layer#enable OpsworksStaticWebLayer#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// upscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.27.0/docs/resources/opsworks_static_web_layer#upscaling OpsworksStaticWebLayer#upscaling}
	Upscaling *OpsworksStaticWebLayerLoadBasedAutoScalingUpscaling `field:"optional" json:"upscaling" yaml:"upscaling"`
}


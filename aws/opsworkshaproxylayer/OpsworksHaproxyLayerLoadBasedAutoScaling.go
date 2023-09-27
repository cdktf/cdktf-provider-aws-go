// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworkshaproxylayer


type OpsworksHaproxyLayerLoadBasedAutoScaling struct {
	// downscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/opsworks_haproxy_layer#downscaling OpsworksHaproxyLayer#downscaling}
	Downscaling *OpsworksHaproxyLayerLoadBasedAutoScalingDownscaling `field:"optional" json:"downscaling" yaml:"downscaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/opsworks_haproxy_layer#enable OpsworksHaproxyLayer#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// upscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/opsworks_haproxy_layer#upscaling OpsworksHaproxyLayer#upscaling}
	Upscaling *OpsworksHaproxyLayerLoadBasedAutoScalingUpscaling `field:"optional" json:"upscaling" yaml:"upscaling"`
}


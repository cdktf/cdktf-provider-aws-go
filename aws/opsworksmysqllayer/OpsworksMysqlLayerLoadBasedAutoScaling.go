// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksmysqllayer


type OpsworksMysqlLayerLoadBasedAutoScaling struct {
	// downscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/opsworks_mysql_layer#downscaling OpsworksMysqlLayer#downscaling}
	Downscaling *OpsworksMysqlLayerLoadBasedAutoScalingDownscaling `field:"optional" json:"downscaling" yaml:"downscaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/opsworks_mysql_layer#enable OpsworksMysqlLayer#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// upscaling block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/opsworks_mysql_layer#upscaling OpsworksMysqlLayer#upscaling}
	Upscaling *OpsworksMysqlLayerLoadBasedAutoScalingUpscaling `field:"optional" json:"upscaling" yaml:"upscaling"`
}


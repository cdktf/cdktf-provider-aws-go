// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksphpapplayer


type OpsworksPhpAppLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/opsworks_php_app_layer#enabled OpsworksPhpAppLayer#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/opsworks_php_app_layer#log_streams OpsworksPhpAppLayer#log_streams}
	LogStreams interface{} `field:"optional" json:"logStreams" yaml:"logStreams"`
}


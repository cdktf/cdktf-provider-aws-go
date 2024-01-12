// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksecsclusterlayer


type OpsworksEcsClusterLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_ecs_cluster_layer#enabled OpsworksEcsClusterLayer#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/opsworks_ecs_cluster_layer#log_streams OpsworksEcsClusterLayer#log_streams}
	LogStreams interface{} `field:"optional" json:"logStreams" yaml:"logStreams"`
}


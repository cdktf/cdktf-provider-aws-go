// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksmemcachedlayer


type OpsworksMemcachedLayerCloudwatchConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/opsworks_memcached_layer#enabled OpsworksMemcachedLayer#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// log_streams block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/opsworks_memcached_layer#log_streams OpsworksMemcachedLayer#log_streams}
	LogStreams interface{} `field:"optional" json:"logStreams" yaml:"logStreams"`
}


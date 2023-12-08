// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworkshaproxylayer


type OpsworksHaproxyLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_haproxy_layer#file OpsworksHaproxyLayer#file}.
	File *string `field:"required" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_haproxy_layer#log_group_name OpsworksHaproxyLayer#log_group_name}.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_haproxy_layer#batch_count OpsworksHaproxyLayer#batch_count}.
	BatchCount *float64 `field:"optional" json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_haproxy_layer#batch_size OpsworksHaproxyLayer#batch_size}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_haproxy_layer#buffer_duration OpsworksHaproxyLayer#buffer_duration}.
	BufferDuration *float64 `field:"optional" json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_haproxy_layer#datetime_format OpsworksHaproxyLayer#datetime_format}.
	DatetimeFormat *string `field:"optional" json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_haproxy_layer#encoding OpsworksHaproxyLayer#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_haproxy_layer#file_fingerprint_lines OpsworksHaproxyLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `field:"optional" json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_haproxy_layer#initial_position OpsworksHaproxyLayer#initial_position}.
	InitialPosition *string `field:"optional" json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_haproxy_layer#multiline_start_pattern OpsworksHaproxyLayer#multiline_start_pattern}.
	MultilineStartPattern *string `field:"optional" json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.30.0/docs/resources/opsworks_haproxy_layer#time_zone OpsworksHaproxyLayer#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworkscustomlayer


type OpsworksCustomLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/opsworks_custom_layer#file OpsworksCustomLayer#file}.
	File *string `field:"required" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/opsworks_custom_layer#log_group_name OpsworksCustomLayer#log_group_name}.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/opsworks_custom_layer#batch_count OpsworksCustomLayer#batch_count}.
	BatchCount *float64 `field:"optional" json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/opsworks_custom_layer#batch_size OpsworksCustomLayer#batch_size}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/opsworks_custom_layer#buffer_duration OpsworksCustomLayer#buffer_duration}.
	BufferDuration *float64 `field:"optional" json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/opsworks_custom_layer#datetime_format OpsworksCustomLayer#datetime_format}.
	DatetimeFormat *string `field:"optional" json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/opsworks_custom_layer#encoding OpsworksCustomLayer#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/opsworks_custom_layer#file_fingerprint_lines OpsworksCustomLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `field:"optional" json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/opsworks_custom_layer#initial_position OpsworksCustomLayer#initial_position}.
	InitialPosition *string `field:"optional" json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/opsworks_custom_layer#multiline_start_pattern OpsworksCustomLayer#multiline_start_pattern}.
	MultilineStartPattern *string `field:"optional" json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/opsworks_custom_layer#time_zone OpsworksCustomLayer#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}


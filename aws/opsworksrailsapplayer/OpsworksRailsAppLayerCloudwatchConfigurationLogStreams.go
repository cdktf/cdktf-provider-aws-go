// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksrailsapplayer


type OpsworksRailsAppLayerCloudwatchConfigurationLogStreams struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/opsworks_rails_app_layer#file OpsworksRailsAppLayer#file}.
	File *string `field:"required" json:"file" yaml:"file"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/opsworks_rails_app_layer#log_group_name OpsworksRailsAppLayer#log_group_name}.
	LogGroupName *string `field:"required" json:"logGroupName" yaml:"logGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/opsworks_rails_app_layer#batch_count OpsworksRailsAppLayer#batch_count}.
	BatchCount *float64 `field:"optional" json:"batchCount" yaml:"batchCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/opsworks_rails_app_layer#batch_size OpsworksRailsAppLayer#batch_size}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/opsworks_rails_app_layer#buffer_duration OpsworksRailsAppLayer#buffer_duration}.
	BufferDuration *float64 `field:"optional" json:"bufferDuration" yaml:"bufferDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/opsworks_rails_app_layer#datetime_format OpsworksRailsAppLayer#datetime_format}.
	DatetimeFormat *string `field:"optional" json:"datetimeFormat" yaml:"datetimeFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/opsworks_rails_app_layer#encoding OpsworksRailsAppLayer#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/opsworks_rails_app_layer#file_fingerprint_lines OpsworksRailsAppLayer#file_fingerprint_lines}.
	FileFingerprintLines *string `field:"optional" json:"fileFingerprintLines" yaml:"fileFingerprintLines"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/opsworks_rails_app_layer#initial_position OpsworksRailsAppLayer#initial_position}.
	InitialPosition *string `field:"optional" json:"initialPosition" yaml:"initialPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/opsworks_rails_app_layer#multiline_start_pattern OpsworksRailsAppLayer#multiline_start_pattern}.
	MultilineStartPattern *string `field:"optional" json:"multilineStartPattern" yaml:"multilineStartPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/opsworks_rails_app_layer#time_zone OpsworksRailsAppLayer#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}


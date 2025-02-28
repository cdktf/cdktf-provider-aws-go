// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksnodejsapplayer


type OpsworksNodejsAppLayerLoadBasedAutoScalingUpscaling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/opsworks_nodejs_app_layer#alarms OpsworksNodejsAppLayer#alarms}.
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/opsworks_nodejs_app_layer#cpu_threshold OpsworksNodejsAppLayer#cpu_threshold}.
	CpuThreshold *float64 `field:"optional" json:"cpuThreshold" yaml:"cpuThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/opsworks_nodejs_app_layer#ignore_metrics_time OpsworksNodejsAppLayer#ignore_metrics_time}.
	IgnoreMetricsTime *float64 `field:"optional" json:"ignoreMetricsTime" yaml:"ignoreMetricsTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/opsworks_nodejs_app_layer#instance_count OpsworksNodejsAppLayer#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/opsworks_nodejs_app_layer#load_threshold OpsworksNodejsAppLayer#load_threshold}.
	LoadThreshold *float64 `field:"optional" json:"loadThreshold" yaml:"loadThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/opsworks_nodejs_app_layer#memory_threshold OpsworksNodejsAppLayer#memory_threshold}.
	MemoryThreshold *float64 `field:"optional" json:"memoryThreshold" yaml:"memoryThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/opsworks_nodejs_app_layer#thresholds_wait_time OpsworksNodejsAppLayer#thresholds_wait_time}.
	ThresholdsWaitTime *float64 `field:"optional" json:"thresholdsWaitTime" yaml:"thresholdsWaitTime"`
}


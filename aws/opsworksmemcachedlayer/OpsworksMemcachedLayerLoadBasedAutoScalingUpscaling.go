// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksmemcachedlayer


type OpsworksMemcachedLayerLoadBasedAutoScalingUpscaling struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/opsworks_memcached_layer#alarms OpsworksMemcachedLayer#alarms}.
	Alarms *[]*string `field:"optional" json:"alarms" yaml:"alarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/opsworks_memcached_layer#cpu_threshold OpsworksMemcachedLayer#cpu_threshold}.
	CpuThreshold *float64 `field:"optional" json:"cpuThreshold" yaml:"cpuThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/opsworks_memcached_layer#ignore_metrics_time OpsworksMemcachedLayer#ignore_metrics_time}.
	IgnoreMetricsTime *float64 `field:"optional" json:"ignoreMetricsTime" yaml:"ignoreMetricsTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/opsworks_memcached_layer#instance_count OpsworksMemcachedLayer#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/opsworks_memcached_layer#load_threshold OpsworksMemcachedLayer#load_threshold}.
	LoadThreshold *float64 `field:"optional" json:"loadThreshold" yaml:"loadThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/opsworks_memcached_layer#memory_threshold OpsworksMemcachedLayer#memory_threshold}.
	MemoryThreshold *float64 `field:"optional" json:"memoryThreshold" yaml:"memoryThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/opsworks_memcached_layer#thresholds_wait_time OpsworksMemcachedLayer#thresholds_wait_time}.
	ThresholdsWaitTime *float64 `field:"optional" json:"thresholdsWaitTime" yaml:"thresholdsWaitTime"`
}


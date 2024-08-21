// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package internetmonitormonitor


type InternetmonitorMonitorInternetMeasurementsLogDeliveryS3Config struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/internetmonitor_monitor#bucket_name InternetmonitorMonitor#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/internetmonitor_monitor#bucket_prefix InternetmonitorMonitor#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/internetmonitor_monitor#log_delivery_status InternetmonitorMonitor#log_delivery_status}.
	LogDeliveryStatus *string `field:"optional" json:"logDeliveryStatus" yaml:"logDeliveryStatus"`
}


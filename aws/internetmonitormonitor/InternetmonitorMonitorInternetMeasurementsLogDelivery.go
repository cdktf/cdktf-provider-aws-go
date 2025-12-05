// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package internetmonitormonitor


type InternetmonitorMonitorInternetMeasurementsLogDelivery struct {
	// s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/internetmonitor_monitor#s3_config InternetmonitorMonitor#s3_config}
	S3Config *InternetmonitorMonitorInternetMeasurementsLogDeliveryS3Config `field:"optional" json:"s3Config" yaml:"s3Config"`
}


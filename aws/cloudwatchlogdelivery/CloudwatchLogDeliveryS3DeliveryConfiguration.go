// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudwatchlogdelivery


type CloudwatchLogDeliveryS3DeliveryConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/cloudwatch_log_delivery#enable_hive_compatible_path CloudwatchLogDelivery#enable_hive_compatible_path}.
	EnableHiveCompatiblePath interface{} `field:"optional" json:"enableHiveCompatiblePath" yaml:"enableHiveCompatiblePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/cloudwatch_log_delivery#suffix_path CloudwatchLogDelivery#suffix_path}.
	SuffixPath *string `field:"optional" json:"suffixPath" yaml:"suffixPath"`
}


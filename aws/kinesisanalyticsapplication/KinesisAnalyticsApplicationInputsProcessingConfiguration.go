// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsapplication


type KinesisAnalyticsApplicationInputsProcessingConfiguration struct {
	// lambda block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/kinesis_analytics_application#lambda KinesisAnalyticsApplication#lambda}
	Lambda *KinesisAnalyticsApplicationInputsProcessingConfigurationLambda `field:"required" json:"lambda" yaml:"lambda"`
}


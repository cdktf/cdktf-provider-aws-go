// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputProcessingConfiguration struct {
	// input_lambda_processor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/kinesisanalyticsv2_application#input_lambda_processor Kinesisanalyticsv2Application#input_lambda_processor}
	InputLambdaProcessor *Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputInputProcessingConfigurationInputLambdaProcessor `field:"required" json:"inputLambdaProcessor" yaml:"inputLambdaProcessor"`
}


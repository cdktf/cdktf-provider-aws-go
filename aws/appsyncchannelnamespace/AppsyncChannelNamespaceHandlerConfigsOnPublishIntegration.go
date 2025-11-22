// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncchannelnamespace


type AppsyncChannelNamespaceHandlerConfigsOnPublishIntegration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appsync_channel_namespace#data_source_name AppsyncChannelNamespace#data_source_name}.
	DataSourceName *string `field:"required" json:"dataSourceName" yaml:"dataSourceName"`
	// lambda_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/appsync_channel_namespace#lambda_config AppsyncChannelNamespace#lambda_config}
	LambdaConfig interface{} `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
}


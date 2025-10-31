// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncchannelnamespace


type AppsyncChannelNamespaceHandlerConfigsOnSubscribe struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/appsync_channel_namespace#behavior AppsyncChannelNamespace#behavior}.
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// integration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/appsync_channel_namespace#integration AppsyncChannelNamespace#integration}
	Integration interface{} `field:"optional" json:"integration" yaml:"integration"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncchannelnamespace


type AppsyncChannelNamespaceHandlerConfigs struct {
	// on_publish block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/appsync_channel_namespace#on_publish AppsyncChannelNamespace#on_publish}
	OnPublish interface{} `field:"optional" json:"onPublish" yaml:"onPublish"`
	// on_subscribe block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/appsync_channel_namespace#on_subscribe AppsyncChannelNamespace#on_subscribe}
	OnSubscribe interface{} `field:"optional" json:"onSubscribe" yaml:"onSubscribe"`
}


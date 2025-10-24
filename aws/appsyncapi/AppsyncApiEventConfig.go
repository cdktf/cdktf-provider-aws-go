// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appsyncapi


type AppsyncApiEventConfig struct {
	// auth_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/appsync_api#auth_provider AppsyncApi#auth_provider}
	AuthProvider interface{} `field:"optional" json:"authProvider" yaml:"authProvider"`
	// connection_auth_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/appsync_api#connection_auth_mode AppsyncApi#connection_auth_mode}
	ConnectionAuthMode interface{} `field:"optional" json:"connectionAuthMode" yaml:"connectionAuthMode"`
	// default_publish_auth_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/appsync_api#default_publish_auth_mode AppsyncApi#default_publish_auth_mode}
	DefaultPublishAuthMode interface{} `field:"optional" json:"defaultPublishAuthMode" yaml:"defaultPublishAuthMode"`
	// default_subscribe_auth_mode block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/appsync_api#default_subscribe_auth_mode AppsyncApi#default_subscribe_auth_mode}
	DefaultSubscribeAuthMode interface{} `field:"optional" json:"defaultSubscribeAuthMode" yaml:"defaultSubscribeAuthMode"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/appsync_api#log_config AppsyncApi#log_config}
	LogConfig interface{} `field:"optional" json:"logConfig" yaml:"logConfig"`
}


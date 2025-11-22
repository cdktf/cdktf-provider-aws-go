// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectquickconnect


type ConnectQuickConnectQuickConnectConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/connect_quick_connect#quick_connect_type ConnectQuickConnect#quick_connect_type}.
	QuickConnectType *string `field:"required" json:"quickConnectType" yaml:"quickConnectType"`
	// phone_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/connect_quick_connect#phone_config ConnectQuickConnect#phone_config}
	PhoneConfig interface{} `field:"optional" json:"phoneConfig" yaml:"phoneConfig"`
	// queue_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/connect_quick_connect#queue_config ConnectQuickConnect#queue_config}
	QueueConfig interface{} `field:"optional" json:"queueConfig" yaml:"queueConfig"`
	// user_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/connect_quick_connect#user_config ConnectQuickConnect#user_config}
	UserConfig interface{} `field:"optional" json:"userConfig" yaml:"userConfig"`
}


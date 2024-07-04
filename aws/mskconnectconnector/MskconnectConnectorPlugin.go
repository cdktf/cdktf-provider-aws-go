// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskconnectconnector


type MskconnectConnectorPlugin struct {
	// custom_plugin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.57.0/docs/resources/mskconnect_connector#custom_plugin MskconnectConnector#custom_plugin}
	CustomPlugin *MskconnectConnectorPluginCustomPlugin `field:"required" json:"customPlugin" yaml:"customPlugin"`
}


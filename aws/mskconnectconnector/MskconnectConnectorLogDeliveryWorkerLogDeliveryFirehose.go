// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskconnectconnector


type MskconnectConnectorLogDeliveryWorkerLogDeliveryFirehose struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/mskconnect_connector#enabled MskconnectConnector#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/mskconnect_connector#delivery_stream MskconnectConnector#delivery_stream}.
	DeliveryStream *string `field:"optional" json:"deliveryStream" yaml:"deliveryStream"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2configurationset


type Sesv2ConfigurationSetDeliveryOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/sesv2_configuration_set#max_delivery_seconds Sesv2ConfigurationSet#max_delivery_seconds}.
	MaxDeliverySeconds *float64 `field:"optional" json:"maxDeliverySeconds" yaml:"maxDeliverySeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/sesv2_configuration_set#sending_pool_name Sesv2ConfigurationSet#sending_pool_name}.
	SendingPoolName *string `field:"optional" json:"sendingPoolName" yaml:"sendingPoolName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/sesv2_configuration_set#tls_policy Sesv2ConfigurationSet#tls_policy}.
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
}


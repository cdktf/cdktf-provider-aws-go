// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2configurationset


type Sesv2ConfigurationSetTrackingOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/sesv2_configuration_set#custom_redirect_domain Sesv2ConfigurationSet#custom_redirect_domain}.
	CustomRedirectDomain *string `field:"required" json:"customRedirectDomain" yaml:"customRedirectDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/sesv2_configuration_set#https_policy Sesv2ConfigurationSet#https_policy}.
	HttpsPolicy *string `field:"optional" json:"httpsPolicy" yaml:"httpsPolicy"`
}


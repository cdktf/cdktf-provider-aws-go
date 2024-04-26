// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotdomainconfiguration


type IotDomainConfigurationAuthorizerConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/iot_domain_configuration#allow_authorizer_override IotDomainConfiguration#allow_authorizer_override}.
	AllowAuthorizerOverride interface{} `field:"optional" json:"allowAuthorizerOverride" yaml:"allowAuthorizerOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/iot_domain_configuration#default_authorizer_name IotDomainConfiguration#default_authorizer_name}.
	DefaultAuthorizerName *string `field:"optional" json:"defaultAuthorizerName" yaml:"defaultAuthorizerName"`
}


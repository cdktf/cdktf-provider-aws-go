// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotdomainconfiguration


type IotDomainConfigurationTlsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/iot_domain_configuration#security_policy IotDomainConfiguration#security_policy}.
	SecurityPolicy *string `field:"optional" json:"securityPolicy" yaml:"securityPolicy"`
}


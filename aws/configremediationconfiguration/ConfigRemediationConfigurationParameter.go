// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configremediationconfiguration


type ConfigRemediationConfigurationParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_remediation_configuration#name ConfigRemediationConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_remediation_configuration#resource_value ConfigRemediationConfiguration#resource_value}.
	ResourceValue *string `field:"optional" json:"resourceValue" yaml:"resourceValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_remediation_configuration#static_value ConfigRemediationConfiguration#static_value}.
	StaticValue *string `field:"optional" json:"staticValue" yaml:"staticValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/config_remediation_configuration#static_values ConfigRemediationConfiguration#static_values}.
	StaticValues *[]*string `field:"optional" json:"staticValues" yaml:"staticValues"`
}


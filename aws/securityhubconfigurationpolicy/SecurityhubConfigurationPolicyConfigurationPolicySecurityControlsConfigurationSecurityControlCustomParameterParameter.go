// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubconfigurationpolicy


type SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securityhub_configuration_policy#name SecurityhubConfigurationPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securityhub_configuration_policy#value_type SecurityhubConfigurationPolicy#value_type}.
	ValueType *string `field:"required" json:"valueType" yaml:"valueType"`
	// bool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securityhub_configuration_policy#bool SecurityhubConfigurationPolicy#bool}
	Bool *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterBool `field:"optional" json:"bool" yaml:"bool"`
	// double block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securityhub_configuration_policy#double SecurityhubConfigurationPolicy#double}
	Double *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterDouble `field:"optional" json:"double" yaml:"double"`
	// enum block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securityhub_configuration_policy#enum SecurityhubConfigurationPolicy#enum}
	Enum *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterEnum `field:"optional" json:"enum" yaml:"enum"`
	// enum_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securityhub_configuration_policy#enum_list SecurityhubConfigurationPolicy#enum_list}
	EnumList *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterEnumListStruct `field:"optional" json:"enumList" yaml:"enumList"`
	// int block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securityhub_configuration_policy#int SecurityhubConfigurationPolicy#int}
	Int *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterInt `field:"optional" json:"int" yaml:"int"`
	// int_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securityhub_configuration_policy#int_list SecurityhubConfigurationPolicy#int_list}
	IntList *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterIntListStruct `field:"optional" json:"intList" yaml:"intList"`
	// string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securityhub_configuration_policy#string SecurityhubConfigurationPolicy#string}
	String *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterString `field:"optional" json:"string" yaml:"string"`
	// string_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/securityhub_configuration_policy#string_list SecurityhubConfigurationPolicy#string_list}
	StringList *SecurityhubConfigurationPolicyConfigurationPolicySecurityControlsConfigurationSecurityControlCustomParameterParameterStringListStruct `field:"optional" json:"stringList" yaml:"stringList"`
}


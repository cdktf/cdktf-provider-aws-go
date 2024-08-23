// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotcacertificate


type IotCaCertificateRegistrationConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/iot_ca_certificate#role_arn IotCaCertificate#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/iot_ca_certificate#template_body IotCaCertificate#template_body}.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/iot_ca_certificate#template_name IotCaCertificate#template_name}.
	TemplateName *string `field:"optional" json:"templateName" yaml:"templateName"`
}


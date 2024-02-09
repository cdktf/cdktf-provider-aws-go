// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanagercontrol


type AuditmanagerControlControlMappingSources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/auditmanager_control#source_name AuditmanagerControl#source_name}.
	SourceName *string `field:"required" json:"sourceName" yaml:"sourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/auditmanager_control#source_set_up_option AuditmanagerControl#source_set_up_option}.
	SourceSetUpOption *string `field:"required" json:"sourceSetUpOption" yaml:"sourceSetUpOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/auditmanager_control#source_type AuditmanagerControl#source_type}.
	SourceType *string `field:"required" json:"sourceType" yaml:"sourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/auditmanager_control#source_description AuditmanagerControl#source_description}.
	SourceDescription *string `field:"optional" json:"sourceDescription" yaml:"sourceDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/auditmanager_control#source_frequency AuditmanagerControl#source_frequency}.
	SourceFrequency *string `field:"optional" json:"sourceFrequency" yaml:"sourceFrequency"`
	// source_keyword block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/auditmanager_control#source_keyword AuditmanagerControl#source_keyword}
	SourceKeyword interface{} `field:"optional" json:"sourceKeyword" yaml:"sourceKeyword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.36.0/docs/resources/auditmanager_control#troubleshooting_text AuditmanagerControl#troubleshooting_text}.
	TroubleshootingText *string `field:"optional" json:"troubleshootingText" yaml:"troubleshootingText"`
}


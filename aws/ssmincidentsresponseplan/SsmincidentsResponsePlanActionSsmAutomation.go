// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmincidentsresponseplan


type SsmincidentsResponsePlanActionSsmAutomation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/ssmincidents_response_plan#document_name SsmincidentsResponsePlan#document_name}.
	DocumentName *string `field:"required" json:"documentName" yaml:"documentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/ssmincidents_response_plan#role_arn SsmincidentsResponsePlan#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/ssmincidents_response_plan#document_version SsmincidentsResponsePlan#document_version}.
	DocumentVersion *string `field:"optional" json:"documentVersion" yaml:"documentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/ssmincidents_response_plan#dynamic_parameters SsmincidentsResponsePlan#dynamic_parameters}.
	DynamicParameters *map[string]*string `field:"optional" json:"dynamicParameters" yaml:"dynamicParameters"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/ssmincidents_response_plan#parameter SsmincidentsResponsePlan#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.24.0/docs/resources/ssmincidents_response_plan#target_account SsmincidentsResponsePlan#target_account}.
	TargetAccount *string `field:"optional" json:"targetAccount" yaml:"targetAccount"`
}


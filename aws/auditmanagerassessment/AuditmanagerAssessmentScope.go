// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessment


type AuditmanagerAssessmentScope struct {
	// aws_accounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/auditmanager_assessment#aws_accounts AuditmanagerAssessment#aws_accounts}
	AwsAccounts interface{} `field:"optional" json:"awsAccounts" yaml:"awsAccounts"`
	// aws_services block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/auditmanager_assessment#aws_services AuditmanagerAssessment#aws_services}
	AwsServices interface{} `field:"optional" json:"awsServices" yaml:"awsServices"`
}


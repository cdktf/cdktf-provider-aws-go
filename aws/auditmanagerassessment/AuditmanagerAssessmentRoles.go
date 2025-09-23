// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessment


type AuditmanagerAssessmentRoles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/auditmanager_assessment#role_arn AuditmanagerAssessment#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/auditmanager_assessment#role_type AuditmanagerAssessment#role_type}.
	RoleType *string `field:"required" json:"roleType" yaml:"roleType"`
}


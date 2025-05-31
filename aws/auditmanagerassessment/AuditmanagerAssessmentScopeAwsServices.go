// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessment


type AuditmanagerAssessmentScopeAwsServices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/auditmanager_assessment#service_name AuditmanagerAssessment#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guarddutyorganizationconfiguration


type GuarddutyOrganizationConfigurationDatasourcesKubernetes struct {
	// audit_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.64.0/docs/resources/guardduty_organization_configuration#audit_logs GuarddutyOrganizationConfiguration#audit_logs}
	AuditLogs *GuarddutyOrganizationConfigurationDatasourcesKubernetesAuditLogs `field:"required" json:"auditLogs" yaml:"auditLogs"`
}


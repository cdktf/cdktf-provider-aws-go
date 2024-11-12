// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package guarddutydetector


type GuarddutyDetectorDatasourcesKubernetes struct {
	// audit_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/guardduty_detector#audit_logs GuarddutyDetector#audit_logs}
	AuditLogs *GuarddutyDetectorDatasourcesKubernetesAuditLogs `field:"required" json:"auditLogs" yaml:"auditLogs"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupplan


type BackupPlanRuleScanAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/backup_plan#malware_scanner BackupPlan#malware_scanner}.
	MalwareScanner *string `field:"required" json:"malwareScanner" yaml:"malwareScanner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/backup_plan#scan_mode BackupPlan#scan_mode}.
	ScanMode *string `field:"required" json:"scanMode" yaml:"scanMode"`
}


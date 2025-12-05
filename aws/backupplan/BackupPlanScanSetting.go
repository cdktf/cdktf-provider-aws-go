// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupplan


type BackupPlanScanSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/backup_plan#malware_scanner BackupPlan#malware_scanner}.
	MalwareScanner *string `field:"required" json:"malwareScanner" yaml:"malwareScanner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/backup_plan#resource_types BackupPlan#resource_types}.
	ResourceTypes *[]*string `field:"required" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/backup_plan#scanner_role_arn BackupPlan#scanner_role_arn}.
	ScannerRoleArn *string `field:"required" json:"scannerRoleArn" yaml:"scannerRoleArn"`
}


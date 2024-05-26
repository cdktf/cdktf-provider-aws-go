// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupreportplan


type BackupReportPlanReportSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/backup_report_plan#report_template BackupReportPlan#report_template}.
	ReportTemplate *string `field:"required" json:"reportTemplate" yaml:"reportTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/backup_report_plan#accounts BackupReportPlan#accounts}.
	Accounts *[]*string `field:"optional" json:"accounts" yaml:"accounts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/backup_report_plan#framework_arns BackupReportPlan#framework_arns}.
	FrameworkArns *[]*string `field:"optional" json:"frameworkArns" yaml:"frameworkArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/backup_report_plan#number_of_frameworks BackupReportPlan#number_of_frameworks}.
	NumberOfFrameworks *float64 `field:"optional" json:"numberOfFrameworks" yaml:"numberOfFrameworks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/backup_report_plan#organization_units BackupReportPlan#organization_units}.
	OrganizationUnits *[]*string `field:"optional" json:"organizationUnits" yaml:"organizationUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.1/docs/resources/backup_report_plan#regions BackupReportPlan#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}


package backupreportplan


type BackupReportPlanReportSetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/backup_report_plan#report_template BackupReportPlan#report_template}.
	ReportTemplate *string `field:"required" json:"reportTemplate" yaml:"reportTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/backup_report_plan#framework_arns BackupReportPlan#framework_arns}.
	FrameworkArns *[]*string `field:"optional" json:"frameworkArns" yaml:"frameworkArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.64.0/docs/resources/backup_report_plan#number_of_frameworks BackupReportPlan#number_of_frameworks}.
	NumberOfFrameworks *float64 `field:"optional" json:"numberOfFrameworks" yaml:"numberOfFrameworks"`
}


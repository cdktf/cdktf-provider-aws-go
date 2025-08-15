// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backupreportplan


type BackupReportPlanReportDeliveryChannel struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/backup_report_plan#s3_bucket_name BackupReportPlan#s3_bucket_name}.
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/backup_report_plan#formats BackupReportPlan#formats}.
	Formats *[]*string `field:"optional" json:"formats" yaml:"formats"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/backup_report_plan#s3_key_prefix BackupReportPlan#s3_key_prefix}.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}


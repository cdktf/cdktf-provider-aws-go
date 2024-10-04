// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbinstance


type DbInstanceRestoreToPointInTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/db_instance#restore_time DbInstance#restore_time}.
	RestoreTime *string `field:"optional" json:"restoreTime" yaml:"restoreTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/db_instance#source_db_instance_automated_backups_arn DbInstance#source_db_instance_automated_backups_arn}.
	SourceDbInstanceAutomatedBackupsArn *string `field:"optional" json:"sourceDbInstanceAutomatedBackupsArn" yaml:"sourceDbInstanceAutomatedBackupsArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/db_instance#source_db_instance_identifier DbInstance#source_db_instance_identifier}.
	SourceDbInstanceIdentifier *string `field:"optional" json:"sourceDbInstanceIdentifier" yaml:"sourceDbInstanceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/db_instance#source_dbi_resource_id DbInstance#source_dbi_resource_id}.
	SourceDbiResourceId *string `field:"optional" json:"sourceDbiResourceId" yaml:"sourceDbiResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/db_instance#use_latest_restorable_time DbInstance#use_latest_restorable_time}.
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
}


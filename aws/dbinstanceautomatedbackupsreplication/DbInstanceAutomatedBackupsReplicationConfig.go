// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbinstanceautomatedbackupsreplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DbInstanceAutomatedBackupsReplicationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/db_instance_automated_backups_replication#source_db_instance_arn DbInstanceAutomatedBackupsReplication#source_db_instance_arn}.
	SourceDbInstanceArn *string `field:"required" json:"sourceDbInstanceArn" yaml:"sourceDbInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/db_instance_automated_backups_replication#id DbInstanceAutomatedBackupsReplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/db_instance_automated_backups_replication#kms_key_id DbInstanceAutomatedBackupsReplication#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/db_instance_automated_backups_replication#pre_signed_url DbInstanceAutomatedBackupsReplication#pre_signed_url}.
	PreSignedUrl *string `field:"optional" json:"preSignedUrl" yaml:"preSignedUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/db_instance_automated_backups_replication#retention_period DbInstanceAutomatedBackupsReplication#retention_period}.
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/db_instance_automated_backups_replication#timeouts DbInstanceAutomatedBackupsReplication#timeouts}
	Timeouts *DbInstanceAutomatedBackupsReplicationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


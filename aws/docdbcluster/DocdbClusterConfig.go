// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package docdbcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DocdbClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#apply_immediately DocdbCluster#apply_immediately}.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#availability_zones DocdbCluster#availability_zones}.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#backup_retention_period DocdbCluster#backup_retention_period}.
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#cluster_identifier DocdbCluster#cluster_identifier}.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#cluster_identifier_prefix DocdbCluster#cluster_identifier_prefix}.
	ClusterIdentifierPrefix *string `field:"optional" json:"clusterIdentifierPrefix" yaml:"clusterIdentifierPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#cluster_members DocdbCluster#cluster_members}.
	ClusterMembers *[]*string `field:"optional" json:"clusterMembers" yaml:"clusterMembers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#db_cluster_parameter_group_name DocdbCluster#db_cluster_parameter_group_name}.
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#db_subnet_group_name DocdbCluster#db_subnet_group_name}.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#deletion_protection DocdbCluster#deletion_protection}.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#enabled_cloudwatch_logs_exports DocdbCluster#enabled_cloudwatch_logs_exports}.
	EnabledCloudwatchLogsExports *[]*string `field:"optional" json:"enabledCloudwatchLogsExports" yaml:"enabledCloudwatchLogsExports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#engine DocdbCluster#engine}.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#engine_version DocdbCluster#engine_version}.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#final_snapshot_identifier DocdbCluster#final_snapshot_identifier}.
	FinalSnapshotIdentifier *string `field:"optional" json:"finalSnapshotIdentifier" yaml:"finalSnapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#global_cluster_identifier DocdbCluster#global_cluster_identifier}.
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#id DocdbCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#kms_key_id DocdbCluster#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#master_password DocdbCluster#master_password}.
	MasterPassword *string `field:"optional" json:"masterPassword" yaml:"masterPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#master_username DocdbCluster#master_username}.
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#port DocdbCluster#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#preferred_backup_window DocdbCluster#preferred_backup_window}.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#preferred_maintenance_window DocdbCluster#preferred_maintenance_window}.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#skip_final_snapshot DocdbCluster#skip_final_snapshot}.
	SkipFinalSnapshot interface{} `field:"optional" json:"skipFinalSnapshot" yaml:"skipFinalSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#snapshot_identifier DocdbCluster#snapshot_identifier}.
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#storage_encrypted DocdbCluster#storage_encrypted}.
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#tags DocdbCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#tags_all DocdbCluster#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#timeouts DocdbCluster#timeouts}
	Timeouts *DocdbClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.16.1/docs/resources/docdb_cluster#vpc_security_group_ids DocdbCluster#vpc_security_group_ids}.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}


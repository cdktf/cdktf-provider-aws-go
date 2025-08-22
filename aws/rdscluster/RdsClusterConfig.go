// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdscluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#engine RdsCluster#engine}.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#allocated_storage RdsCluster#allocated_storage}.
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#allow_major_version_upgrade RdsCluster#allow_major_version_upgrade}.
	AllowMajorVersionUpgrade interface{} `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#apply_immediately RdsCluster#apply_immediately}.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#availability_zones RdsCluster#availability_zones}.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#backtrack_window RdsCluster#backtrack_window}.
	BacktrackWindow *float64 `field:"optional" json:"backtrackWindow" yaml:"backtrackWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#backup_retention_period RdsCluster#backup_retention_period}.
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#ca_certificate_identifier RdsCluster#ca_certificate_identifier}.
	CaCertificateIdentifier *string `field:"optional" json:"caCertificateIdentifier" yaml:"caCertificateIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#cluster_identifier RdsCluster#cluster_identifier}.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#cluster_identifier_prefix RdsCluster#cluster_identifier_prefix}.
	ClusterIdentifierPrefix *string `field:"optional" json:"clusterIdentifierPrefix" yaml:"clusterIdentifierPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#cluster_members RdsCluster#cluster_members}.
	ClusterMembers *[]*string `field:"optional" json:"clusterMembers" yaml:"clusterMembers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#cluster_scalability_type RdsCluster#cluster_scalability_type}.
	ClusterScalabilityType *string `field:"optional" json:"clusterScalabilityType" yaml:"clusterScalabilityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#copy_tags_to_snapshot RdsCluster#copy_tags_to_snapshot}.
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#database_insights_mode RdsCluster#database_insights_mode}.
	DatabaseInsightsMode *string `field:"optional" json:"databaseInsightsMode" yaml:"databaseInsightsMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#database_name RdsCluster#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#db_cluster_instance_class RdsCluster#db_cluster_instance_class}.
	DbClusterInstanceClass *string `field:"optional" json:"dbClusterInstanceClass" yaml:"dbClusterInstanceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#db_cluster_parameter_group_name RdsCluster#db_cluster_parameter_group_name}.
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#db_instance_parameter_group_name RdsCluster#db_instance_parameter_group_name}.
	DbInstanceParameterGroupName *string `field:"optional" json:"dbInstanceParameterGroupName" yaml:"dbInstanceParameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#db_subnet_group_name RdsCluster#db_subnet_group_name}.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#db_system_id RdsCluster#db_system_id}.
	DbSystemId *string `field:"optional" json:"dbSystemId" yaml:"dbSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#delete_automated_backups RdsCluster#delete_automated_backups}.
	DeleteAutomatedBackups interface{} `field:"optional" json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#deletion_protection RdsCluster#deletion_protection}.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#domain RdsCluster#domain}.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#domain_iam_role_name RdsCluster#domain_iam_role_name}.
	DomainIamRoleName *string `field:"optional" json:"domainIamRoleName" yaml:"domainIamRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#enabled_cloudwatch_logs_exports RdsCluster#enabled_cloudwatch_logs_exports}.
	EnabledCloudwatchLogsExports *[]*string `field:"optional" json:"enabledCloudwatchLogsExports" yaml:"enabledCloudwatchLogsExports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#enable_global_write_forwarding RdsCluster#enable_global_write_forwarding}.
	EnableGlobalWriteForwarding interface{} `field:"optional" json:"enableGlobalWriteForwarding" yaml:"enableGlobalWriteForwarding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#enable_http_endpoint RdsCluster#enable_http_endpoint}.
	EnableHttpEndpoint interface{} `field:"optional" json:"enableHttpEndpoint" yaml:"enableHttpEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#enable_local_write_forwarding RdsCluster#enable_local_write_forwarding}.
	EnableLocalWriteForwarding interface{} `field:"optional" json:"enableLocalWriteForwarding" yaml:"enableLocalWriteForwarding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#engine_lifecycle_support RdsCluster#engine_lifecycle_support}.
	EngineLifecycleSupport *string `field:"optional" json:"engineLifecycleSupport" yaml:"engineLifecycleSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#engine_mode RdsCluster#engine_mode}.
	EngineMode *string `field:"optional" json:"engineMode" yaml:"engineMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#engine_version RdsCluster#engine_version}.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#final_snapshot_identifier RdsCluster#final_snapshot_identifier}.
	FinalSnapshotIdentifier *string `field:"optional" json:"finalSnapshotIdentifier" yaml:"finalSnapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#global_cluster_identifier RdsCluster#global_cluster_identifier}.
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#iam_database_authentication_enabled RdsCluster#iam_database_authentication_enabled}.
	IamDatabaseAuthenticationEnabled interface{} `field:"optional" json:"iamDatabaseAuthenticationEnabled" yaml:"iamDatabaseAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#iam_roles RdsCluster#iam_roles}.
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#id RdsCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#iops RdsCluster#iops}.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#kms_key_id RdsCluster#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#manage_master_user_password RdsCluster#manage_master_user_password}.
	ManageMasterUserPassword interface{} `field:"optional" json:"manageMasterUserPassword" yaml:"manageMasterUserPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#master_password RdsCluster#master_password}.
	MasterPassword *string `field:"optional" json:"masterPassword" yaml:"masterPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#master_password_wo RdsCluster#master_password_wo}.
	MasterPasswordWo *string `field:"optional" json:"masterPasswordWo" yaml:"masterPasswordWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#master_password_wo_version RdsCluster#master_password_wo_version}.
	MasterPasswordWoVersion *float64 `field:"optional" json:"masterPasswordWoVersion" yaml:"masterPasswordWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#master_username RdsCluster#master_username}.
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#master_user_secret_kms_key_id RdsCluster#master_user_secret_kms_key_id}.
	MasterUserSecretKmsKeyId *string `field:"optional" json:"masterUserSecretKmsKeyId" yaml:"masterUserSecretKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#monitoring_interval RdsCluster#monitoring_interval}.
	MonitoringInterval *float64 `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#monitoring_role_arn RdsCluster#monitoring_role_arn}.
	MonitoringRoleArn *string `field:"optional" json:"monitoringRoleArn" yaml:"monitoringRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#network_type RdsCluster#network_type}.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#performance_insights_enabled RdsCluster#performance_insights_enabled}.
	PerformanceInsightsEnabled interface{} `field:"optional" json:"performanceInsightsEnabled" yaml:"performanceInsightsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#performance_insights_kms_key_id RdsCluster#performance_insights_kms_key_id}.
	PerformanceInsightsKmsKeyId *string `field:"optional" json:"performanceInsightsKmsKeyId" yaml:"performanceInsightsKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#performance_insights_retention_period RdsCluster#performance_insights_retention_period}.
	PerformanceInsightsRetentionPeriod *float64 `field:"optional" json:"performanceInsightsRetentionPeriod" yaml:"performanceInsightsRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#port RdsCluster#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#preferred_backup_window RdsCluster#preferred_backup_window}.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#preferred_maintenance_window RdsCluster#preferred_maintenance_window}.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#region RdsCluster#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#replication_source_identifier RdsCluster#replication_source_identifier}.
	ReplicationSourceIdentifier *string `field:"optional" json:"replicationSourceIdentifier" yaml:"replicationSourceIdentifier"`
	// restore_to_point_in_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#restore_to_point_in_time RdsCluster#restore_to_point_in_time}
	RestoreToPointInTime *RdsClusterRestoreToPointInTime `field:"optional" json:"restoreToPointInTime" yaml:"restoreToPointInTime"`
	// s3_import block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#s3_import RdsCluster#s3_import}
	S3Import *RdsClusterS3Import `field:"optional" json:"s3Import" yaml:"s3Import"`
	// scaling_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#scaling_configuration RdsCluster#scaling_configuration}
	ScalingConfiguration *RdsClusterScalingConfiguration `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// serverlessv2_scaling_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#serverlessv2_scaling_configuration RdsCluster#serverlessv2_scaling_configuration}
	Serverlessv2ScalingConfiguration *RdsClusterServerlessv2ScalingConfiguration `field:"optional" json:"serverlessv2ScalingConfiguration" yaml:"serverlessv2ScalingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#skip_final_snapshot RdsCluster#skip_final_snapshot}.
	SkipFinalSnapshot interface{} `field:"optional" json:"skipFinalSnapshot" yaml:"skipFinalSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#snapshot_identifier RdsCluster#snapshot_identifier}.
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#source_region RdsCluster#source_region}.
	SourceRegion *string `field:"optional" json:"sourceRegion" yaml:"sourceRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#storage_encrypted RdsCluster#storage_encrypted}.
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#storage_type RdsCluster#storage_type}.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#tags RdsCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#tags_all RdsCluster#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#timeouts RdsCluster#timeouts}
	Timeouts *RdsClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/rds_cluster#vpc_security_group_ids RdsCluster#vpc_security_group_ids}.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}


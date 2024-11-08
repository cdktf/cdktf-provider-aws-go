// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdscluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/rdscluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/rds_cluster aws_rds_cluster}.
type RdsCluster interface {
	cdktf.TerraformResource
	AllocatedStorage() *float64
	SetAllocatedStorage(val *float64)
	AllocatedStorageInput() *float64
	AllowMajorVersionUpgrade() interface{}
	SetAllowMajorVersionUpgrade(val interface{})
	AllowMajorVersionUpgradeInput() interface{}
	ApplyImmediately() interface{}
	SetApplyImmediately(val interface{})
	ApplyImmediatelyInput() interface{}
	Arn() *string
	AvailabilityZones() *[]*string
	SetAvailabilityZones(val *[]*string)
	AvailabilityZonesInput() *[]*string
	BacktrackWindow() *float64
	SetBacktrackWindow(val *float64)
	BacktrackWindowInput() *float64
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	BackupRetentionPeriodInput() *float64
	CaCertificateIdentifier() *string
	SetCaCertificateIdentifier(val *string)
	CaCertificateIdentifierInput() *string
	CaCertificateValidTill() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterIdentifier() *string
	SetClusterIdentifier(val *string)
	ClusterIdentifierInput() *string
	ClusterIdentifierPrefix() *string
	SetClusterIdentifierPrefix(val *string)
	ClusterIdentifierPrefixInput() *string
	ClusterMembers() *[]*string
	SetClusterMembers(val *[]*string)
	ClusterMembersInput() *[]*string
	ClusterResourceId() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToSnapshot() interface{}
	SetCopyTagsToSnapshot(val interface{})
	CopyTagsToSnapshotInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DbClusterInstanceClass() *string
	SetDbClusterInstanceClass(val *string)
	DbClusterInstanceClassInput() *string
	DbClusterParameterGroupName() *string
	SetDbClusterParameterGroupName(val *string)
	DbClusterParameterGroupNameInput() *string
	DbInstanceParameterGroupName() *string
	SetDbInstanceParameterGroupName(val *string)
	DbInstanceParameterGroupNameInput() *string
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	DbSubnetGroupNameInput() *string
	DbSystemId() *string
	SetDbSystemId(val *string)
	DbSystemIdInput() *string
	DeleteAutomatedBackups() interface{}
	SetDeleteAutomatedBackups(val interface{})
	DeleteAutomatedBackupsInput() interface{}
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Domain() *string
	SetDomain(val *string)
	DomainIamRoleName() *string
	SetDomainIamRoleName(val *string)
	DomainIamRoleNameInput() *string
	DomainInput() *string
	EnabledCloudwatchLogsExports() *[]*string
	SetEnabledCloudwatchLogsExports(val *[]*string)
	EnabledCloudwatchLogsExportsInput() *[]*string
	EnableGlobalWriteForwarding() interface{}
	SetEnableGlobalWriteForwarding(val interface{})
	EnableGlobalWriteForwardingInput() interface{}
	EnableHttpEndpoint() interface{}
	SetEnableHttpEndpoint(val interface{})
	EnableHttpEndpointInput() interface{}
	EnableLocalWriteForwarding() interface{}
	SetEnableLocalWriteForwarding(val interface{})
	EnableLocalWriteForwardingInput() interface{}
	Endpoint() *string
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
	EngineLifecycleSupport() *string
	SetEngineLifecycleSupport(val *string)
	EngineLifecycleSupportInput() *string
	EngineMode() *string
	SetEngineMode(val *string)
	EngineModeInput() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionActual() *string
	EngineVersionInput() *string
	FinalSnapshotIdentifier() *string
	SetFinalSnapshotIdentifier(val *string)
	FinalSnapshotIdentifierInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlobalClusterIdentifier() *string
	SetGlobalClusterIdentifier(val *string)
	GlobalClusterIdentifierInput() *string
	HostedZoneId() *string
	IamDatabaseAuthenticationEnabled() interface{}
	SetIamDatabaseAuthenticationEnabled(val interface{})
	IamDatabaseAuthenticationEnabledInput() interface{}
	IamRoles() *[]*string
	SetIamRoles(val *[]*string)
	IamRolesInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManageMasterUserPassword() interface{}
	SetManageMasterUserPassword(val interface{})
	ManageMasterUserPasswordInput() interface{}
	MasterPassword() *string
	SetMasterPassword(val *string)
	MasterPasswordInput() *string
	MasterUsername() *string
	SetMasterUsername(val *string)
	MasterUsernameInput() *string
	MasterUserSecret() RdsClusterMasterUserSecretList
	MasterUserSecretKmsKeyId() *string
	SetMasterUserSecretKmsKeyId(val *string)
	MasterUserSecretKmsKeyIdInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
	// The tree node.
	Node() constructs.Node
	PerformanceInsightsEnabled() interface{}
	SetPerformanceInsightsEnabled(val interface{})
	PerformanceInsightsEnabledInput() interface{}
	PerformanceInsightsKmsKeyId() *string
	SetPerformanceInsightsKmsKeyId(val *string)
	PerformanceInsightsKmsKeyIdInput() *string
	PerformanceInsightsRetentionPeriod() *float64
	SetPerformanceInsightsRetentionPeriod(val *float64)
	PerformanceInsightsRetentionPeriodInput() *float64
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	PreferredBackupWindowInput() *string
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	PreferredMaintenanceWindowInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	ReaderEndpoint() *string
	ReplicationSourceIdentifier() *string
	SetReplicationSourceIdentifier(val *string)
	ReplicationSourceIdentifierInput() *string
	RestoreToPointInTime() RdsClusterRestoreToPointInTimeOutputReference
	RestoreToPointInTimeInput() *RdsClusterRestoreToPointInTime
	S3Import() RdsClusterS3ImportOutputReference
	S3ImportInput() *RdsClusterS3Import
	ScalingConfiguration() RdsClusterScalingConfigurationOutputReference
	ScalingConfigurationInput() *RdsClusterScalingConfiguration
	Serverlessv2ScalingConfiguration() RdsClusterServerlessv2ScalingConfigurationOutputReference
	Serverlessv2ScalingConfigurationInput() *RdsClusterServerlessv2ScalingConfiguration
	SkipFinalSnapshot() interface{}
	SetSkipFinalSnapshot(val interface{})
	SkipFinalSnapshotInput() interface{}
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	SnapshotIdentifierInput() *string
	SourceRegion() *string
	SetSourceRegion(val *string)
	SourceRegionInput() *string
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	StorageEncryptedInput() interface{}
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() RdsClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSecurityGroupIdsInput() *[]*string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutRestoreToPointInTime(value *RdsClusterRestoreToPointInTime)
	PutS3Import(value *RdsClusterS3Import)
	PutScalingConfiguration(value *RdsClusterScalingConfiguration)
	PutServerlessv2ScalingConfiguration(value *RdsClusterServerlessv2ScalingConfiguration)
	PutTimeouts(value *RdsClusterTimeouts)
	ResetAllocatedStorage()
	ResetAllowMajorVersionUpgrade()
	ResetApplyImmediately()
	ResetAvailabilityZones()
	ResetBacktrackWindow()
	ResetBackupRetentionPeriod()
	ResetCaCertificateIdentifier()
	ResetClusterIdentifier()
	ResetClusterIdentifierPrefix()
	ResetClusterMembers()
	ResetCopyTagsToSnapshot()
	ResetDatabaseName()
	ResetDbClusterInstanceClass()
	ResetDbClusterParameterGroupName()
	ResetDbInstanceParameterGroupName()
	ResetDbSubnetGroupName()
	ResetDbSystemId()
	ResetDeleteAutomatedBackups()
	ResetDeletionProtection()
	ResetDomain()
	ResetDomainIamRoleName()
	ResetEnabledCloudwatchLogsExports()
	ResetEnableGlobalWriteForwarding()
	ResetEnableHttpEndpoint()
	ResetEnableLocalWriteForwarding()
	ResetEngineLifecycleSupport()
	ResetEngineMode()
	ResetEngineVersion()
	ResetFinalSnapshotIdentifier()
	ResetGlobalClusterIdentifier()
	ResetIamDatabaseAuthenticationEnabled()
	ResetIamRoles()
	ResetId()
	ResetIops()
	ResetKmsKeyId()
	ResetManageMasterUserPassword()
	ResetMasterPassword()
	ResetMasterUsername()
	ResetMasterUserSecretKmsKeyId()
	ResetNetworkType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerformanceInsightsEnabled()
	ResetPerformanceInsightsKmsKeyId()
	ResetPerformanceInsightsRetentionPeriod()
	ResetPort()
	ResetPreferredBackupWindow()
	ResetPreferredMaintenanceWindow()
	ResetReplicationSourceIdentifier()
	ResetRestoreToPointInTime()
	ResetS3Import()
	ResetScalingConfiguration()
	ResetServerlessv2ScalingConfiguration()
	ResetSkipFinalSnapshot()
	ResetSnapshotIdentifier()
	ResetSourceRegion()
	ResetStorageEncrypted()
	ResetStorageType()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVpcSecurityGroupIds()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for RdsCluster
type jsiiProxy_RdsCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RdsCluster) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) AllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) AllowMajorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) AvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) AvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"availabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) BacktrackWindow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backtrackWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) BacktrackWindowInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backtrackWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) BackupRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) CaCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) CaCertificateIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) CaCertificateValidTill() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateValidTill",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ClusterIdentifierPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ClusterIdentifierPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ClusterMembers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterMembers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ClusterMembersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"clusterMembersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ClusterResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) CopyTagsToSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DbClusterInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DbClusterInstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterInstanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DbClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DbClusterParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DbInstanceParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DbInstanceParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DbSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DbSystemIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DeleteAutomatedBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DeleteAutomatedBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DomainIamRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EnabledCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EnabledCloudwatchLogsExportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EnableGlobalWriteForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGlobalWriteForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EnableGlobalWriteForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableGlobalWriteForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EnableHttpEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EnableHttpEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHttpEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EnableLocalWriteForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalWriteForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EnableLocalWriteForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableLocalWriteForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EngineLifecycleSupport() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EngineLifecycleSupportInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineLifecycleSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EngineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EngineModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EngineVersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) FinalSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) FinalSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) GlobalClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) GlobalClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) IamDatabaseAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) IamDatabaseAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) IamRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) IamRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ManageMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ManageMasterUserPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) MasterPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) MasterPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) MasterUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) MasterUserSecret() RdsClusterMasterUserSecretList {
	var returns RdsClusterMasterUserSecretList
	_jsii_.Get(
		j,
		"masterUserSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) MasterUserSecretKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) MasterUserSecretKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) PerformanceInsightsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) PerformanceInsightsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) PerformanceInsightsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) PerformanceInsightsRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) PreferredBackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ReaderEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ReplicationSourceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSourceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ReplicationSourceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationSourceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) RestoreToPointInTime() RdsClusterRestoreToPointInTimeOutputReference {
	var returns RdsClusterRestoreToPointInTimeOutputReference
	_jsii_.Get(
		j,
		"restoreToPointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) RestoreToPointInTimeInput() *RdsClusterRestoreToPointInTime {
	var returns *RdsClusterRestoreToPointInTime
	_jsii_.Get(
		j,
		"restoreToPointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) S3Import() RdsClusterS3ImportOutputReference {
	var returns RdsClusterS3ImportOutputReference
	_jsii_.Get(
		j,
		"s3Import",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) S3ImportInput() *RdsClusterS3Import {
	var returns *RdsClusterS3Import
	_jsii_.Get(
		j,
		"s3ImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ScalingConfiguration() RdsClusterScalingConfigurationOutputReference {
	var returns RdsClusterScalingConfigurationOutputReference
	_jsii_.Get(
		j,
		"scalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) ScalingConfigurationInput() *RdsClusterScalingConfiguration {
	var returns *RdsClusterScalingConfiguration
	_jsii_.Get(
		j,
		"scalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Serverlessv2ScalingConfiguration() RdsClusterServerlessv2ScalingConfigurationOutputReference {
	var returns RdsClusterServerlessv2ScalingConfigurationOutputReference
	_jsii_.Get(
		j,
		"serverlessv2ScalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Serverlessv2ScalingConfigurationInput() *RdsClusterServerlessv2ScalingConfiguration {
	var returns *RdsClusterServerlessv2ScalingConfiguration
	_jsii_.Get(
		j,
		"serverlessv2ScalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) SkipFinalSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) SkipFinalSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) SnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) SourceRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) StorageEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) Timeouts() RdsClusterTimeoutsOutputReference {
	var returns RdsClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsCluster) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/rds_cluster aws_rds_cluster} Resource.
func NewRdsCluster(scope constructs.Construct, id *string, config *RdsClusterConfig) RdsCluster {
	_init_.Initialize()

	if err := validateNewRdsClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsCluster{}

	_jsii_.Create(
		"@cdktf/provider-aws.rdsCluster.RdsCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/rds_cluster aws_rds_cluster} Resource.
func NewRdsCluster_Override(r RdsCluster, scope constructs.Construct, id *string, config *RdsClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.rdsCluster.RdsCluster",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RdsCluster)SetAllocatedStorage(val *float64) {
	if err := j.validateSetAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetAllowMajorVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowMajorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetAvailabilityZones(val *[]*string) {
	if err := j.validateSetAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZones",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetBacktrackWindow(val *float64) {
	if err := j.validateSetBacktrackWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backtrackWindow",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetBackupRetentionPeriod(val *float64) {
	if err := j.validateSetBackupRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetCaCertificateIdentifier(val *string) {
	if err := j.validateSetCaCertificateIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetClusterIdentifier(val *string) {
	if err := j.validateSetClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetClusterIdentifierPrefix(val *string) {
	if err := j.validateSetClusterIdentifierPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIdentifierPrefix",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetClusterMembers(val *[]*string) {
	if err := j.validateSetClusterMembersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterMembers",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetDbClusterInstanceClass(val *string) {
	if err := j.validateSetDbClusterInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterInstanceClass",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetDbClusterParameterGroupName(val *string) {
	if err := j.validateSetDbClusterParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetDbInstanceParameterGroupName(val *string) {
	if err := j.validateSetDbInstanceParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbInstanceParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetDbSubnetGroupName(val *string) {
	if err := j.validateSetDbSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetDbSystemId(val *string) {
	if err := j.validateSetDbSystemIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSystemId",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetDeleteAutomatedBackups(val interface{}) {
	if err := j.validateSetDeleteAutomatedBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAutomatedBackups",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetDomainIamRoleName(val *string) {
	if err := j.validateSetDomainIamRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetEnabledCloudwatchLogsExports(val *[]*string) {
	if err := j.validateSetEnabledCloudwatchLogsExportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetEnableGlobalWriteForwarding(val interface{}) {
	if err := j.validateSetEnableGlobalWriteForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableGlobalWriteForwarding",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetEnableHttpEndpoint(val interface{}) {
	if err := j.validateSetEnableHttpEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHttpEndpoint",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetEnableLocalWriteForwarding(val interface{}) {
	if err := j.validateSetEnableLocalWriteForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableLocalWriteForwarding",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetEngineLifecycleSupport(val *string) {
	if err := j.validateSetEngineLifecycleSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineLifecycleSupport",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetEngineMode(val *string) {
	if err := j.validateSetEngineModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineMode",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetFinalSnapshotIdentifier(val *string) {
	if err := j.validateSetFinalSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetGlobalClusterIdentifier(val *string) {
	if err := j.validateSetGlobalClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetIamDatabaseAuthenticationEnabled(val interface{}) {
	if err := j.validateSetIamDatabaseAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamDatabaseAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetIamRoles(val *[]*string) {
	if err := j.validateSetIamRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoles",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetManageMasterUserPassword(val interface{}) {
	if err := j.validateSetManageMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterUserPassword",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetMasterPassword(val *string) {
	if err := j.validateSetMasterPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPassword",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetMasterUsername(val *string) {
	if err := j.validateSetMasterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetMasterUserSecretKmsKeyId(val *string) {
	if err := j.validateSetMasterUserSecretKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserSecretKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetPerformanceInsightsEnabled(val interface{}) {
	if err := j.validateSetPerformanceInsightsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsEnabled",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetPerformanceInsightsKmsKeyId(val *string) {
	if err := j.validateSetPerformanceInsightsKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetPerformanceInsightsRetentionPeriod(val *float64) {
	if err := j.validateSetPerformanceInsightsRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetPreferredBackupWindow(val *string) {
	if err := j.validateSetPreferredBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetReplicationSourceIdentifier(val *string) {
	if err := j.validateSetReplicationSourceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationSourceIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetSkipFinalSnapshot(val interface{}) {
	if err := j.validateSetSkipFinalSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalSnapshot",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetSnapshotIdentifier(val *string) {
	if err := j.validateSetSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetSourceRegion(val *string) {
	if err := j.validateSetSourceRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRegion",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_RdsCluster)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTF code for importing a RdsCluster resource upon running "cdktf plan <stack-name>".
func RdsCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRdsCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.rdsCluster.RdsCluster",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func RdsCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.rdsCluster.RdsCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.rdsCluster.RdsCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.rdsCluster.RdsCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RdsCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.rdsCluster.RdsCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RdsCluster) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RdsCluster) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RdsCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RdsCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RdsCluster) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsCluster) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RdsCluster) PutRestoreToPointInTime(value *RdsClusterRestoreToPointInTime) {
	if err := r.validatePutRestoreToPointInTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRestoreToPointInTime",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsCluster) PutS3Import(value *RdsClusterS3Import) {
	if err := r.validatePutS3ImportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putS3Import",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsCluster) PutScalingConfiguration(value *RdsClusterScalingConfiguration) {
	if err := r.validatePutScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putScalingConfiguration",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsCluster) PutServerlessv2ScalingConfiguration(value *RdsClusterServerlessv2ScalingConfiguration) {
	if err := r.validatePutServerlessv2ScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putServerlessv2ScalingConfiguration",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsCluster) PutTimeouts(value *RdsClusterTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsCluster) ResetAllocatedStorage() {
	_jsii_.InvokeVoid(
		r,
		"resetAllocatedStorage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetAllowMajorVersionUpgrade() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowMajorVersionUpgrade",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		r,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetAvailabilityZones() {
	_jsii_.InvokeVoid(
		r,
		"resetAvailabilityZones",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetBacktrackWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetBacktrackWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetBackupRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetBackupRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetCaCertificateIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetCaCertificateIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetClusterIdentifierPrefix() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterIdentifierPrefix",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetClusterMembers() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterMembers",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetCopyTagsToSnapshot() {
	_jsii_.InvokeVoid(
		r,
		"resetCopyTagsToSnapshot",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		r,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetDbClusterInstanceClass() {
	_jsii_.InvokeVoid(
		r,
		"resetDbClusterInstanceClass",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetDbClusterParameterGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbClusterParameterGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetDbInstanceParameterGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbInstanceParameterGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetDbSubnetGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSubnetGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetDbSystemId() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSystemId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetDeleteAutomatedBackups() {
	_jsii_.InvokeVoid(
		r,
		"resetDeleteAutomatedBackups",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		r,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetDomain() {
	_jsii_.InvokeVoid(
		r,
		"resetDomain",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetDomainIamRoleName() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainIamRoleName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetEnabledCloudwatchLogsExports() {
	_jsii_.InvokeVoid(
		r,
		"resetEnabledCloudwatchLogsExports",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetEnableGlobalWriteForwarding() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableGlobalWriteForwarding",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetEnableHttpEndpoint() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableHttpEndpoint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetEnableLocalWriteForwarding() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableLocalWriteForwarding",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetEngineLifecycleSupport() {
	_jsii_.InvokeVoid(
		r,
		"resetEngineLifecycleSupport",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetEngineMode() {
	_jsii_.InvokeVoid(
		r,
		"resetEngineMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetFinalSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetFinalSnapshotIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetGlobalClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetGlobalClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetIamDatabaseAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetIamDatabaseAuthenticationEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetIamRoles() {
	_jsii_.InvokeVoid(
		r,
		"resetIamRoles",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetIops() {
	_jsii_.InvokeVoid(
		r,
		"resetIops",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetManageMasterUserPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetManageMasterUserPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetMasterPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetMasterUsername() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUsername",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetMasterUserSecretKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUserSecretKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetNetworkType() {
	_jsii_.InvokeVoid(
		r,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetPerformanceInsightsEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetPerformanceInsightsEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetPerformanceInsightsKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetPerformanceInsightsKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetPerformanceInsightsRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetPerformanceInsightsRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetPort() {
	_jsii_.InvokeVoid(
		r,
		"resetPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetPreferredBackupWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetPreferredBackupWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetReplicationSourceIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetReplicationSourceIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetRestoreToPointInTime() {
	_jsii_.InvokeVoid(
		r,
		"resetRestoreToPointInTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetS3Import() {
	_jsii_.InvokeVoid(
		r,
		"resetS3Import",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetScalingConfiguration() {
	_jsii_.InvokeVoid(
		r,
		"resetScalingConfiguration",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetServerlessv2ScalingConfiguration() {
	_jsii_.InvokeVoid(
		r,
		"resetServerlessv2ScalingConfiguration",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetSkipFinalSnapshot() {
	_jsii_.InvokeVoid(
		r,
		"resetSkipFinalSnapshot",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSnapshotIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetSourceRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetStorageEncrypted() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageEncrypted",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetStorageType() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		r,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		r,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


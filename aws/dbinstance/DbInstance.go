// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dbinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/db_instance aws_db_instance}.
type DbInstance interface {
	cdktf.TerraformResource
	Address() *string
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
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	AutoMinorVersionUpgradeInput() interface{}
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	BackupRetentionPeriodInput() *float64
	BackupTarget() *string
	SetBackupTarget(val *string)
	BackupTargetInput() *string
	BackupWindow() *string
	SetBackupWindow(val *string)
	BackupWindowInput() *string
	BlueGreenUpdate() DbInstanceBlueGreenUpdateOutputReference
	BlueGreenUpdateInput() *DbInstanceBlueGreenUpdate
	CaCertIdentifier() *string
	SetCaCertIdentifier(val *string)
	CaCertIdentifierInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CharacterSetName() *string
	SetCharacterSetName(val *string)
	CharacterSetNameInput() *string
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
	CustomerOwnedIpEnabled() interface{}
	SetCustomerOwnedIpEnabled(val interface{})
	CustomerOwnedIpEnabledInput() interface{}
	CustomIamInstanceProfile() *string
	SetCustomIamInstanceProfile(val *string)
	CustomIamInstanceProfileInput() *string
	DbName() *string
	SetDbName(val *string)
	DbNameInput() *string
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	DbSubnetGroupNameInput() *string
	DedicatedLogVolume() interface{}
	SetDedicatedLogVolume(val interface{})
	DedicatedLogVolumeInput() interface{}
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
	DomainAuthSecretArn() *string
	SetDomainAuthSecretArn(val *string)
	DomainAuthSecretArnInput() *string
	DomainDnsIps() *[]*string
	SetDomainDnsIps(val *[]*string)
	DomainDnsIpsInput() *[]*string
	DomainFqdn() *string
	SetDomainFqdn(val *string)
	DomainFqdnInput() *string
	DomainIamRoleName() *string
	SetDomainIamRoleName(val *string)
	DomainIamRoleNameInput() *string
	DomainInput() *string
	DomainOu() *string
	SetDomainOu(val *string)
	DomainOuInput() *string
	EnabledCloudwatchLogsExports() *[]*string
	SetEnabledCloudwatchLogsExports(val *[]*string)
	EnabledCloudwatchLogsExportsInput() *[]*string
	Endpoint() *string
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
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
	HostedZoneId() *string
	IamDatabaseAuthenticationEnabled() interface{}
	SetIamDatabaseAuthenticationEnabled(val interface{})
	IamDatabaseAuthenticationEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	Identifier() *string
	SetIdentifier(val *string)
	IdentifierInput() *string
	IdentifierPrefix() *string
	SetIdentifierPrefix(val *string)
	IdentifierPrefixInput() *string
	IdInput() *string
	InstanceClass() *string
	SetInstanceClass(val *string)
	InstanceClassInput() *string
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	LatestRestorableTime() *string
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListenerEndpoint() DbInstanceListenerEndpointList
	MaintenanceWindow() *string
	SetMaintenanceWindow(val *string)
	MaintenanceWindowInput() *string
	ManageMasterUserPassword() interface{}
	SetManageMasterUserPassword(val interface{})
	ManageMasterUserPasswordInput() interface{}
	MasterUserSecret() DbInstanceMasterUserSecretList
	MasterUserSecretKmsKeyId() *string
	SetMasterUserSecretKmsKeyId(val *string)
	MasterUserSecretKmsKeyIdInput() *string
	MaxAllocatedStorage() *float64
	SetMaxAllocatedStorage(val *float64)
	MaxAllocatedStorageInput() *float64
	MonitoringInterval() *float64
	SetMonitoringInterval(val *float64)
	MonitoringIntervalInput() *float64
	MonitoringRoleArn() *string
	SetMonitoringRoleArn(val *string)
	MonitoringRoleArnInput() *string
	MultiAz() interface{}
	SetMultiAz(val interface{})
	MultiAzInput() interface{}
	NcharCharacterSetName() *string
	SetNcharCharacterSetName(val *string)
	NcharCharacterSetNameInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
	// The tree node.
	Node() constructs.Node
	OptionGroupName() *string
	SetOptionGroupName(val *string)
	OptionGroupNameInput() *string
	ParameterGroupName() *string
	SetParameterGroupName(val *string)
	ParameterGroupNameInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
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
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	PubliclyAccessibleInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReplicaMode() *string
	SetReplicaMode(val *string)
	ReplicaModeInput() *string
	Replicas() *[]*string
	ReplicateSourceDb() *string
	SetReplicateSourceDb(val *string)
	ReplicateSourceDbInput() *string
	ResourceId() *string
	RestoreToPointInTime() DbInstanceRestoreToPointInTimeOutputReference
	RestoreToPointInTimeInput() *DbInstanceRestoreToPointInTime
	S3Import() DbInstanceS3ImportOutputReference
	S3ImportInput() *DbInstanceS3Import
	SkipFinalSnapshot() interface{}
	SetSkipFinalSnapshot(val interface{})
	SkipFinalSnapshotInput() interface{}
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	SnapshotIdentifierInput() *string
	Status() *string
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	StorageEncryptedInput() interface{}
	StorageThroughput() *float64
	SetStorageThroughput(val *float64)
	StorageThroughputInput() *float64
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
	Timeouts() DbInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	PutBlueGreenUpdate(value *DbInstanceBlueGreenUpdate)
	PutRestoreToPointInTime(value *DbInstanceRestoreToPointInTime)
	PutS3Import(value *DbInstanceS3Import)
	PutTimeouts(value *DbInstanceTimeouts)
	ResetAllocatedStorage()
	ResetAllowMajorVersionUpgrade()
	ResetApplyImmediately()
	ResetAutoMinorVersionUpgrade()
	ResetAvailabilityZone()
	ResetBackupRetentionPeriod()
	ResetBackupTarget()
	ResetBackupWindow()
	ResetBlueGreenUpdate()
	ResetCaCertIdentifier()
	ResetCharacterSetName()
	ResetCopyTagsToSnapshot()
	ResetCustomerOwnedIpEnabled()
	ResetCustomIamInstanceProfile()
	ResetDbName()
	ResetDbSubnetGroupName()
	ResetDedicatedLogVolume()
	ResetDeleteAutomatedBackups()
	ResetDeletionProtection()
	ResetDomain()
	ResetDomainAuthSecretArn()
	ResetDomainDnsIps()
	ResetDomainFqdn()
	ResetDomainIamRoleName()
	ResetDomainOu()
	ResetEnabledCloudwatchLogsExports()
	ResetEngine()
	ResetEngineVersion()
	ResetFinalSnapshotIdentifier()
	ResetIamDatabaseAuthenticationEnabled()
	ResetId()
	ResetIdentifier()
	ResetIdentifierPrefix()
	ResetIops()
	ResetKmsKeyId()
	ResetLicenseModel()
	ResetMaintenanceWindow()
	ResetManageMasterUserPassword()
	ResetMasterUserSecretKmsKeyId()
	ResetMaxAllocatedStorage()
	ResetMonitoringInterval()
	ResetMonitoringRoleArn()
	ResetMultiAz()
	ResetNcharCharacterSetName()
	ResetNetworkType()
	ResetOptionGroupName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameterGroupName()
	ResetPassword()
	ResetPerformanceInsightsEnabled()
	ResetPerformanceInsightsKmsKeyId()
	ResetPerformanceInsightsRetentionPeriod()
	ResetPort()
	ResetPubliclyAccessible()
	ResetReplicaMode()
	ResetReplicateSourceDb()
	ResetRestoreToPointInTime()
	ResetS3Import()
	ResetSkipFinalSnapshot()
	ResetSnapshotIdentifier()
	ResetStorageEncrypted()
	ResetStorageThroughput()
	ResetStorageType()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetTimezone()
	ResetUsername()
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

// The jsii proxy struct for DbInstance
type jsiiProxy_DbInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DbInstance) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) AllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) AllowMajorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) BackupRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) BackupTarget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) BackupTargetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) BackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) BackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) BlueGreenUpdate() DbInstanceBlueGreenUpdateOutputReference {
	var returns DbInstanceBlueGreenUpdateOutputReference
	_jsii_.Get(
		j,
		"blueGreenUpdate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) BlueGreenUpdateInput() *DbInstanceBlueGreenUpdate {
	var returns *DbInstanceBlueGreenUpdate
	_jsii_.Get(
		j,
		"blueGreenUpdateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) CaCertIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) CaCertIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) CharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) CharacterSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) CopyTagsToSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) CustomerOwnedIpEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerOwnedIpEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) CustomerOwnedIpEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customerOwnedIpEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) CustomIamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) CustomIamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DbSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DedicatedLogVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedLogVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DedicatedLogVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedLogVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DeleteAutomatedBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DeleteAutomatedBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DomainAuthSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DomainAuthSecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DomainDnsIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DomainDnsIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DomainFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DomainFqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DomainIamRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DomainOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) DomainOuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) EnabledCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) EnabledCloudwatchLogsExportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enabledCloudwatchLogsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) EngineVersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) FinalSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) FinalSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) HostedZoneId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostedZoneId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) IamDatabaseAuthenticationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) IamDatabaseAuthenticationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iamDatabaseAuthenticationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Identifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) IdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) IdentifierPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) IdentifierPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identifierPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) InstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) InstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) LatestRestorableTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"latestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ListenerEndpoint() DbInstanceListenerEndpointList {
	var returns DbInstanceListenerEndpointList
	_jsii_.Get(
		j,
		"listenerEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ManageMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ManageMasterUserPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MasterUserSecret() DbInstanceMasterUserSecretList {
	var returns DbInstanceMasterUserSecretList
	_jsii_.Get(
		j,
		"masterUserSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MasterUserSecretKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MasterUserSecretKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserSecretKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MaxAllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MaxAllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MonitoringIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MonitoringRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) MultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) NcharCharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) NcharCharacterSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) OptionGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) PerformanceInsightsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) PerformanceInsightsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceInsightsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) PerformanceInsightsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) PerformanceInsightsRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ReplicaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ReplicaModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Replicas() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"replicas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ReplicateSourceDb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicateSourceDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ReplicateSourceDbInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicateSourceDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) RestoreToPointInTime() DbInstanceRestoreToPointInTimeOutputReference {
	var returns DbInstanceRestoreToPointInTimeOutputReference
	_jsii_.Get(
		j,
		"restoreToPointInTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) RestoreToPointInTimeInput() *DbInstanceRestoreToPointInTime {
	var returns *DbInstanceRestoreToPointInTime
	_jsii_.Get(
		j,
		"restoreToPointInTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) S3Import() DbInstanceS3ImportOutputReference {
	var returns DbInstanceS3ImportOutputReference
	_jsii_.Get(
		j,
		"s3Import",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) S3ImportInput() *DbInstanceS3Import {
	var returns *DbInstanceS3Import
	_jsii_.Get(
		j,
		"s3ImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) SkipFinalSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) SkipFinalSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) SnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) StorageEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) StorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) StorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Timeouts() DbInstanceTimeoutsOutputReference {
	var returns DbInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbInstance) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/db_instance aws_db_instance} Resource.
func NewDbInstance(scope constructs.Construct, id *string, config *DbInstanceConfig) DbInstance {
	_init_.Initialize()

	if err := validateNewDbInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DbInstance{}

	_jsii_.Create(
		"@cdktf/provider-aws.dbInstance.DbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/db_instance aws_db_instance} Resource.
func NewDbInstance_Override(d DbInstance, scope constructs.Construct, id *string, config *DbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dbInstance.DbInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DbInstance)SetAllocatedStorage(val *float64) {
	if err := j.validateSetAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetAllowMajorVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowMajorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetBackupRetentionPeriod(val *float64) {
	if err := j.validateSetBackupRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetBackupTarget(val *string) {
	if err := j.validateSetBackupTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupTarget",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetBackupWindow(val *string) {
	if err := j.validateSetBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupWindow",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetCaCertIdentifier(val *string) {
	if err := j.validateSetCaCertIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertIdentifier",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetCharacterSetName(val *string) {
	if err := j.validateSetCharacterSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSetName",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetCustomerOwnedIpEnabled(val interface{}) {
	if err := j.validateSetCustomerOwnedIpEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerOwnedIpEnabled",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetCustomIamInstanceProfile(val *string) {
	if err := j.validateSetCustomIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customIamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDbSubnetGroupName(val *string) {
	if err := j.validateSetDbSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDedicatedLogVolume(val interface{}) {
	if err := j.validateSetDedicatedLogVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedLogVolume",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDeleteAutomatedBackups(val interface{}) {
	if err := j.validateSetDeleteAutomatedBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAutomatedBackups",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDomainAuthSecretArn(val *string) {
	if err := j.validateSetDomainAuthSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAuthSecretArn",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDomainDnsIps(val *[]*string) {
	if err := j.validateSetDomainDnsIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainDnsIps",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDomainFqdn(val *string) {
	if err := j.validateSetDomainFqdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainFqdn",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDomainIamRoleName(val *string) {
	if err := j.validateSetDomainIamRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetDomainOu(val *string) {
	if err := j.validateSetDomainOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainOu",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetEnabledCloudwatchLogsExports(val *[]*string) {
	if err := j.validateSetEnabledCloudwatchLogsExportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetFinalSnapshotIdentifier(val *string) {
	if err := j.validateSetFinalSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetIamDatabaseAuthenticationEnabled(val interface{}) {
	if err := j.validateSetIamDatabaseAuthenticationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamDatabaseAuthenticationEnabled",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetIdentifier(val *string) {
	if err := j.validateSetIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifier",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetIdentifierPrefix(val *string) {
	if err := j.validateSetIdentifierPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identifierPrefix",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetInstanceClass(val *string) {
	if err := j.validateSetInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceClass",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetMaintenanceWindow(val *string) {
	if err := j.validateSetMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetManageMasterUserPassword(val interface{}) {
	if err := j.validateSetManageMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterUserPassword",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetMasterUserSecretKmsKeyId(val *string) {
	if err := j.validateSetMasterUserSecretKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserSecretKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetMaxAllocatedStorage(val *float64) {
	if err := j.validateSetMaxAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllocatedStorage",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetMonitoringInterval(val *float64) {
	if err := j.validateSetMonitoringIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInterval",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetMonitoringRoleArn(val *string) {
	if err := j.validateSetMonitoringRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringRoleArn",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetMultiAz(val interface{}) {
	if err := j.validateSetMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetNcharCharacterSetName(val *string) {
	if err := j.validateSetNcharCharacterSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ncharCharacterSetName",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetOptionGroupName(val *string) {
	if err := j.validateSetOptionGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionGroupName",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetParameterGroupName(val *string) {
	if err := j.validateSetParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterGroupName",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetPerformanceInsightsEnabled(val interface{}) {
	if err := j.validateSetPerformanceInsightsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsEnabled",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetPerformanceInsightsKmsKeyId(val *string) {
	if err := j.validateSetPerformanceInsightsKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetPerformanceInsightsRetentionPeriod(val *float64) {
	if err := j.validateSetPerformanceInsightsRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetReplicaMode(val *string) {
	if err := j.validateSetReplicaModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaMode",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetReplicateSourceDb(val *string) {
	if err := j.validateSetReplicateSourceDbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicateSourceDb",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetSkipFinalSnapshot(val interface{}) {
	if err := j.validateSetSkipFinalSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalSnapshot",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetSnapshotIdentifier(val *string) {
	if err := j.validateSetSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetStorageThroughput(val *float64) {
	if err := j.validateSetStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageThroughput",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_DbInstance)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
}

// Generates CDKTF code for importing a DbInstance resource upon running "cdktf plan <stack-name>".
func DbInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDbInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dbInstance.DbInstance",
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
func DbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDbInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dbInstance.DbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DbInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDbInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dbInstance.DbInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DbInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDbInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dbInstance.DbInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dbInstance.DbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DbInstance) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DbInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DbInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DbInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DbInstance) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DbInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DbInstance) PutBlueGreenUpdate(value *DbInstanceBlueGreenUpdate) {
	if err := d.validatePutBlueGreenUpdateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putBlueGreenUpdate",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DbInstance) PutRestoreToPointInTime(value *DbInstanceRestoreToPointInTime) {
	if err := d.validatePutRestoreToPointInTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRestoreToPointInTime",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DbInstance) PutS3Import(value *DbInstanceS3Import) {
	if err := d.validatePutS3ImportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putS3Import",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DbInstance) PutTimeouts(value *DbInstanceTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DbInstance) ResetAllocatedStorage() {
	_jsii_.InvokeVoid(
		d,
		"resetAllocatedStorage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetAllowMajorVersionUpgrade() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowMajorVersionUpgrade",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		d,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		d,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		d,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetBackupRetentionPeriod() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupRetentionPeriod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetBackupTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetBackupWindow() {
	_jsii_.InvokeVoid(
		d,
		"resetBackupWindow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetBlueGreenUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetBlueGreenUpdate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetCaCertIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetCaCertIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetCharacterSetName() {
	_jsii_.InvokeVoid(
		d,
		"resetCharacterSetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetCopyTagsToSnapshot() {
	_jsii_.InvokeVoid(
		d,
		"resetCopyTagsToSnapshot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetCustomerOwnedIpEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomerOwnedIpEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetCustomIamInstanceProfile() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomIamInstanceProfile",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetDbName() {
	_jsii_.InvokeVoid(
		d,
		"resetDbName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetDbSubnetGroupName() {
	_jsii_.InvokeVoid(
		d,
		"resetDbSubnetGroupName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetDedicatedLogVolume() {
	_jsii_.InvokeVoid(
		d,
		"resetDedicatedLogVolume",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetDeleteAutomatedBackups() {
	_jsii_.InvokeVoid(
		d,
		"resetDeleteAutomatedBackups",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetDomain() {
	_jsii_.InvokeVoid(
		d,
		"resetDomain",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetDomainAuthSecretArn() {
	_jsii_.InvokeVoid(
		d,
		"resetDomainAuthSecretArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetDomainDnsIps() {
	_jsii_.InvokeVoid(
		d,
		"resetDomainDnsIps",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetDomainFqdn() {
	_jsii_.InvokeVoid(
		d,
		"resetDomainFqdn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetDomainIamRoleName() {
	_jsii_.InvokeVoid(
		d,
		"resetDomainIamRoleName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetDomainOu() {
	_jsii_.InvokeVoid(
		d,
		"resetDomainOu",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetEnabledCloudwatchLogsExports() {
	_jsii_.InvokeVoid(
		d,
		"resetEnabledCloudwatchLogsExports",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetEngine() {
	_jsii_.InvokeVoid(
		d,
		"resetEngine",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetFinalSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetFinalSnapshotIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetIamDatabaseAuthenticationEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetIamDatabaseAuthenticationEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetIdentifierPrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetIdentifierPrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetIops() {
	_jsii_.InvokeVoid(
		d,
		"resetIops",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		d,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		d,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetManageMasterUserPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetManageMasterUserPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetMasterUserSecretKmsKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetMasterUserSecretKmsKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetMaxAllocatedStorage() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxAllocatedStorage",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetMonitoringInterval() {
	_jsii_.InvokeVoid(
		d,
		"resetMonitoringInterval",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetMonitoringRoleArn() {
	_jsii_.InvokeVoid(
		d,
		"resetMonitoringRoleArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetMultiAz() {
	_jsii_.InvokeVoid(
		d,
		"resetMultiAz",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetNcharCharacterSetName() {
	_jsii_.InvokeVoid(
		d,
		"resetNcharCharacterSetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetNetworkType() {
	_jsii_.InvokeVoid(
		d,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetOptionGroupName() {
	_jsii_.InvokeVoid(
		d,
		"resetOptionGroupName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetParameterGroupName() {
	_jsii_.InvokeVoid(
		d,
		"resetParameterGroupName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetPerformanceInsightsEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetPerformanceInsightsEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetPerformanceInsightsKmsKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetPerformanceInsightsKmsKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetPerformanceInsightsRetentionPeriod() {
	_jsii_.InvokeVoid(
		d,
		"resetPerformanceInsightsRetentionPeriod",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetPort() {
	_jsii_.InvokeVoid(
		d,
		"resetPort",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		d,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetReplicaMode() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicaMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetReplicateSourceDb() {
	_jsii_.InvokeVoid(
		d,
		"resetReplicateSourceDb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetRestoreToPointInTime() {
	_jsii_.InvokeVoid(
		d,
		"resetRestoreToPointInTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetS3Import() {
	_jsii_.InvokeVoid(
		d,
		"resetS3Import",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetSkipFinalSnapshot() {
	_jsii_.InvokeVoid(
		d,
		"resetSkipFinalSnapshot",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetSnapshotIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetStorageEncrypted() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageEncrypted",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetStorageThroughput() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageThroughput",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetStorageType() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetTimezone() {
	_jsii_.InvokeVoid(
		d,
		"resetTimezone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		d,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


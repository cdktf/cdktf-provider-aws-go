package redshiftcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v16/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v16/redshiftcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/redshift_cluster aws_redshift_cluster}.
type RedshiftCluster interface {
	cdktf.TerraformResource
	AllowVersionUpgrade() interface{}
	SetAllowVersionUpgrade(val interface{})
	AllowVersionUpgradeInput() interface{}
	ApplyImmediately() interface{}
	SetApplyImmediately(val interface{})
	ApplyImmediatelyInput() interface{}
	AquaConfigurationStatus() *string
	SetAquaConfigurationStatus(val *string)
	AquaConfigurationStatusInput() *string
	Arn() *string
	AutomatedSnapshotRetentionPeriod() *float64
	SetAutomatedSnapshotRetentionPeriod(val *float64)
	AutomatedSnapshotRetentionPeriodInput() *float64
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	AvailabilityZoneRelocationEnabled() interface{}
	SetAvailabilityZoneRelocationEnabled(val interface{})
	AvailabilityZoneRelocationEnabledInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterIdentifier() *string
	SetClusterIdentifier(val *string)
	ClusterIdentifierInput() *string
	ClusterNamespaceArn() *string
	ClusterNodes() RedshiftClusterClusterNodesList
	ClusterParameterGroupName() *string
	SetClusterParameterGroupName(val *string)
	ClusterParameterGroupNameInput() *string
	ClusterPublicKey() *string
	SetClusterPublicKey(val *string)
	ClusterPublicKeyInput() *string
	ClusterRevisionNumber() *string
	SetClusterRevisionNumber(val *string)
	ClusterRevisionNumberInput() *string
	ClusterSubnetGroupName() *string
	SetClusterSubnetGroupName(val *string)
	ClusterSubnetGroupNameInput() *string
	ClusterType() *string
	SetClusterType(val *string)
	ClusterTypeInput() *string
	ClusterVersion() *string
	SetClusterVersion(val *string)
	ClusterVersionInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	DefaultIamRoleArn() *string
	SetDefaultIamRoleArn(val *string)
	DefaultIamRoleArnInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DnsName() *string
	ElasticIp() *string
	SetElasticIp(val *string)
	ElasticIpInput() *string
	Encrypted() interface{}
	SetEncrypted(val interface{})
	EncryptedInput() interface{}
	Endpoint() *string
	SetEndpoint(val *string)
	EndpointInput() *string
	EnhancedVpcRouting() interface{}
	SetEnhancedVpcRouting(val interface{})
	EnhancedVpcRoutingInput() interface{}
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
	IamRoles() *[]*string
	SetIamRoles(val *[]*string)
	IamRolesInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Logging() RedshiftClusterLoggingOutputReference
	LoggingInput() *RedshiftClusterLogging
	MaintenanceTrackName() *string
	SetMaintenanceTrackName(val *string)
	MaintenanceTrackNameInput() *string
	ManualSnapshotRetentionPeriod() *float64
	SetManualSnapshotRetentionPeriod(val *float64)
	ManualSnapshotRetentionPeriodInput() *float64
	MasterPassword() *string
	SetMasterPassword(val *string)
	MasterPasswordInput() *string
	MasterUsername() *string
	SetMasterUsername(val *string)
	MasterUsernameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeType() *string
	SetNodeType(val *string)
	NodeTypeInput() *string
	NumberOfNodes() *float64
	SetNumberOfNodes(val *float64)
	NumberOfNodesInput() *float64
	OwnerAccount() *string
	SetOwnerAccount(val *string)
	OwnerAccountInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
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
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	PubliclyAccessibleInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	SkipFinalSnapshot() interface{}
	SetSkipFinalSnapshot(val interface{})
	SkipFinalSnapshotInput() interface{}
	SnapshotClusterIdentifier() *string
	SetSnapshotClusterIdentifier(val *string)
	SnapshotClusterIdentifierInput() *string
	SnapshotCopy() RedshiftClusterSnapshotCopyOutputReference
	SnapshotCopyInput() *RedshiftClusterSnapshotCopy
	SnapshotIdentifier() *string
	SetSnapshotIdentifier(val *string)
	SnapshotIdentifierInput() *string
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
	Timeouts() RedshiftClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcSecurityGroupIds() *[]*string
	SetVpcSecurityGroupIds(val *[]*string)
	VpcSecurityGroupIdsInput() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutLogging(value *RedshiftClusterLogging)
	PutSnapshotCopy(value *RedshiftClusterSnapshotCopy)
	PutTimeouts(value *RedshiftClusterTimeouts)
	ResetAllowVersionUpgrade()
	ResetApplyImmediately()
	ResetAquaConfigurationStatus()
	ResetAutomatedSnapshotRetentionPeriod()
	ResetAvailabilityZone()
	ResetAvailabilityZoneRelocationEnabled()
	ResetClusterParameterGroupName()
	ResetClusterPublicKey()
	ResetClusterRevisionNumber()
	ResetClusterSubnetGroupName()
	ResetClusterType()
	ResetClusterVersion()
	ResetDatabaseName()
	ResetDefaultIamRoleArn()
	ResetElasticIp()
	ResetEncrypted()
	ResetEndpoint()
	ResetEnhancedVpcRouting()
	ResetFinalSnapshotIdentifier()
	ResetIamRoles()
	ResetId()
	ResetKmsKeyId()
	ResetLogging()
	ResetMaintenanceTrackName()
	ResetManualSnapshotRetentionPeriod()
	ResetMasterPassword()
	ResetMasterUsername()
	ResetNumberOfNodes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwnerAccount()
	ResetPort()
	ResetPreferredMaintenanceWindow()
	ResetPubliclyAccessible()
	ResetSkipFinalSnapshot()
	ResetSnapshotClusterIdentifier()
	ResetSnapshotCopy()
	ResetSnapshotIdentifier()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetVpcSecurityGroupIds()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for RedshiftCluster
type jsiiProxy_RedshiftCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RedshiftCluster) AllowVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AllowVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AquaConfigurationStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aquaConfigurationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AquaConfigurationStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aquaConfigurationStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AutomatedSnapshotRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automatedSnapshotRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AutomatedSnapshotRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"automatedSnapshotRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AvailabilityZoneRelocationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneRelocationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) AvailabilityZoneRelocationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"availabilityZoneRelocationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterNodes() RedshiftClusterClusterNodesList {
	var returns RedshiftClusterClusterNodesList
	_jsii_.Get(
		j,
		"clusterNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterPublicKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterPublicKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterPublicKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterPublicKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterRevisionNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterRevisionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterRevisionNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterRevisionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ClusterVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DefaultIamRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIamRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DefaultIamRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultIamRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) DnsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ElasticIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ElasticIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) EncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Endpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) EndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) EnhancedVpcRouting() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedVpcRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) EnhancedVpcRoutingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedVpcRoutingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) FinalSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) FinalSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) IamRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) IamRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Logging() RedshiftClusterLoggingOutputReference {
	var returns RedshiftClusterLoggingOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) LoggingInput() *RedshiftClusterLogging {
	var returns *RedshiftClusterLogging
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MaintenanceTrackName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceTrackName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MaintenanceTrackNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceTrackNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ManualSnapshotRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manualSnapshotRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) ManualSnapshotRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"manualSnapshotRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MasterPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MasterPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) MasterUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) NumberOfNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) NumberOfNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) OwnerAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) OwnerAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SkipFinalSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SkipFinalSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotCopy() RedshiftClusterSnapshotCopyOutputReference {
	var returns RedshiftClusterSnapshotCopyOutputReference
	_jsii_.Get(
		j,
		"snapshotCopy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotCopyInput() *RedshiftClusterSnapshotCopy {
	var returns *RedshiftClusterSnapshotCopy
	_jsii_.Get(
		j,
		"snapshotCopyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) SnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) Timeouts() RedshiftClusterTimeoutsOutputReference {
	var returns RedshiftClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) VpcSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftCluster) VpcSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupIdsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/redshift_cluster aws_redshift_cluster} Resource.
func NewRedshiftCluster(scope constructs.Construct, id *string, config *RedshiftClusterConfig) RedshiftCluster {
	_init_.Initialize()

	if err := validateNewRedshiftClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftCluster{}

	_jsii_.Create(
		"@cdktf/provider-aws.redshiftCluster.RedshiftCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.10.0/docs/resources/redshift_cluster aws_redshift_cluster} Resource.
func NewRedshiftCluster_Override(r RedshiftCluster, scope constructs.Construct, id *string, config *RedshiftClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.redshiftCluster.RedshiftCluster",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetAllowVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetAquaConfigurationStatus(val *string) {
	if err := j.validateSetAquaConfigurationStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aquaConfigurationStatus",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetAutomatedSnapshotRetentionPeriod(val *float64) {
	if err := j.validateSetAutomatedSnapshotRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automatedSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetAvailabilityZoneRelocationEnabled(val interface{}) {
	if err := j.validateSetAvailabilityZoneRelocationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZoneRelocationEnabled",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterIdentifier(val *string) {
	if err := j.validateSetClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterParameterGroupName(val *string) {
	if err := j.validateSetClusterParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterPublicKey(val *string) {
	if err := j.validateSetClusterPublicKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterPublicKey",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterRevisionNumber(val *string) {
	if err := j.validateSetClusterRevisionNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterRevisionNumber",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterSubnetGroupName(val *string) {
	if err := j.validateSetClusterSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterType(val *string) {
	if err := j.validateSetClusterTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterType",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetClusterVersion(val *string) {
	if err := j.validateSetClusterVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterVersion",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetDefaultIamRoleArn(val *string) {
	if err := j.validateSetDefaultIamRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultIamRoleArn",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetElasticIp(val *string) {
	if err := j.validateSetElasticIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticIp",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetEncrypted(val interface{}) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetEndpoint(val *string) {
	if err := j.validateSetEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpoint",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetEnhancedVpcRouting(val interface{}) {
	if err := j.validateSetEnhancedVpcRoutingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enhancedVpcRouting",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetFinalSnapshotIdentifier(val *string) {
	if err := j.validateSetFinalSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetIamRoles(val *[]*string) {
	if err := j.validateSetIamRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamRoles",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetMaintenanceTrackName(val *string) {
	if err := j.validateSetMaintenanceTrackNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceTrackName",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetManualSnapshotRetentionPeriod(val *float64) {
	if err := j.validateSetManualSnapshotRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manualSnapshotRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetMasterPassword(val *string) {
	if err := j.validateSetMasterPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterPassword",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetMasterUsername(val *string) {
	if err := j.validateSetMasterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetNumberOfNodes(val *float64) {
	if err := j.validateSetNumberOfNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfNodes",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetOwnerAccount(val *string) {
	if err := j.validateSetOwnerAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownerAccount",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetSkipFinalSnapshot(val interface{}) {
	if err := j.validateSetSkipFinalSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalSnapshot",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetSnapshotClusterIdentifier(val *string) {
	if err := j.validateSetSnapshotClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetSnapshotIdentifier(val *string) {
	if err := j.validateSetSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_RedshiftCluster)SetVpcSecurityGroupIds(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroupIds",
		val,
	)
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
func RedshiftCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.redshiftCluster.RedshiftCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedshiftCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.redshiftCluster.RedshiftCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RedshiftCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRedshiftCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.redshiftCluster.RedshiftCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RedshiftCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.redshiftCluster.RedshiftCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RedshiftCluster) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RedshiftCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RedshiftCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RedshiftCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RedshiftCluster) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RedshiftCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RedshiftCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RedshiftCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RedshiftCluster) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RedshiftCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RedshiftCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RedshiftCluster) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RedshiftCluster) PutLogging(value *RedshiftClusterLogging) {
	if err := r.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putLogging",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftCluster) PutSnapshotCopy(value *RedshiftClusterSnapshotCopy) {
	if err := r.validatePutSnapshotCopyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSnapshotCopy",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftCluster) PutTimeouts(value *RedshiftClusterTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetAllowVersionUpgrade() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowVersionUpgrade",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		r,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetAquaConfigurationStatus() {
	_jsii_.InvokeVoid(
		r,
		"resetAquaConfigurationStatus",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetAutomatedSnapshotRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetAutomatedSnapshotRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		r,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetAvailabilityZoneRelocationEnabled() {
	_jsii_.InvokeVoid(
		r,
		"resetAvailabilityZoneRelocationEnabled",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClusterParameterGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterParameterGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClusterPublicKey() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterPublicKey",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClusterRevisionNumber() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterRevisionNumber",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClusterSubnetGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterSubnetGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClusterType() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetClusterVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetClusterVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetDatabaseName() {
	_jsii_.InvokeVoid(
		r,
		"resetDatabaseName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetDefaultIamRoleArn() {
	_jsii_.InvokeVoid(
		r,
		"resetDefaultIamRoleArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetElasticIp() {
	_jsii_.InvokeVoid(
		r,
		"resetElasticIp",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetEncrypted() {
	_jsii_.InvokeVoid(
		r,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetEndpoint() {
	_jsii_.InvokeVoid(
		r,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetEnhancedVpcRouting() {
	_jsii_.InvokeVoid(
		r,
		"resetEnhancedVpcRouting",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetFinalSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetFinalSnapshotIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetIamRoles() {
	_jsii_.InvokeVoid(
		r,
		"resetIamRoles",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetLogging() {
	_jsii_.InvokeVoid(
		r,
		"resetLogging",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetMaintenanceTrackName() {
	_jsii_.InvokeVoid(
		r,
		"resetMaintenanceTrackName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetManualSnapshotRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetManualSnapshotRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetMasterPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetMasterUsername() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUsername",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetNumberOfNodes() {
	_jsii_.InvokeVoid(
		r,
		"resetNumberOfNodes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetOwnerAccount() {
	_jsii_.InvokeVoid(
		r,
		"resetOwnerAccount",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetPort() {
	_jsii_.InvokeVoid(
		r,
		"resetPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		r,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetSkipFinalSnapshot() {
	_jsii_.InvokeVoid(
		r,
		"resetSkipFinalSnapshot",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetSnapshotClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSnapshotClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetSnapshotCopy() {
	_jsii_.InvokeVoid(
		r,
		"resetSnapshotCopy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSnapshotIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		r,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) ResetVpcSecurityGroupIds() {
	_jsii_.InvokeVoid(
		r,
		"resetVpcSecurityGroupIds",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RedshiftCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


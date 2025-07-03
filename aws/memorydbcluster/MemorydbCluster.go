// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorydbcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/memorydbcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/memorydb_cluster aws_memorydb_cluster}.
type MemorydbCluster interface {
	cdktf.TerraformResource
	AclName() *string
	SetAclName(val *string)
	AclNameInput() *string
	Arn() *string
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	AutoMinorVersionUpgradeInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterEndpoint() MemorydbClusterClusterEndpointList
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
	DataTiering() interface{}
	SetDataTiering(val interface{})
	DataTieringInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
	EnginePatchVersion() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	FinalSnapshotName() *string
	SetFinalSnapshotName(val *string)
	FinalSnapshotNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KmsKeyArn() *string
	SetKmsKeyArn(val *string)
	KmsKeyArnInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaintenanceWindow() *string
	SetMaintenanceWindow(val *string)
	MaintenanceWindowInput() *string
	MultiRegionClusterName() *string
	SetMultiRegionClusterName(val *string)
	MultiRegionClusterNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
	// The tree node.
	Node() constructs.Node
	NodeType() *string
	SetNodeType(val *string)
	NodeTypeInput() *string
	NumReplicasPerShard() *float64
	SetNumReplicasPerShard(val *float64)
	NumReplicasPerShardInput() *float64
	NumShards() *float64
	SetNumShards(val *float64)
	NumShardsInput() *float64
	ParameterGroupName() *string
	SetParameterGroupName(val *string)
	ParameterGroupNameInput() *string
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
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	Shards() MemorydbClusterShardsList
	SnapshotArns() *[]*string
	SetSnapshotArns(val *[]*string)
	SnapshotArnsInput() *[]*string
	SnapshotName() *string
	SetSnapshotName(val *string)
	SnapshotNameInput() *string
	SnapshotRetentionLimit() *float64
	SetSnapshotRetentionLimit(val *float64)
	SnapshotRetentionLimitInput() *float64
	SnapshotWindow() *string
	SetSnapshotWindow(val *string)
	SnapshotWindowInput() *string
	SnsTopicArn() *string
	SetSnsTopicArn(val *string)
	SnsTopicArnInput() *string
	SubnetGroupName() *string
	SetSubnetGroupName(val *string)
	SubnetGroupNameInput() *string
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
	Timeouts() MemorydbClusterTimeoutsOutputReference
	TimeoutsInput() interface{}
	TlsEnabled() interface{}
	SetTlsEnabled(val interface{})
	TlsEnabledInput() interface{}
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
	PutTimeouts(value *MemorydbClusterTimeouts)
	ResetAutoMinorVersionUpgrade()
	ResetDataTiering()
	ResetDescription()
	ResetEngine()
	ResetEngineVersion()
	ResetFinalSnapshotName()
	ResetId()
	ResetKmsKeyArn()
	ResetMaintenanceWindow()
	ResetMultiRegionClusterName()
	ResetName()
	ResetNamePrefix()
	ResetNumReplicasPerShard()
	ResetNumShards()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameterGroupName()
	ResetPort()
	ResetRegion()
	ResetSecurityGroupIds()
	ResetSnapshotArns()
	ResetSnapshotName()
	ResetSnapshotRetentionLimit()
	ResetSnapshotWindow()
	ResetSnsTopicArn()
	ResetSubnetGroupName()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetTlsEnabled()
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

// The jsii proxy struct for MemorydbCluster
type jsiiProxy_MemorydbCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MemorydbCluster) AclName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) AclNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aclNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) ClusterEndpoint() MemorydbClusterClusterEndpointList {
	var returns MemorydbClusterClusterEndpointList
	_jsii_.Get(
		j,
		"clusterEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) DataTiering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTiering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) DataTieringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataTieringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) EnginePatchVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enginePatchVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) FinalSnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) FinalSnapshotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) MaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) MaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) MultiRegionClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiRegionClusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) MultiRegionClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"multiRegionClusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) NumReplicasPerShard() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numReplicasPerShard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) NumReplicasPerShardInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numReplicasPerShardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) NumShards() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numShards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) NumShardsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numShardsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) ParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Shards() MemorydbClusterShardsList {
	var returns MemorydbClusterShardsList
	_jsii_.Get(
		j,
		"shards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SnapshotArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SnapshotArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SnapshotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SnapshotRetentionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SnapshotRetentionLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SnapshotWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SnapshotWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SnsTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SnsTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snsTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) SubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) Timeouts() MemorydbClusterTimeoutsOutputReference {
	var returns MemorydbClusterTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) TlsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MemorydbCluster) TlsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tlsEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/memorydb_cluster aws_memorydb_cluster} Resource.
func NewMemorydbCluster(scope constructs.Construct, id *string, config *MemorydbClusterConfig) MemorydbCluster {
	_init_.Initialize()

	if err := validateNewMemorydbClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MemorydbCluster{}

	_jsii_.Create(
		"@cdktf/provider-aws.memorydbCluster.MemorydbCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.2.0/docs/resources/memorydb_cluster aws_memorydb_cluster} Resource.
func NewMemorydbCluster_Override(m MemorydbCluster, scope constructs.Construct, id *string, config *MemorydbClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.memorydbCluster.MemorydbCluster",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetAclName(val *string) {
	if err := j.validateSetAclNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aclName",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetDataTiering(val interface{}) {
	if err := j.validateSetDataTieringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTiering",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetFinalSnapshotName(val *string) {
	if err := j.validateSetFinalSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotName",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetMaintenanceWindow(val *string) {
	if err := j.validateSetMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetMultiRegionClusterName(val *string) {
	if err := j.validateSetMultiRegionClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiRegionClusterName",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetNumReplicasPerShard(val *float64) {
	if err := j.validateSetNumReplicasPerShardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numReplicasPerShard",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetNumShards(val *float64) {
	if err := j.validateSetNumShardsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numShards",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetParameterGroupName(val *string) {
	if err := j.validateSetParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterGroupName",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetSnapshotArns(val *[]*string) {
	if err := j.validateSetSnapshotArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotArns",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetSnapshotName(val *string) {
	if err := j.validateSetSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotName",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetSnapshotRetentionLimit(val *float64) {
	if err := j.validateSetSnapshotRetentionLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotRetentionLimit",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetSnapshotWindow(val *string) {
	if err := j.validateSetSnapshotWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotWindow",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetSnsTopicArn(val *string) {
	if err := j.validateSetSnsTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snsTopicArn",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetSubnetGroupName(val *string) {
	if err := j.validateSetSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetGroupName",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_MemorydbCluster)SetTlsEnabled(val interface{}) {
	if err := j.validateSetTlsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsEnabled",
		val,
	)
}

// Generates CDKTF code for importing a MemorydbCluster resource upon running "cdktf plan <stack-name>".
func MemorydbCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMemorydbCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.memorydbCluster.MemorydbCluster",
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
func MemorydbCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMemorydbCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.memorydbCluster.MemorydbCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MemorydbCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMemorydbCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.memorydbCluster.MemorydbCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MemorydbCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMemorydbCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.memorydbCluster.MemorydbCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MemorydbCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.memorydbCluster.MemorydbCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MemorydbCluster) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MemorydbCluster) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MemorydbCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MemorydbCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MemorydbCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MemorydbCluster) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MemorydbCluster) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MemorydbCluster) PutTimeouts(value *MemorydbClusterTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		m,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetDataTiering() {
	_jsii_.InvokeVoid(
		m,
		"resetDataTiering",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetDescription() {
	_jsii_.InvokeVoid(
		m,
		"resetDescription",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetEngine() {
	_jsii_.InvokeVoid(
		m,
		"resetEngine",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		m,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetFinalSnapshotName() {
	_jsii_.InvokeVoid(
		m,
		"resetFinalSnapshotName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		m,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		m,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetMultiRegionClusterName() {
	_jsii_.InvokeVoid(
		m,
		"resetMultiRegionClusterName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetName() {
	_jsii_.InvokeVoid(
		m,
		"resetName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		m,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetNumReplicasPerShard() {
	_jsii_.InvokeVoid(
		m,
		"resetNumReplicasPerShard",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetNumShards() {
	_jsii_.InvokeVoid(
		m,
		"resetNumShards",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetParameterGroupName() {
	_jsii_.InvokeVoid(
		m,
		"resetParameterGroupName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetPort() {
	_jsii_.InvokeVoid(
		m,
		"resetPort",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetRegion() {
	_jsii_.InvokeVoid(
		m,
		"resetRegion",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		m,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetSnapshotArns() {
	_jsii_.InvokeVoid(
		m,
		"resetSnapshotArns",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetSnapshotName() {
	_jsii_.InvokeVoid(
		m,
		"resetSnapshotName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetSnapshotRetentionLimit() {
	_jsii_.InvokeVoid(
		m,
		"resetSnapshotRetentionLimit",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetSnapshotWindow() {
	_jsii_.InvokeVoid(
		m,
		"resetSnapshotWindow",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetSnsTopicArn() {
	_jsii_.InvokeVoid(
		m,
		"resetSnsTopicArn",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetSubnetGroupName() {
	_jsii_.InvokeVoid(
		m,
		"resetSubnetGroupName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		m,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) ResetTlsEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetTlsEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MemorydbCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MemorydbCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


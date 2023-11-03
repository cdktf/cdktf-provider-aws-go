// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticachecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/elasticachecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/elasticache_cluster aws_elasticache_cluster}.
type ElasticacheCluster interface {
	cdktf.TerraformResource
	ApplyImmediately() interface{}
	SetApplyImmediately(val interface{})
	ApplyImmediatelyInput() interface{}
	Arn() *string
	AutoMinorVersionUpgrade() *string
	SetAutoMinorVersionUpgrade(val *string)
	AutoMinorVersionUpgradeInput() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	AzMode() *string
	SetAzMode(val *string)
	AzModeInput() *string
	CacheNodes() ElasticacheClusterCacheNodesList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterAddress() *string
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ConfigurationEndpoint() *string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpDiscovery() *string
	SetIpDiscovery(val *string)
	IpDiscoveryInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogDeliveryConfiguration() ElasticacheClusterLogDeliveryConfigurationList
	LogDeliveryConfigurationInput() interface{}
	MaintenanceWindow() *string
	SetMaintenanceWindow(val *string)
	MaintenanceWindowInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
	// The tree node.
	Node() constructs.Node
	NodeType() *string
	SetNodeType(val *string)
	NodeTypeInput() *string
	NotificationTopicArn() *string
	SetNotificationTopicArn(val *string)
	NotificationTopicArnInput() *string
	NumCacheNodes() *float64
	SetNumCacheNodes(val *float64)
	NumCacheNodesInput() *float64
	OutpostMode() *string
	SetOutpostMode(val *string)
	OutpostModeInput() *string
	ParameterGroupName() *string
	SetParameterGroupName(val *string)
	ParameterGroupNameInput() *string
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	PreferredAvailabilityZones() *[]*string
	SetPreferredAvailabilityZones(val *[]*string)
	PreferredAvailabilityZonesInput() *[]*string
	PreferredOutpostArn() *string
	SetPreferredOutpostArn(val *string)
	PreferredOutpostArnInput() *string
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
	ReplicationGroupId() *string
	SetReplicationGroupId(val *string)
	ReplicationGroupIdInput() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
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
	TransitEncryptionEnabled() interface{}
	SetTransitEncryptionEnabled(val interface{})
	TransitEncryptionEnabledInput() interface{}
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutLogDeliveryConfiguration(value interface{})
	ResetApplyImmediately()
	ResetAutoMinorVersionUpgrade()
	ResetAvailabilityZone()
	ResetAzMode()
	ResetEngine()
	ResetEngineVersion()
	ResetFinalSnapshotIdentifier()
	ResetId()
	ResetIpDiscovery()
	ResetLogDeliveryConfiguration()
	ResetMaintenanceWindow()
	ResetNetworkType()
	ResetNodeType()
	ResetNotificationTopicArn()
	ResetNumCacheNodes()
	ResetOutpostMode()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameterGroupName()
	ResetPort()
	ResetPreferredAvailabilityZones()
	ResetPreferredOutpostArn()
	ResetReplicationGroupId()
	ResetSecurityGroupIds()
	ResetSnapshotArns()
	ResetSnapshotName()
	ResetSnapshotRetentionLimit()
	ResetSnapshotWindow()
	ResetSubnetGroupName()
	ResetTags()
	ResetTagsAll()
	ResetTransitEncryptionEnabled()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ElasticacheCluster
type jsiiProxy_ElasticacheCluster struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElasticacheCluster) ApplyImmediately() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediately",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) ApplyImmediatelyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applyImmediatelyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) AutoMinorVersionUpgrade() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) AutoMinorVersionUpgradeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) AzMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) AzModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"azModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) CacheNodes() ElasticacheClusterCacheNodesList {
	var returns ElasticacheClusterCacheNodesList
	_jsii_.Get(
		j,
		"cacheNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) ClusterAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) ConfigurationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) EngineVersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) FinalSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) FinalSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"finalSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) IpDiscovery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) IpDiscoveryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipDiscoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) LogDeliveryConfiguration() ElasticacheClusterLogDeliveryConfigurationList {
	var returns ElasticacheClusterLogDeliveryConfigurationList
	_jsii_.Get(
		j,
		"logDeliveryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) LogDeliveryConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logDeliveryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) MaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) MaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) NodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) NodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) NotificationTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) NotificationTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) NumCacheNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCacheNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) NumCacheNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numCacheNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) OutpostMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) OutpostModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outpostModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) ParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) ParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) PreferredAvailabilityZones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredAvailabilityZones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) PreferredAvailabilityZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredAvailabilityZonesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) PreferredOutpostArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredOutpostArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) PreferredOutpostArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredOutpostArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) ReplicationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) ReplicationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SnapshotArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SnapshotArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"snapshotArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SnapshotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SnapshotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SnapshotRetentionLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SnapshotRetentionLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"snapshotRetentionLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SnapshotWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SnapshotWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) SubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) TransitEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheCluster) TransitEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transitEncryptionEnabledInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/elasticache_cluster aws_elasticache_cluster} Resource.
func NewElasticacheCluster(scope constructs.Construct, id *string, config *ElasticacheClusterConfig) ElasticacheCluster {
	_init_.Initialize()

	if err := validateNewElasticacheClusterParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticacheCluster{}

	_jsii_.Create(
		"@cdktf/provider-aws.elasticacheCluster.ElasticacheCluster",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/elasticache_cluster aws_elasticache_cluster} Resource.
func NewElasticacheCluster_Override(e ElasticacheCluster, scope constructs.Construct, id *string, config *ElasticacheClusterConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.elasticacheCluster.ElasticacheCluster",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetApplyImmediately(val interface{}) {
	if err := j.validateSetApplyImmediatelyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applyImmediately",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetAutoMinorVersionUpgrade(val *string) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetAzMode(val *string) {
	if err := j.validateSetAzModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"azMode",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetFinalSnapshotIdentifier(val *string) {
	if err := j.validateSetFinalSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetIpDiscovery(val *string) {
	if err := j.validateSetIpDiscoveryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipDiscovery",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetMaintenanceWindow(val *string) {
	if err := j.validateSetMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetNodeType(val *string) {
	if err := j.validateSetNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeType",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetNotificationTopicArn(val *string) {
	if err := j.validateSetNotificationTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notificationTopicArn",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetNumCacheNodes(val *float64) {
	if err := j.validateSetNumCacheNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numCacheNodes",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetOutpostMode(val *string) {
	if err := j.validateSetOutpostModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outpostMode",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetParameterGroupName(val *string) {
	if err := j.validateSetParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterGroupName",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetPreferredAvailabilityZones(val *[]*string) {
	if err := j.validateSetPreferredAvailabilityZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredAvailabilityZones",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetPreferredOutpostArn(val *string) {
	if err := j.validateSetPreferredOutpostArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredOutpostArn",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetReplicationGroupId(val *string) {
	if err := j.validateSetReplicationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationGroupId",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetSnapshotArns(val *[]*string) {
	if err := j.validateSetSnapshotArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotArns",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetSnapshotName(val *string) {
	if err := j.validateSetSnapshotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotName",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetSnapshotRetentionLimit(val *float64) {
	if err := j.validateSetSnapshotRetentionLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotRetentionLimit",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetSnapshotWindow(val *string) {
	if err := j.validateSetSnapshotWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotWindow",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetSubnetGroupName(val *string) {
	if err := j.validateSetSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetGroupName",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_ElasticacheCluster)SetTransitEncryptionEnabled(val interface{}) {
	if err := j.validateSetTransitEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transitEncryptionEnabled",
		val,
	)
}

// Generates CDKTF code for importing a ElasticacheCluster resource upon running "cdktf plan <stack-name>".
func ElasticacheCluster_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateElasticacheCluster_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elasticacheCluster.ElasticacheCluster",
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
func ElasticacheCluster_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheCluster_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elasticacheCluster.ElasticacheCluster",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticacheCluster_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheCluster_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elasticacheCluster.ElasticacheCluster",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticacheCluster_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheCluster_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.elasticacheCluster.ElasticacheCluster",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElasticacheCluster_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.elasticacheCluster.ElasticacheCluster",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElasticacheCluster) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ElasticacheCluster) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElasticacheCluster) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ElasticacheCluster) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ElasticacheCluster) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElasticacheCluster) PutLogDeliveryConfiguration(value interface{}) {
	if err := e.validatePutLogDeliveryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLogDeliveryConfiguration",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetApplyImmediately() {
	_jsii_.InvokeVoid(
		e,
		"resetApplyImmediately",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		e,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetAzMode() {
	_jsii_.InvokeVoid(
		e,
		"resetAzMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetEngine() {
	_jsii_.InvokeVoid(
		e,
		"resetEngine",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetFinalSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		e,
		"resetFinalSnapshotIdentifier",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetIpDiscovery() {
	_jsii_.InvokeVoid(
		e,
		"resetIpDiscovery",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetLogDeliveryConfiguration() {
	_jsii_.InvokeVoid(
		e,
		"resetLogDeliveryConfiguration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetMaintenanceWindow() {
	_jsii_.InvokeVoid(
		e,
		"resetMaintenanceWindow",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetNetworkType() {
	_jsii_.InvokeVoid(
		e,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetNodeType() {
	_jsii_.InvokeVoid(
		e,
		"resetNodeType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetNotificationTopicArn() {
	_jsii_.InvokeVoid(
		e,
		"resetNotificationTopicArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetNumCacheNodes() {
	_jsii_.InvokeVoid(
		e,
		"resetNumCacheNodes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetOutpostMode() {
	_jsii_.InvokeVoid(
		e,
		"resetOutpostMode",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetParameterGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetParameterGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetPort() {
	_jsii_.InvokeVoid(
		e,
		"resetPort",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetPreferredAvailabilityZones() {
	_jsii_.InvokeVoid(
		e,
		"resetPreferredAvailabilityZones",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetPreferredOutpostArn() {
	_jsii_.InvokeVoid(
		e,
		"resetPreferredOutpostArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetReplicationGroupId() {
	_jsii_.InvokeVoid(
		e,
		"resetReplicationGroupId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetSnapshotArns() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotArns",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetSnapshotName() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetSnapshotRetentionLimit() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotRetentionLimit",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetSnapshotWindow() {
	_jsii_.InvokeVoid(
		e,
		"resetSnapshotWindow",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetSubnetGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetSubnetGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetTagsAll() {
	_jsii_.InvokeVoid(
		e,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) ResetTransitEncryptionEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetTransitEncryptionEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheCluster) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheCluster) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


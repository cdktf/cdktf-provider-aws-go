// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package opsworksinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/opsworksinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/opsworks_instance aws_opsworks_instance}.
type OpsworksInstance interface {
	cdktf.TerraformResource
	AgentVersion() *string
	SetAgentVersion(val *string)
	AgentVersionInput() *string
	AmiId() *string
	SetAmiId(val *string)
	AmiIdInput() *string
	Architecture() *string
	SetArchitecture(val *string)
	ArchitectureInput() *string
	AutoScalingType() *string
	SetAutoScalingType(val *string)
	AutoScalingTypeInput() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	CreatedAt() *string
	SetCreatedAt(val *string)
	CreatedAtInput() *string
	DeleteEbs() interface{}
	SetDeleteEbs(val interface{})
	DeleteEbsInput() interface{}
	DeleteEip() interface{}
	SetDeleteEip(val interface{})
	DeleteEipInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EbsBlockDevice() OpsworksInstanceEbsBlockDeviceList
	EbsBlockDeviceInput() interface{}
	EbsOptimized() interface{}
	SetEbsOptimized(val interface{})
	EbsOptimizedInput() interface{}
	Ec2InstanceId() *string
	EcsClusterArn() *string
	SetEcsClusterArn(val *string)
	EcsClusterArnInput() *string
	ElasticIp() *string
	SetElasticIp(val *string)
	ElasticIpInput() *string
	EphemeralBlockDevice() OpsworksInstanceEphemeralBlockDeviceList
	EphemeralBlockDeviceInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hostname() *string
	SetHostname(val *string)
	HostnameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InfrastructureClass() *string
	SetInfrastructureClass(val *string)
	InfrastructureClassInput() *string
	InstallUpdatesOnBoot() interface{}
	SetInstallUpdatesOnBoot(val interface{})
	InstallUpdatesOnBootInput() interface{}
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	InstanceProfileArnInput() *string
	InstanceType() *string
	SetInstanceType(val *string)
	InstanceTypeInput() *string
	LastServiceErrorId() *string
	LayerIds() *[]*string
	SetLayerIds(val *[]*string)
	LayerIdsInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Os() *string
	SetOs(val *string)
	OsInput() *string
	Platform() *string
	PrivateDns() *string
	PrivateIp() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PublicDns() *string
	PublicIp() *string
	// Experimental.
	RawOverrides() interface{}
	RegisteredBy() *string
	ReportedAgentVersion() *string
	ReportedOsFamily() *string
	ReportedOsName() *string
	ReportedOsVersion() *string
	RootBlockDevice() OpsworksInstanceRootBlockDeviceList
	RootBlockDeviceInput() interface{}
	RootDeviceType() *string
	SetRootDeviceType(val *string)
	RootDeviceTypeInput() *string
	RootDeviceVolumeId() *string
	SecurityGroupIds() *[]*string
	SetSecurityGroupIds(val *[]*string)
	SecurityGroupIdsInput() *[]*string
	SshHostDsaKeyFingerprint() *string
	SshHostRsaKeyFingerprint() *string
	SshKeyName() *string
	SetSshKeyName(val *string)
	SshKeyNameInput() *string
	StackId() *string
	SetStackId(val *string)
	StackIdInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	Status() *string
	SetStatus(val *string)
	StatusInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	Tenancy() *string
	SetTenancy(val *string)
	TenancyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() OpsworksInstanceTimeoutsOutputReference
	TimeoutsInput() interface{}
	VirtualizationType() *string
	SetVirtualizationType(val *string)
	VirtualizationTypeInput() *string
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
	PutEbsBlockDevice(value interface{})
	PutEphemeralBlockDevice(value interface{})
	PutRootBlockDevice(value interface{})
	PutTimeouts(value *OpsworksInstanceTimeouts)
	ResetAgentVersion()
	ResetAmiId()
	ResetArchitecture()
	ResetAutoScalingType()
	ResetAvailabilityZone()
	ResetCreatedAt()
	ResetDeleteEbs()
	ResetDeleteEip()
	ResetEbsBlockDevice()
	ResetEbsOptimized()
	ResetEcsClusterArn()
	ResetElasticIp()
	ResetEphemeralBlockDevice()
	ResetHostname()
	ResetId()
	ResetInfrastructureClass()
	ResetInstallUpdatesOnBoot()
	ResetInstanceProfileArn()
	ResetInstanceType()
	ResetOs()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRootBlockDevice()
	ResetRootDeviceType()
	ResetSecurityGroupIds()
	ResetSshKeyName()
	ResetState()
	ResetStatus()
	ResetSubnetId()
	ResetTenancy()
	ResetTimeouts()
	ResetVirtualizationType()
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

// The jsii proxy struct for OpsworksInstance
type jsiiProxy_OpsworksInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OpsworksInstance) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AgentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AmiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AutoScalingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AutoScalingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) CreatedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) CreatedAtInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdAtInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) DeleteEbs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteEbs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) DeleteEbsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteEbsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) DeleteEip() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteEip",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) DeleteEipInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteEipInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EbsBlockDevice() OpsworksInstanceEbsBlockDeviceList {
	var returns OpsworksInstanceEbsBlockDeviceList
	_jsii_.Get(
		j,
		"ebsBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EbsBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EbsOptimized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EbsOptimizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ebsOptimizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Ec2InstanceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ec2InstanceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EcsClusterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsClusterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EcsClusterArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecsClusterArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ElasticIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ElasticIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"elasticIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EphemeralBlockDevice() OpsworksInstanceEphemeralBlockDeviceList {
	var returns OpsworksInstanceEphemeralBlockDeviceList
	_jsii_.Get(
		j,
		"ephemeralBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) EphemeralBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ephemeralBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Hostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) HostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InfrastructureClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InfrastructureClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"infrastructureClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstallUpdatesOnBoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstallUpdatesOnBootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installUpdatesOnBootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) InstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) LastServiceErrorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastServiceErrorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) LayerIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layerIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) LayerIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"layerIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Os() *string {
	var returns *string
	_jsii_.Get(
		j,
		"os",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) OsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PrivateDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PrivateIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PublicDns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicDns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) PublicIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RegisteredBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registeredBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedAgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedAgentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedOsFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedOsFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedOsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedOsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) ReportedOsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportedOsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootBlockDevice() OpsworksInstanceRootBlockDeviceList {
	var returns OpsworksInstanceRootBlockDeviceList
	_jsii_.Get(
		j,
		"rootBlockDevice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootBlockDeviceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootBlockDeviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootDeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootDeviceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) RootDeviceVolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootDeviceVolumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SshHostDsaKeyFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshHostDsaKeyFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SshHostRsaKeyFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshHostRsaKeyFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SshKeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKeyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SshKeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKeyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) StackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) StackIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stackIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) Timeouts() OpsworksInstanceTimeoutsOutputReference {
	var returns OpsworksInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) VirtualizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsworksInstance) VirtualizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualizationTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/opsworks_instance aws_opsworks_instance} Resource.
func NewOpsworksInstance(scope constructs.Construct, id *string, config *OpsworksInstanceConfig) OpsworksInstance {
	_init_.Initialize()

	if err := validateNewOpsworksInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsworksInstance{}

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksInstance.OpsworksInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/opsworks_instance aws_opsworks_instance} Resource.
func NewOpsworksInstance_Override(o OpsworksInstance, scope constructs.Construct, id *string, config *OpsworksInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opsworksInstance.OpsworksInstance",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetAgentVersion(val *string) {
	if err := j.validateSetAgentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetAmiId(val *string) {
	if err := j.validateSetAmiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amiId",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetAutoScalingType(val *string) {
	if err := j.validateSetAutoScalingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingType",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetCreatedAt(val *string) {
	if err := j.validateSetCreatedAtParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createdAt",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetDeleteEbs(val interface{}) {
	if err := j.validateSetDeleteEbsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteEbs",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetDeleteEip(val interface{}) {
	if err := j.validateSetDeleteEipParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteEip",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetEbsOptimized(val interface{}) {
	if err := j.validateSetEbsOptimizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ebsOptimized",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetEcsClusterArn(val *string) {
	if err := j.validateSetEcsClusterArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecsClusterArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetElasticIp(val *string) {
	if err := j.validateSetElasticIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticIp",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetHostname(val *string) {
	if err := j.validateSetHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostname",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetInfrastructureClass(val *string) {
	if err := j.validateSetInfrastructureClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"infrastructureClass",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetInstallUpdatesOnBoot(val interface{}) {
	if err := j.validateSetInstallUpdatesOnBootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"installUpdatesOnBoot",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetInstanceProfileArn(val *string) {
	if err := j.validateSetInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetInstanceType(val *string) {
	if err := j.validateSetInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceType",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetLayerIds(val *[]*string) {
	if err := j.validateSetLayerIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"layerIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetOs(val *string) {
	if err := j.validateSetOsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"os",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetRootDeviceType(val *string) {
	if err := j.validateSetRootDeviceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootDeviceType",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetSecurityGroupIds(val *[]*string) {
	if err := j.validateSetSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupIds",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetSshKeyName(val *string) {
	if err := j.validateSetSshKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKeyName",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetStackId(val *string) {
	if err := j.validateSetStackIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stackId",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_OpsworksInstance)SetVirtualizationType(val *string) {
	if err := j.validateSetVirtualizationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualizationType",
		val,
	)
}

// Generates CDKTF code for importing a OpsworksInstance resource upon running "cdktf plan <stack-name>".
func OpsworksInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOpsworksInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksInstance.OpsworksInstance",
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
func OpsworksInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksInstance.OpsworksInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsworksInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksInstance.OpsworksInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OpsworksInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOpsworksInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.opsworksInstance.OpsworksInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OpsworksInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.opsworksInstance.OpsworksInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OpsworksInstance) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OpsworksInstance) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OpsworksInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OpsworksInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OpsworksInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OpsworksInstance) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OpsworksInstance) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OpsworksInstance) PutEbsBlockDevice(value interface{}) {
	if err := o.validatePutEbsBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEbsBlockDevice",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksInstance) PutEphemeralBlockDevice(value interface{}) {
	if err := o.validatePutEphemeralBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEphemeralBlockDevice",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksInstance) PutRootBlockDevice(value interface{}) {
	if err := o.validatePutRootBlockDeviceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRootBlockDevice",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksInstance) PutTimeouts(value *OpsworksInstanceTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetAgentVersion() {
	_jsii_.InvokeVoid(
		o,
		"resetAgentVersion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetAmiId() {
	_jsii_.InvokeVoid(
		o,
		"resetAmiId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetArchitecture() {
	_jsii_.InvokeVoid(
		o,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetAutoScalingType() {
	_jsii_.InvokeVoid(
		o,
		"resetAutoScalingType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		o,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetCreatedAt() {
	_jsii_.InvokeVoid(
		o,
		"resetCreatedAt",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetDeleteEbs() {
	_jsii_.InvokeVoid(
		o,
		"resetDeleteEbs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetDeleteEip() {
	_jsii_.InvokeVoid(
		o,
		"resetDeleteEip",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetEbsBlockDevice() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsBlockDevice",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetEbsOptimized() {
	_jsii_.InvokeVoid(
		o,
		"resetEbsOptimized",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetEcsClusterArn() {
	_jsii_.InvokeVoid(
		o,
		"resetEcsClusterArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetElasticIp() {
	_jsii_.InvokeVoid(
		o,
		"resetElasticIp",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetEphemeralBlockDevice() {
	_jsii_.InvokeVoid(
		o,
		"resetEphemeralBlockDevice",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetHostname() {
	_jsii_.InvokeVoid(
		o,
		"resetHostname",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetInfrastructureClass() {
	_jsii_.InvokeVoid(
		o,
		"resetInfrastructureClass",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetInstallUpdatesOnBoot() {
	_jsii_.InvokeVoid(
		o,
		"resetInstallUpdatesOnBoot",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetInstanceType() {
	_jsii_.InvokeVoid(
		o,
		"resetInstanceType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetOs() {
	_jsii_.InvokeVoid(
		o,
		"resetOs",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetRootBlockDevice() {
	_jsii_.InvokeVoid(
		o,
		"resetRootBlockDevice",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetRootDeviceType() {
	_jsii_.InvokeVoid(
		o,
		"resetRootDeviceType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetSecurityGroupIds() {
	_jsii_.InvokeVoid(
		o,
		"resetSecurityGroupIds",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetSshKeyName() {
	_jsii_.InvokeVoid(
		o,
		"resetSshKeyName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetState() {
	_jsii_.InvokeVoid(
		o,
		"resetState",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetStatus() {
	_jsii_.InvokeVoid(
		o,
		"resetStatus",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetSubnetId() {
	_jsii_.InvokeVoid(
		o,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetTenancy() {
	_jsii_.InvokeVoid(
		o,
		"resetTenancy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) ResetVirtualizationType() {
	_jsii_.InvokeVoid(
		o,
		"resetVirtualizationType",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OpsworksInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpsworksInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


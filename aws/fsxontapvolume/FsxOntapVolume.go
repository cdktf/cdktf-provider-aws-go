// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxontapvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/fsxontapvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_ontap_volume aws_fsx_ontap_volume}.
type FsxOntapVolume interface {
	cdktf.TerraformResource
	AggregateConfiguration() FsxOntapVolumeAggregateConfigurationOutputReference
	AggregateConfigurationInput() *FsxOntapVolumeAggregateConfiguration
	Arn() *string
	BypassSnaplockEnterpriseRetention() interface{}
	SetBypassSnaplockEnterpriseRetention(val interface{})
	BypassSnaplockEnterpriseRetentionInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToBackups() interface{}
	SetCopyTagsToBackups(val interface{})
	CopyTagsToBackupsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FileSystemId() *string
	FinalBackupTags() *map[string]*string
	SetFinalBackupTags(val *map[string]*string)
	FinalBackupTagsInput() *map[string]*string
	FlexcacheEndpointType() *string
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
	JunctionPath() *string
	SetJunctionPath(val *string)
	JunctionPathInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OntapVolumeType() *string
	SetOntapVolumeType(val *string)
	OntapVolumeTypeInput() *string
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
	SecurityStyle() *string
	SetSecurityStyle(val *string)
	SecurityStyleInput() *string
	SizeInBytes() *string
	SetSizeInBytes(val *string)
	SizeInBytesInput() *string
	SizeInMegabytes() *float64
	SetSizeInMegabytes(val *float64)
	SizeInMegabytesInput() *float64
	SkipFinalBackup() interface{}
	SetSkipFinalBackup(val interface{})
	SkipFinalBackupInput() interface{}
	SnaplockConfiguration() FsxOntapVolumeSnaplockConfigurationOutputReference
	SnaplockConfigurationInput() *FsxOntapVolumeSnaplockConfiguration
	SnapshotPolicy() *string
	SetSnapshotPolicy(val *string)
	SnapshotPolicyInput() *string
	StorageEfficiencyEnabled() interface{}
	SetStorageEfficiencyEnabled(val interface{})
	StorageEfficiencyEnabledInput() interface{}
	StorageVirtualMachineId() *string
	SetStorageVirtualMachineId(val *string)
	StorageVirtualMachineIdInput() *string
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
	TieringPolicy() FsxOntapVolumeTieringPolicyOutputReference
	TieringPolicyInput() *FsxOntapVolumeTieringPolicy
	Timeouts() FsxOntapVolumeTimeoutsOutputReference
	TimeoutsInput() interface{}
	Uuid() *string
	VolumeStyle() *string
	SetVolumeStyle(val *string)
	VolumeStyleInput() *string
	VolumeType() *string
	SetVolumeType(val *string)
	VolumeTypeInput() *string
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
	PutAggregateConfiguration(value *FsxOntapVolumeAggregateConfiguration)
	PutSnaplockConfiguration(value *FsxOntapVolumeSnaplockConfiguration)
	PutTieringPolicy(value *FsxOntapVolumeTieringPolicy)
	PutTimeouts(value *FsxOntapVolumeTimeouts)
	ResetAggregateConfiguration()
	ResetBypassSnaplockEnterpriseRetention()
	ResetCopyTagsToBackups()
	ResetFinalBackupTags()
	ResetId()
	ResetJunctionPath()
	ResetOntapVolumeType()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSecurityStyle()
	ResetSizeInBytes()
	ResetSizeInMegabytes()
	ResetSkipFinalBackup()
	ResetSnaplockConfiguration()
	ResetSnapshotPolicy()
	ResetStorageEfficiencyEnabled()
	ResetTags()
	ResetTagsAll()
	ResetTieringPolicy()
	ResetTimeouts()
	ResetVolumeStyle()
	ResetVolumeType()
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

// The jsii proxy struct for FsxOntapVolume
type jsiiProxy_FsxOntapVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxOntapVolume) AggregateConfiguration() FsxOntapVolumeAggregateConfigurationOutputReference {
	var returns FsxOntapVolumeAggregateConfigurationOutputReference
	_jsii_.Get(
		j,
		"aggregateConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) AggregateConfigurationInput() *FsxOntapVolumeAggregateConfiguration {
	var returns *FsxOntapVolumeAggregateConfiguration
	_jsii_.Get(
		j,
		"aggregateConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) BypassSnaplockEnterpriseRetention() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassSnaplockEnterpriseRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) BypassSnaplockEnterpriseRetentionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassSnaplockEnterpriseRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) CopyTagsToBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) CopyTagsToBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) FileSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) FinalBackupTags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) FinalBackupTagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"finalBackupTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) FlexcacheEndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"flexcacheEndpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) JunctionPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"junctionPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) JunctionPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"junctionPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) OntapVolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ontapVolumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) OntapVolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ontapVolumeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SecurityStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SecurityStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SizeInBytes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SizeInBytesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInBytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SizeInMegabytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInMegabytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SizeInMegabytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sizeInMegabytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SkipFinalBackup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SkipFinalBackupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipFinalBackupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SnaplockConfiguration() FsxOntapVolumeSnaplockConfigurationOutputReference {
	var returns FsxOntapVolumeSnaplockConfigurationOutputReference
	_jsii_.Get(
		j,
		"snaplockConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SnaplockConfigurationInput() *FsxOntapVolumeSnaplockConfiguration {
	var returns *FsxOntapVolumeSnaplockConfiguration
	_jsii_.Get(
		j,
		"snaplockConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SnapshotPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) SnapshotPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) StorageEfficiencyEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEfficiencyEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) StorageEfficiencyEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEfficiencyEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) StorageVirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageVirtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) StorageVirtualMachineIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageVirtualMachineIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TieringPolicy() FsxOntapVolumeTieringPolicyOutputReference {
	var returns FsxOntapVolumeTieringPolicyOutputReference
	_jsii_.Get(
		j,
		"tieringPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TieringPolicyInput() *FsxOntapVolumeTieringPolicy {
	var returns *FsxOntapVolumeTieringPolicy
	_jsii_.Get(
		j,
		"tieringPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Timeouts() FsxOntapVolumeTimeoutsOutputReference {
	var returns FsxOntapVolumeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) VolumeStyle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeStyle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) VolumeStyleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeStyleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOntapVolume) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_ontap_volume aws_fsx_ontap_volume} Resource.
func NewFsxOntapVolume(scope constructs.Construct, id *string, config *FsxOntapVolumeConfig) FsxOntapVolume {
	_init_.Initialize()

	if err := validateNewFsxOntapVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxOntapVolume{}

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOntapVolume.FsxOntapVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fsx_ontap_volume aws_fsx_ontap_volume} Resource.
func NewFsxOntapVolume_Override(f FsxOntapVolume, scope constructs.Construct, id *string, config *FsxOntapVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOntapVolume.FsxOntapVolume",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetBypassSnaplockEnterpriseRetention(val interface{}) {
	if err := j.validateSetBypassSnaplockEnterpriseRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bypassSnaplockEnterpriseRetention",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetCopyTagsToBackups(val interface{}) {
	if err := j.validateSetCopyTagsToBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToBackups",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetFinalBackupTags(val *map[string]*string) {
	if err := j.validateSetFinalBackupTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"finalBackupTags",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetJunctionPath(val *string) {
	if err := j.validateSetJunctionPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"junctionPath",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetOntapVolumeType(val *string) {
	if err := j.validateSetOntapVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ontapVolumeType",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetSecurityStyle(val *string) {
	if err := j.validateSetSecurityStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityStyle",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetSizeInBytes(val *string) {
	if err := j.validateSetSizeInBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInBytes",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetSizeInMegabytes(val *float64) {
	if err := j.validateSetSizeInMegabytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sizeInMegabytes",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetSkipFinalBackup(val interface{}) {
	if err := j.validateSetSkipFinalBackupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipFinalBackup",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetSnapshotPolicy(val *string) {
	if err := j.validateSetSnapshotPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotPolicy",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetStorageEfficiencyEnabled(val interface{}) {
	if err := j.validateSetStorageEfficiencyEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEfficiencyEnabled",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetStorageVirtualMachineId(val *string) {
	if err := j.validateSetStorageVirtualMachineIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageVirtualMachineId",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetVolumeStyle(val *string) {
	if err := j.validateSetVolumeStyleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeStyle",
		val,
	)
}

func (j *jsiiProxy_FsxOntapVolume)SetVolumeType(val *string) {
	if err := j.validateSetVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

// Generates CDKTF code for importing a FsxOntapVolume resource upon running "cdktf plan <stack-name>".
func FsxOntapVolume_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFsxOntapVolume_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxOntapVolume.FsxOntapVolume",
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
func FsxOntapVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxOntapVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxOntapVolume.FsxOntapVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FsxOntapVolume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxOntapVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxOntapVolume.FsxOntapVolume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FsxOntapVolume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxOntapVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxOntapVolume.FsxOntapVolume",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxOntapVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.fsxOntapVolume.FsxOntapVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FsxOntapVolume) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FsxOntapVolume) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FsxOntapVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FsxOntapVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) MoveFromId(id *string) {
	if err := f.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveFromId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FsxOntapVolume) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FsxOntapVolume) MoveToId(id *string) {
	if err := f.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveToId",
		[]interface{}{id},
	)
}

func (f *jsiiProxy_FsxOntapVolume) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxOntapVolume) PutAggregateConfiguration(value *FsxOntapVolumeAggregateConfiguration) {
	if err := f.validatePutAggregateConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putAggregateConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapVolume) PutSnaplockConfiguration(value *FsxOntapVolumeSnaplockConfiguration) {
	if err := f.validatePutSnaplockConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putSnaplockConfiguration",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapVolume) PutTieringPolicy(value *FsxOntapVolumeTieringPolicy) {
	if err := f.validatePutTieringPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTieringPolicy",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapVolume) PutTimeouts(value *FsxOntapVolumeTimeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetAggregateConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetAggregateConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetBypassSnaplockEnterpriseRetention() {
	_jsii_.InvokeVoid(
		f,
		"resetBypassSnaplockEnterpriseRetention",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetCopyTagsToBackups() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToBackups",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetFinalBackupTags() {
	_jsii_.InvokeVoid(
		f,
		"resetFinalBackupTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetJunctionPath() {
	_jsii_.InvokeVoid(
		f,
		"resetJunctionPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetOntapVolumeType() {
	_jsii_.InvokeVoid(
		f,
		"resetOntapVolumeType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetRegion() {
	_jsii_.InvokeVoid(
		f,
		"resetRegion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetSecurityStyle() {
	_jsii_.InvokeVoid(
		f,
		"resetSecurityStyle",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetSizeInBytes() {
	_jsii_.InvokeVoid(
		f,
		"resetSizeInBytes",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetSizeInMegabytes() {
	_jsii_.InvokeVoid(
		f,
		"resetSizeInMegabytes",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetSkipFinalBackup() {
	_jsii_.InvokeVoid(
		f,
		"resetSkipFinalBackup",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetSnaplockConfiguration() {
	_jsii_.InvokeVoid(
		f,
		"resetSnaplockConfiguration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetSnapshotPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetSnapshotPolicy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetStorageEfficiencyEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageEfficiencyEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetTieringPolicy() {
	_jsii_.InvokeVoid(
		f,
		"resetTieringPolicy",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetVolumeStyle() {
	_jsii_.InvokeVoid(
		f,
		"resetVolumeStyle",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) ResetVolumeType() {
	_jsii_.InvokeVoid(
		f,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOntapVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOntapVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


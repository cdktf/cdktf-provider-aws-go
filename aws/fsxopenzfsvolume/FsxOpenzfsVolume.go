// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fsxopenzfsvolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/fsxopenzfsvolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/fsx_openzfs_volume aws_fsx_openzfs_volume}.
type FsxOpenzfsVolume interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToSnapshots() interface{}
	SetCopyTagsToSnapshots(val interface{})
	CopyTagsToSnapshotsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DataCompressionType() *string
	SetDataCompressionType(val *string)
	DataCompressionTypeInput() *string
	DeleteVolumeOptions() *[]*string
	SetDeleteVolumeOptions(val *[]*string)
	DeleteVolumeOptionsInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	NfsExports() FsxOpenzfsVolumeNfsExportsOutputReference
	NfsExportsInput() *FsxOpenzfsVolumeNfsExports
	// The tree node.
	Node() constructs.Node
	OriginSnapshot() FsxOpenzfsVolumeOriginSnapshotOutputReference
	OriginSnapshotInput() *FsxOpenzfsVolumeOriginSnapshot
	ParentVolumeId() *string
	SetParentVolumeId(val *string)
	ParentVolumeIdInput() *string
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
	ReadOnly() interface{}
	SetReadOnly(val interface{})
	ReadOnlyInput() interface{}
	RecordSizeKib() *float64
	SetRecordSizeKib(val *float64)
	RecordSizeKibInput() *float64
	StorageCapacityQuotaGib() *float64
	SetStorageCapacityQuotaGib(val *float64)
	StorageCapacityQuotaGibInput() *float64
	StorageCapacityReservationGib() *float64
	SetStorageCapacityReservationGib(val *float64)
	StorageCapacityReservationGibInput() *float64
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
	Timeouts() FsxOpenzfsVolumeTimeoutsOutputReference
	TimeoutsInput() interface{}
	UserAndGroupQuotas() FsxOpenzfsVolumeUserAndGroupQuotasList
	UserAndGroupQuotasInput() interface{}
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutNfsExports(value *FsxOpenzfsVolumeNfsExports)
	PutOriginSnapshot(value *FsxOpenzfsVolumeOriginSnapshot)
	PutTimeouts(value *FsxOpenzfsVolumeTimeouts)
	PutUserAndGroupQuotas(value interface{})
	ResetCopyTagsToSnapshots()
	ResetDataCompressionType()
	ResetDeleteVolumeOptions()
	ResetId()
	ResetNfsExports()
	ResetOriginSnapshot()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReadOnly()
	ResetRecordSizeKib()
	ResetStorageCapacityQuotaGib()
	ResetStorageCapacityReservationGib()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetUserAndGroupQuotas()
	ResetVolumeType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for FsxOpenzfsVolume
type jsiiProxy_FsxOpenzfsVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_FsxOpenzfsVolume) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) CopyTagsToSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) CopyTagsToSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) DeleteVolumeOptions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deleteVolumeOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) DeleteVolumeOptionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"deleteVolumeOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) NfsExports() FsxOpenzfsVolumeNfsExportsOutputReference {
	var returns FsxOpenzfsVolumeNfsExportsOutputReference
	_jsii_.Get(
		j,
		"nfsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) NfsExportsInput() *FsxOpenzfsVolumeNfsExports {
	var returns *FsxOpenzfsVolumeNfsExports
	_jsii_.Get(
		j,
		"nfsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) OriginSnapshot() FsxOpenzfsVolumeOriginSnapshotOutputReference {
	var returns FsxOpenzfsVolumeOriginSnapshotOutputReference
	_jsii_.Get(
		j,
		"originSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) OriginSnapshotInput() *FsxOpenzfsVolumeOriginSnapshot {
	var returns *FsxOpenzfsVolumeOriginSnapshot
	_jsii_.Get(
		j,
		"originSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) ParentVolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentVolumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) ParentVolumeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentVolumeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) RecordSizeKib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) RecordSizeKibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) StorageCapacityQuotaGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityQuotaGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) StorageCapacityQuotaGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityQuotaGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) StorageCapacityReservationGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityReservationGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) StorageCapacityReservationGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageCapacityReservationGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) Timeouts() FsxOpenzfsVolumeTimeoutsOutputReference {
	var returns FsxOpenzfsVolumeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) UserAndGroupQuotas() FsxOpenzfsVolumeUserAndGroupQuotasList {
	var returns FsxOpenzfsVolumeUserAndGroupQuotasList
	_jsii_.Get(
		j,
		"userAndGroupQuotas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) UserAndGroupQuotasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAndGroupQuotasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FsxOpenzfsVolume) VolumeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/fsx_openzfs_volume aws_fsx_openzfs_volume} Resource.
func NewFsxOpenzfsVolume(scope constructs.Construct, id *string, config *FsxOpenzfsVolumeConfig) FsxOpenzfsVolume {
	_init_.Initialize()

	if err := validateNewFsxOpenzfsVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_FsxOpenzfsVolume{}

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOpenzfsVolume.FsxOpenzfsVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.0/docs/resources/fsx_openzfs_volume aws_fsx_openzfs_volume} Resource.
func NewFsxOpenzfsVolume_Override(f FsxOpenzfsVolume, scope constructs.Construct, id *string, config *FsxOpenzfsVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.fsxOpenzfsVolume.FsxOpenzfsVolume",
		[]interface{}{scope, id, config},
		f,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetCopyTagsToSnapshots(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshots",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetDataCompressionType(val *string) {
	if err := j.validateSetDataCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetDeleteVolumeOptions(val *[]*string) {
	if err := j.validateSetDeleteVolumeOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteVolumeOptions",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetParentVolumeId(val *string) {
	if err := j.validateSetParentVolumeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentVolumeId",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetRecordSizeKib(val *float64) {
	if err := j.validateSetRecordSizeKibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordSizeKib",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetStorageCapacityQuotaGib(val *float64) {
	if err := j.validateSetStorageCapacityQuotaGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacityQuotaGib",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetStorageCapacityReservationGib(val *float64) {
	if err := j.validateSetStorageCapacityReservationGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageCapacityReservationGib",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_FsxOpenzfsVolume)SetVolumeType(val *string) {
	if err := j.validateSetVolumeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeType",
		val,
	)
}

// Generates CDKTF code for importing a FsxOpenzfsVolume resource upon running "cdktf plan <stack-name>".
func FsxOpenzfsVolume_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateFsxOpenzfsVolume_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxOpenzfsVolume.FsxOpenzfsVolume",
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
func FsxOpenzfsVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxOpenzfsVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxOpenzfsVolume.FsxOpenzfsVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FsxOpenzfsVolume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxOpenzfsVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxOpenzfsVolume.FsxOpenzfsVolume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func FsxOpenzfsVolume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateFsxOpenzfsVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.fsxOpenzfsVolume.FsxOpenzfsVolume",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func FsxOpenzfsVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.fsxOpenzfsVolume.FsxOpenzfsVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FsxOpenzfsVolume) AddMoveTarget(moveTarget *string) {
	if err := f.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) AddOverride(path *string, value interface{}) {
	if err := f.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FsxOpenzfsVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FsxOpenzfsVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FsxOpenzfsVolume) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FsxOpenzfsVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FsxOpenzfsVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FsxOpenzfsVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FsxOpenzfsVolume) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FsxOpenzfsVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FsxOpenzfsVolume) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := f.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FsxOpenzfsVolume) MoveTo(moveTarget *string, index interface{}) {
	if err := f.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) OverrideLogicalId(newLogicalId *string) {
	if err := f.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) PutNfsExports(value *FsxOpenzfsVolumeNfsExports) {
	if err := f.validatePutNfsExportsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putNfsExports",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) PutOriginSnapshot(value *FsxOpenzfsVolumeOriginSnapshot) {
	if err := f.validatePutOriginSnapshotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putOriginSnapshot",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) PutTimeouts(value *FsxOpenzfsVolumeTimeouts) {
	if err := f.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) PutUserAndGroupQuotas(value interface{}) {
	if err := f.validatePutUserAndGroupQuotasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putUserAndGroupQuotas",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetCopyTagsToSnapshots() {
	_jsii_.InvokeVoid(
		f,
		"resetCopyTagsToSnapshots",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		f,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetDeleteVolumeOptions() {
	_jsii_.InvokeVoid(
		f,
		"resetDeleteVolumeOptions",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetId() {
	_jsii_.InvokeVoid(
		f,
		"resetId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetNfsExports() {
	_jsii_.InvokeVoid(
		f,
		"resetNfsExports",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetOriginSnapshot() {
	_jsii_.InvokeVoid(
		f,
		"resetOriginSnapshot",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		f,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetReadOnly() {
	_jsii_.InvokeVoid(
		f,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetRecordSizeKib() {
	_jsii_.InvokeVoid(
		f,
		"resetRecordSizeKib",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetStorageCapacityQuotaGib() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageCapacityQuotaGib",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetStorageCapacityReservationGib() {
	_jsii_.InvokeVoid(
		f,
		"resetStorageCapacityReservationGib",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetTags() {
	_jsii_.InvokeVoid(
		f,
		"resetTags",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetTagsAll() {
	_jsii_.InvokeVoid(
		f,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetTimeouts() {
	_jsii_.InvokeVoid(
		f,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetUserAndGroupQuotas() {
	_jsii_.InvokeVoid(
		f,
		"resetUserAndGroupQuotas",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) ResetVolumeType() {
	_jsii_.InvokeVoid(
		f,
		"resetVolumeType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FsxOpenzfsVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FsxOpenzfsVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		f,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagegatewaycachediscsivolume

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/storagegatewaycachediscsivolume/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/storagegateway_cached_iscsi_volume aws_storagegateway_cached_iscsi_volume}.
type StoragegatewayCachedIscsiVolume interface {
	cdktf.TerraformResource
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChapEnabled() cdktf.IResolvable
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GatewayArn() *string
	SetGatewayArn(val *string)
	GatewayArnInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KmsEncrypted() interface{}
	SetKmsEncrypted(val interface{})
	KmsEncryptedInput() interface{}
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LunNumber() *float64
	NetworkInterfaceId() *string
	SetNetworkInterfaceId(val *string)
	NetworkInterfaceIdInput() *string
	NetworkInterfacePort() *float64
	// The tree node.
	Node() constructs.Node
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
	SnapshotId() *string
	SetSnapshotId(val *string)
	SnapshotIdInput() *string
	SourceVolumeArn() *string
	SetSourceVolumeArn(val *string)
	SourceVolumeArnInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TargetArn() *string
	TargetName() *string
	SetTargetName(val *string)
	TargetNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VolumeArn() *string
	VolumeId() *string
	VolumeSizeInBytes() *float64
	SetVolumeSizeInBytes(val *float64)
	VolumeSizeInBytesInput() *float64
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
	ResetId()
	ResetKmsEncrypted()
	ResetKmsKey()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSnapshotId()
	ResetSourceVolumeArn()
	ResetTags()
	ResetTagsAll()
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

// The jsii proxy struct for StoragegatewayCachedIscsiVolume
type jsiiProxy_StoragegatewayCachedIscsiVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) ChapEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"chapEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) KmsEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) KmsEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) LunNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lunNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) NetworkInterfacePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkInterfacePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SourceVolumeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVolumeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) SourceVolumeArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceVolumeArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TargetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TargetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) VolumeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) VolumeSizeInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume) VolumeSizeInBytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInBytesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/storagegateway_cached_iscsi_volume aws_storagegateway_cached_iscsi_volume} Resource.
func NewStoragegatewayCachedIscsiVolume(scope constructs.Construct, id *string, config *StoragegatewayCachedIscsiVolumeConfig) StoragegatewayCachedIscsiVolume {
	_init_.Initialize()

	if err := validateNewStoragegatewayCachedIscsiVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StoragegatewayCachedIscsiVolume{}

	_jsii_.Create(
		"@cdktf/provider-aws.storagegatewayCachedIscsiVolume.StoragegatewayCachedIscsiVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.53.0/docs/resources/storagegateway_cached_iscsi_volume aws_storagegateway_cached_iscsi_volume} Resource.
func NewStoragegatewayCachedIscsiVolume_Override(s StoragegatewayCachedIscsiVolume, scope constructs.Construct, id *string, config *StoragegatewayCachedIscsiVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.storagegatewayCachedIscsiVolume.StoragegatewayCachedIscsiVolume",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetGatewayArn(val *string) {
	if err := j.validateSetGatewayArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetKmsEncrypted(val interface{}) {
	if err := j.validateSetKmsEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsEncrypted",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetNetworkInterfaceId(val *string) {
	if err := j.validateSetNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetSourceVolumeArn(val *string) {
	if err := j.validateSetSourceVolumeArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceVolumeArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetTargetName(val *string) {
	if err := j.validateSetTargetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetName",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayCachedIscsiVolume)SetVolumeSizeInBytes(val *float64) {
	if err := j.validateSetVolumeSizeInBytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSizeInBytes",
		val,
	)
}

// Generates CDKTF code for importing a StoragegatewayCachedIscsiVolume resource upon running "cdktf plan <stack-name>".
func StoragegatewayCachedIscsiVolume_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStoragegatewayCachedIscsiVolume_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewayCachedIscsiVolume.StoragegatewayCachedIscsiVolume",
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
func StoragegatewayCachedIscsiVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewayCachedIscsiVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewayCachedIscsiVolume.StoragegatewayCachedIscsiVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StoragegatewayCachedIscsiVolume_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewayCachedIscsiVolume_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewayCachedIscsiVolume.StoragegatewayCachedIscsiVolume",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StoragegatewayCachedIscsiVolume_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewayCachedIscsiVolume_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegatewayCachedIscsiVolume.StoragegatewayCachedIscsiVolume",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayCachedIscsiVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.storagegatewayCachedIscsiVolume.StoragegatewayCachedIscsiVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetKmsEncrypted() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsEncrypted",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetKmsKey() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		s,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetSourceVolumeArn() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceVolumeArn",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayCachedIscsiVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


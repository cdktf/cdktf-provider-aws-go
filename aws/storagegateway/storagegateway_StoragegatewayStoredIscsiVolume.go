package storagegateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-aws-go/aws/v9/storagegateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume aws_storagegateway_stored_iscsi_volume}.
type StoragegatewayStoredIscsiVolume interface {
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskId() *string
	SetDiskId(val *string)
	DiskIdInput() *string
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
	PreserveExistingData() interface{}
	SetPreserveExistingData(val interface{})
	PreserveExistingDataInput() interface{}
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
	VolumeAttachmentStatus() *string
	VolumeId() *string
	VolumeSizeInBytes() *float64
	VolumeStatus() *string
	VolumeType() *string
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
	ResetId()
	ResetKmsEncrypted()
	ResetKmsKey()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSnapshotId()
	ResetTags()
	ResetTagsAll()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for StoragegatewayStoredIscsiVolume
type jsiiProxy_StoragegatewayStoredIscsiVolume struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) ChapEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"chapEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) DiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) DiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) GatewayArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) GatewayArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) KmsEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) KmsEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) LunNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lunNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) NetworkInterfaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) NetworkInterfaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkInterfaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) NetworkInterfacePort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"networkInterfacePort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) PreserveExistingData() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveExistingData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) PreserveExistingDataInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveExistingDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) SnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TargetArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TargetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TargetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) VolumeAttachmentStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeAttachmentStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) VolumeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) VolumeSizeInBytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"volumeSizeInBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) VolumeStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume) VolumeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"volumeType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume aws_storagegateway_stored_iscsi_volume} Resource.
func NewStoragegatewayStoredIscsiVolume(scope constructs.Construct, id *string, config *StoragegatewayStoredIscsiVolumeConfig) StoragegatewayStoredIscsiVolume {
	_init_.Initialize()

	if err := validateNewStoragegatewayStoredIscsiVolumeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StoragegatewayStoredIscsiVolume{}

	_jsii_.Create(
		"@cdktf/provider-aws.storagegateway.StoragegatewayStoredIscsiVolume",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/aws/r/storagegateway_stored_iscsi_volume aws_storagegateway_stored_iscsi_volume} Resource.
func NewStoragegatewayStoredIscsiVolume_Override(s StoragegatewayStoredIscsiVolume, scope constructs.Construct, id *string, config *StoragegatewayStoredIscsiVolumeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.storagegateway.StoragegatewayStoredIscsiVolume",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetDiskId(val *string) {
	if err := j.validateSetDiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetGatewayArn(val *string) {
	if err := j.validateSetGatewayArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayArn",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetKmsEncrypted(val interface{}) {
	if err := j.validateSetKmsEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsEncrypted",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetNetworkInterfaceId(val *string) {
	if err := j.validateSetNetworkInterfaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetPreserveExistingData(val interface{}) {
	if err := j.validateSetPreserveExistingDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveExistingData",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetSnapshotId(val *string) {
	if err := j.validateSetSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotId",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_StoragegatewayStoredIscsiVolume)SetTargetName(val *string) {
	if err := j.validateSetTargetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetName",
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
func StoragegatewayStoredIscsiVolume_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStoragegatewayStoredIscsiVolume_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.storagegateway.StoragegatewayStoredIscsiVolume",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StoragegatewayStoredIscsiVolume_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.storagegateway.StoragegatewayStoredIscsiVolume",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetKmsEncrypted() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsEncrypted",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetKmsKey() {
	_jsii_.InvokeVoid(
		s,
		"resetKmsKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetSnapshotId() {
	_jsii_.InvokeVoid(
		s,
		"resetSnapshotId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ResetTagsAll() {
	_jsii_.InvokeVoid(
		s,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StoragegatewayStoredIscsiVolume) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


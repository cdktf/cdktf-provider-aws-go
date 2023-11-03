// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbsnapshotcopy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/dbsnapshotcopy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/db_snapshot_copy aws_db_snapshot_copy}.
type DbSnapshotCopy interface {
	cdktf.TerraformResource
	AllocatedStorage() *float64
	AvailabilityZone() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyTags() interface{}
	SetCopyTags(val interface{})
	CopyTagsInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DbSnapshotArn() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DestinationRegion() *string
	SetDestinationRegion(val *string)
	DestinationRegionInput() *string
	Encrypted() cdktf.IResolvable
	Engine() *string
	EngineVersion() *string
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
	Iops() *float64
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	LicenseModel() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OptionGroupName() *string
	SetOptionGroupName(val *string)
	OptionGroupNameInput() *string
	Port() *float64
	PresignedUrl() *string
	SetPresignedUrl(val *string)
	PresignedUrlInput() *string
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
	SnapshotType() *string
	SourceDbSnapshotIdentifier() *string
	SetSourceDbSnapshotIdentifier(val *string)
	SourceDbSnapshotIdentifierInput() *string
	SourceRegion() *string
	StorageType() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	TargetCustomAvailabilityZone() *string
	SetTargetCustomAvailabilityZone(val *string)
	TargetCustomAvailabilityZoneInput() *string
	TargetDbSnapshotIdentifier() *string
	SetTargetDbSnapshotIdentifier(val *string)
	TargetDbSnapshotIdentifierInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DbSnapshotCopyTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcId() *string
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
	PutTimeouts(value *DbSnapshotCopyTimeouts)
	ResetCopyTags()
	ResetDestinationRegion()
	ResetId()
	ResetKmsKeyId()
	ResetOptionGroupName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPresignedUrl()
	ResetTags()
	ResetTagsAll()
	ResetTargetCustomAvailabilityZone()
	ResetTimeouts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DbSnapshotCopy
type jsiiProxy_DbSnapshotCopy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DbSnapshotCopy) AllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) CopyTags() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) CopyTagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) DbSnapshotArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSnapshotArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) DestinationRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) DestinationRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"destinationRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Encrypted() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) OptionGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) PresignedUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presignedUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) PresignedUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"presignedUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) SnapshotType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snapshotType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) SourceDbSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) SourceDbSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) TargetCustomAvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCustomAvailabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) TargetCustomAvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetCustomAvailabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) TargetDbSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) TargetDbSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetDbSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) Timeouts() DbSnapshotCopyTimeoutsOutputReference {
	var returns DbSnapshotCopyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DbSnapshotCopy) VpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vpcId",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/db_snapshot_copy aws_db_snapshot_copy} Resource.
func NewDbSnapshotCopy(scope constructs.Construct, id *string, config *DbSnapshotCopyConfig) DbSnapshotCopy {
	_init_.Initialize()

	if err := validateNewDbSnapshotCopyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DbSnapshotCopy{}

	_jsii_.Create(
		"@cdktf/provider-aws.dbSnapshotCopy.DbSnapshotCopy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/db_snapshot_copy aws_db_snapshot_copy} Resource.
func NewDbSnapshotCopy_Override(d DbSnapshotCopy, scope constructs.Construct, id *string, config *DbSnapshotCopyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dbSnapshotCopy.DbSnapshotCopy",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetCopyTags(val interface{}) {
	if err := j.validateSetCopyTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTags",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetDestinationRegion(val *string) {
	if err := j.validateSetDestinationRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"destinationRegion",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetOptionGroupName(val *string) {
	if err := j.validateSetOptionGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionGroupName",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetPresignedUrl(val *string) {
	if err := j.validateSetPresignedUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"presignedUrl",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetSourceDbSnapshotIdentifier(val *string) {
	if err := j.validateSetSourceDbSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetTargetCustomAvailabilityZone(val *string) {
	if err := j.validateSetTargetCustomAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetCustomAvailabilityZone",
		val,
	)
}

func (j *jsiiProxy_DbSnapshotCopy)SetTargetDbSnapshotIdentifier(val *string) {
	if err := j.validateSetTargetDbSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetDbSnapshotIdentifier",
		val,
	)
}

// Generates CDKTF code for importing a DbSnapshotCopy resource upon running "cdktf plan <stack-name>".
func DbSnapshotCopy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDbSnapshotCopy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dbSnapshotCopy.DbSnapshotCopy",
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
func DbSnapshotCopy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDbSnapshotCopy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dbSnapshotCopy.DbSnapshotCopy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DbSnapshotCopy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDbSnapshotCopy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dbSnapshotCopy.DbSnapshotCopy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DbSnapshotCopy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDbSnapshotCopy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dbSnapshotCopy.DbSnapshotCopy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DbSnapshotCopy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dbSnapshotCopy.DbSnapshotCopy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DbSnapshotCopy) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DbSnapshotCopy) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DbSnapshotCopy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DbSnapshotCopy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DbSnapshotCopy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DbSnapshotCopy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DbSnapshotCopy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DbSnapshotCopy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DbSnapshotCopy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DbSnapshotCopy) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DbSnapshotCopy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DbSnapshotCopy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DbSnapshotCopy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DbSnapshotCopy) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DbSnapshotCopy) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DbSnapshotCopy) PutTimeouts(value *DbSnapshotCopyTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DbSnapshotCopy) ResetCopyTags() {
	_jsii_.InvokeVoid(
		d,
		"resetCopyTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbSnapshotCopy) ResetDestinationRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetDestinationRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbSnapshotCopy) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbSnapshotCopy) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		d,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbSnapshotCopy) ResetOptionGroupName() {
	_jsii_.InvokeVoid(
		d,
		"resetOptionGroupName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbSnapshotCopy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbSnapshotCopy) ResetPresignedUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetPresignedUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbSnapshotCopy) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbSnapshotCopy) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbSnapshotCopy) ResetTargetCustomAvailabilityZone() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetCustomAvailabilityZone",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbSnapshotCopy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DbSnapshotCopy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbSnapshotCopy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbSnapshotCopy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DbSnapshotCopy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


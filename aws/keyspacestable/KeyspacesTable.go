// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyspacestable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/keyspacestable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/keyspaces_table aws_keyspaces_table}.
type KeyspacesTable interface {
	cdktf.TerraformResource
	Arn() *string
	CapacitySpecification() KeyspacesTableCapacitySpecificationOutputReference
	CapacitySpecificationInput() *KeyspacesTableCapacitySpecification
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientSideTimestamps() KeyspacesTableClientSideTimestampsOutputReference
	ClientSideTimestampsInput() *KeyspacesTableClientSideTimestamps
	Comment() KeyspacesTableCommentOutputReference
	CommentInput() *KeyspacesTableComment
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
	DefaultTimeToLive() *float64
	SetDefaultTimeToLive(val *float64)
	DefaultTimeToLiveInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EncryptionSpecification() KeyspacesTableEncryptionSpecificationOutputReference
	EncryptionSpecificationInput() *KeyspacesTableEncryptionSpecification
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
	KeyspaceName() *string
	SetKeyspaceName(val *string)
	KeyspaceNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PointInTimeRecovery() KeyspacesTablePointInTimeRecoveryOutputReference
	PointInTimeRecoveryInput() *KeyspacesTablePointInTimeRecovery
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
	SchemaDefinition() KeyspacesTableSchemaDefinitionOutputReference
	SchemaDefinitionInput() *KeyspacesTableSchemaDefinition
	TableName() *string
	SetTableName(val *string)
	TableNameInput() *string
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
	Timeouts() KeyspacesTableTimeoutsOutputReference
	TimeoutsInput() interface{}
	Ttl() KeyspacesTableTtlOutputReference
	TtlInput() *KeyspacesTableTtl
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
	PutCapacitySpecification(value *KeyspacesTableCapacitySpecification)
	PutClientSideTimestamps(value *KeyspacesTableClientSideTimestamps)
	PutComment(value *KeyspacesTableComment)
	PutEncryptionSpecification(value *KeyspacesTableEncryptionSpecification)
	PutPointInTimeRecovery(value *KeyspacesTablePointInTimeRecovery)
	PutSchemaDefinition(value *KeyspacesTableSchemaDefinition)
	PutTimeouts(value *KeyspacesTableTimeouts)
	PutTtl(value *KeyspacesTableTtl)
	ResetCapacitySpecification()
	ResetClientSideTimestamps()
	ResetComment()
	ResetDefaultTimeToLive()
	ResetEncryptionSpecification()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeRecovery()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetTtl()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for KeyspacesTable
type jsiiProxy_KeyspacesTable struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KeyspacesTable) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) CapacitySpecification() KeyspacesTableCapacitySpecificationOutputReference {
	var returns KeyspacesTableCapacitySpecificationOutputReference
	_jsii_.Get(
		j,
		"capacitySpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) CapacitySpecificationInput() *KeyspacesTableCapacitySpecification {
	var returns *KeyspacesTableCapacitySpecification
	_jsii_.Get(
		j,
		"capacitySpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) ClientSideTimestamps() KeyspacesTableClientSideTimestampsOutputReference {
	var returns KeyspacesTableClientSideTimestampsOutputReference
	_jsii_.Get(
		j,
		"clientSideTimestamps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) ClientSideTimestampsInput() *KeyspacesTableClientSideTimestamps {
	var returns *KeyspacesTableClientSideTimestamps
	_jsii_.Get(
		j,
		"clientSideTimestampsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Comment() KeyspacesTableCommentOutputReference {
	var returns KeyspacesTableCommentOutputReference
	_jsii_.Get(
		j,
		"comment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) CommentInput() *KeyspacesTableComment {
	var returns *KeyspacesTableComment
	_jsii_.Get(
		j,
		"commentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) DefaultTimeToLive() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTimeToLive",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) DefaultTimeToLiveInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTimeToLiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) EncryptionSpecification() KeyspacesTableEncryptionSpecificationOutputReference {
	var returns KeyspacesTableEncryptionSpecificationOutputReference
	_jsii_.Get(
		j,
		"encryptionSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) EncryptionSpecificationInput() *KeyspacesTableEncryptionSpecification {
	var returns *KeyspacesTableEncryptionSpecification
	_jsii_.Get(
		j,
		"encryptionSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) KeyspaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyspaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) KeyspaceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyspaceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) PointInTimeRecovery() KeyspacesTablePointInTimeRecoveryOutputReference {
	var returns KeyspacesTablePointInTimeRecoveryOutputReference
	_jsii_.Get(
		j,
		"pointInTimeRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) PointInTimeRecoveryInput() *KeyspacesTablePointInTimeRecovery {
	var returns *KeyspacesTablePointInTimeRecovery
	_jsii_.Get(
		j,
		"pointInTimeRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) SchemaDefinition() KeyspacesTableSchemaDefinitionOutputReference {
	var returns KeyspacesTableSchemaDefinitionOutputReference
	_jsii_.Get(
		j,
		"schemaDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) SchemaDefinitionInput() *KeyspacesTableSchemaDefinition {
	var returns *KeyspacesTableSchemaDefinition
	_jsii_.Get(
		j,
		"schemaDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Timeouts() KeyspacesTableTimeoutsOutputReference {
	var returns KeyspacesTableTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) Ttl() KeyspacesTableTtlOutputReference {
	var returns KeyspacesTableTtlOutputReference
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyspacesTable) TtlInput() *KeyspacesTableTtl {
	var returns *KeyspacesTableTtl
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/keyspaces_table aws_keyspaces_table} Resource.
func NewKeyspacesTable(scope constructs.Construct, id *string, config *KeyspacesTableConfig) KeyspacesTable {
	_init_.Initialize()

	if err := validateNewKeyspacesTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyspacesTable{}

	_jsii_.Create(
		"@cdktf/provider-aws.keyspacesTable.KeyspacesTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/keyspaces_table aws_keyspaces_table} Resource.
func NewKeyspacesTable_Override(k KeyspacesTable, scope constructs.Construct, id *string, config *KeyspacesTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.keyspacesTable.KeyspacesTable",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetDefaultTimeToLive(val *float64) {
	if err := j.validateSetDefaultTimeToLiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTimeToLive",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetKeyspaceName(val *string) {
	if err := j.validateSetKeyspaceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyspaceName",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_KeyspacesTable)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a KeyspacesTable resource upon running "cdktf plan <stack-name>".
func KeyspacesTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKeyspacesTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.keyspacesTable.KeyspacesTable",
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
func KeyspacesTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyspacesTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.keyspacesTable.KeyspacesTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KeyspacesTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyspacesTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.keyspacesTable.KeyspacesTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KeyspacesTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyspacesTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.keyspacesTable.KeyspacesTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KeyspacesTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.keyspacesTable.KeyspacesTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KeyspacesTable) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KeyspacesTable) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KeyspacesTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KeyspacesTable) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KeyspacesTable) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KeyspacesTable) PutCapacitySpecification(value *KeyspacesTableCapacitySpecification) {
	if err := k.validatePutCapacitySpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putCapacitySpecification",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTable) PutClientSideTimestamps(value *KeyspacesTableClientSideTimestamps) {
	if err := k.validatePutClientSideTimestampsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putClientSideTimestamps",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTable) PutComment(value *KeyspacesTableComment) {
	if err := k.validatePutCommentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putComment",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTable) PutEncryptionSpecification(value *KeyspacesTableEncryptionSpecification) {
	if err := k.validatePutEncryptionSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putEncryptionSpecification",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTable) PutPointInTimeRecovery(value *KeyspacesTablePointInTimeRecovery) {
	if err := k.validatePutPointInTimeRecoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putPointInTimeRecovery",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTable) PutSchemaDefinition(value *KeyspacesTableSchemaDefinition) {
	if err := k.validatePutSchemaDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putSchemaDefinition",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTable) PutTimeouts(value *KeyspacesTableTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTable) PutTtl(value *KeyspacesTableTtl) {
	if err := k.validatePutTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTtl",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetCapacitySpecification() {
	_jsii_.InvokeVoid(
		k,
		"resetCapacitySpecification",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetClientSideTimestamps() {
	_jsii_.InvokeVoid(
		k,
		"resetClientSideTimestamps",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetComment() {
	_jsii_.InvokeVoid(
		k,
		"resetComment",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetDefaultTimeToLive() {
	_jsii_.InvokeVoid(
		k,
		"resetDefaultTimeToLive",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetEncryptionSpecification() {
	_jsii_.InvokeVoid(
		k,
		"resetEncryptionSpecification",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetPointInTimeRecovery() {
	_jsii_.InvokeVoid(
		k,
		"resetPointInTimeRecovery",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetTags() {
	_jsii_.InvokeVoid(
		k,
		"resetTags",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetTagsAll() {
	_jsii_.InvokeVoid(
		k,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) ResetTtl() {
	_jsii_.InvokeVoid(
		k,
		"resetTtl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyspacesTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyspacesTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


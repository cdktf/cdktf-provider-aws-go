// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dynamodbtable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/dynamodb_table aws_dynamodb_table}.
type DynamodbTable interface {
	cdktf.TerraformResource
	Arn() *string
	Attribute() DynamodbTableAttributeList
	AttributeInput() interface{}
	BillingMode() *string
	SetBillingMode(val *string)
	BillingModeInput() *string
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
	DeletionProtectionEnabled() interface{}
	SetDeletionProtectionEnabled(val interface{})
	DeletionProtectionEnabledInput() interface{}
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
	GlobalSecondaryIndex() DynamodbTableGlobalSecondaryIndexList
	GlobalSecondaryIndexInput() interface{}
	HashKey() *string
	SetHashKey(val *string)
	HashKeyInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImportTable() DynamodbTableImportTableOutputReference
	ImportTableInput() *DynamodbTableImportTable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocalSecondaryIndex() DynamodbTableLocalSecondaryIndexList
	LocalSecondaryIndexInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PointInTimeRecovery() DynamodbTablePointInTimeRecoveryOutputReference
	PointInTimeRecoveryInput() *DynamodbTablePointInTimeRecovery
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	RangeKey() *string
	SetRangeKey(val *string)
	RangeKeyInput() *string
	// Experimental.
	RawOverrides() interface{}
	ReadCapacity() *float64
	SetReadCapacity(val *float64)
	ReadCapacityInput() *float64
	Replica() DynamodbTableReplicaList
	ReplicaInput() interface{}
	RestoreDateTime() *string
	SetRestoreDateTime(val *string)
	RestoreDateTimeInput() *string
	RestoreSourceName() *string
	SetRestoreSourceName(val *string)
	RestoreSourceNameInput() *string
	RestoreToLatestTime() interface{}
	SetRestoreToLatestTime(val interface{})
	RestoreToLatestTimeInput() interface{}
	ServerSideEncryption() DynamodbTableServerSideEncryptionOutputReference
	ServerSideEncryptionInput() *DynamodbTableServerSideEncryption
	StreamArn() *string
	StreamEnabled() interface{}
	SetStreamEnabled(val interface{})
	StreamEnabledInput() interface{}
	StreamLabel() *string
	StreamViewType() *string
	SetStreamViewType(val *string)
	StreamViewTypeInput() *string
	TableClass() *string
	SetTableClass(val *string)
	TableClassInput() *string
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
	Timeouts() DynamodbTableTimeoutsOutputReference
	TimeoutsInput() interface{}
	Ttl() DynamodbTableTtlOutputReference
	TtlInput() *DynamodbTableTtl
	WriteCapacity() *float64
	SetWriteCapacity(val *float64)
	WriteCapacityInput() *float64
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
	PutAttribute(value interface{})
	PutGlobalSecondaryIndex(value interface{})
	PutImportTable(value *DynamodbTableImportTable)
	PutLocalSecondaryIndex(value interface{})
	PutPointInTimeRecovery(value *DynamodbTablePointInTimeRecovery)
	PutReplica(value interface{})
	PutServerSideEncryption(value *DynamodbTableServerSideEncryption)
	PutTimeouts(value *DynamodbTableTimeouts)
	PutTtl(value *DynamodbTableTtl)
	ResetAttribute()
	ResetBillingMode()
	ResetDeletionProtectionEnabled()
	ResetGlobalSecondaryIndex()
	ResetHashKey()
	ResetId()
	ResetImportTable()
	ResetLocalSecondaryIndex()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPointInTimeRecovery()
	ResetRangeKey()
	ResetReadCapacity()
	ResetReplica()
	ResetRestoreDateTime()
	ResetRestoreSourceName()
	ResetRestoreToLatestTime()
	ResetServerSideEncryption()
	ResetStreamEnabled()
	ResetStreamViewType()
	ResetTableClass()
	ResetTags()
	ResetTagsAll()
	ResetTimeouts()
	ResetTtl()
	ResetWriteCapacity()
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

// The jsii proxy struct for DynamodbTable
type jsiiProxy_DynamodbTable struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DynamodbTable) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Attribute() DynamodbTableAttributeList {
	var returns DynamodbTableAttributeList
	_jsii_.Get(
		j,
		"attribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) AttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"attributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) BillingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) BillingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) DeletionProtectionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) DeletionProtectionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) GlobalSecondaryIndex() DynamodbTableGlobalSecondaryIndexList {
	var returns DynamodbTableGlobalSecondaryIndexList
	_jsii_.Get(
		j,
		"globalSecondaryIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) GlobalSecondaryIndexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"globalSecondaryIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) HashKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) HashKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hashKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ImportTable() DynamodbTableImportTableOutputReference {
	var returns DynamodbTableImportTableOutputReference
	_jsii_.Get(
		j,
		"importTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ImportTableInput() *DynamodbTableImportTable {
	var returns *DynamodbTableImportTable
	_jsii_.Get(
		j,
		"importTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) LocalSecondaryIndex() DynamodbTableLocalSecondaryIndexList {
	var returns DynamodbTableLocalSecondaryIndexList
	_jsii_.Get(
		j,
		"localSecondaryIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) LocalSecondaryIndexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localSecondaryIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) PointInTimeRecovery() DynamodbTablePointInTimeRecoveryOutputReference {
	var returns DynamodbTablePointInTimeRecoveryOutputReference
	_jsii_.Get(
		j,
		"pointInTimeRecovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) PointInTimeRecoveryInput() *DynamodbTablePointInTimeRecovery {
	var returns *DynamodbTablePointInTimeRecovery
	_jsii_.Get(
		j,
		"pointInTimeRecoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) RangeKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) RangeKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rangeKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ReadCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ReadCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Replica() DynamodbTableReplicaList {
	var returns DynamodbTableReplicaList
	_jsii_.Get(
		j,
		"replica",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ReplicaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replicaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) RestoreDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) RestoreDateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreDateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) RestoreSourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreSourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) RestoreSourceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreSourceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) RestoreToLatestTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restoreToLatestTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) RestoreToLatestTimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restoreToLatestTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ServerSideEncryption() DynamodbTableServerSideEncryptionOutputReference {
	var returns DynamodbTableServerSideEncryptionOutputReference
	_jsii_.Get(
		j,
		"serverSideEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) ServerSideEncryptionInput() *DynamodbTableServerSideEncryption {
	var returns *DynamodbTableServerSideEncryption
	_jsii_.Get(
		j,
		"serverSideEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"streamEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamViewType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamViewType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) StreamViewTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamViewTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TableClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TableClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Timeouts() DynamodbTableTimeoutsOutputReference {
	var returns DynamodbTableTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) Ttl() DynamodbTableTtlOutputReference {
	var returns DynamodbTableTtlOutputReference
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) TtlInput() *DynamodbTableTtl {
	var returns *DynamodbTableTtl
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) WriteCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamodbTable) WriteCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeCapacityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/dynamodb_table aws_dynamodb_table} Resource.
func NewDynamodbTable(scope constructs.Construct, id *string, config *DynamodbTableConfig) DynamodbTable {
	_init_.Initialize()

	if err := validateNewDynamodbTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamodbTable{}

	_jsii_.Create(
		"@cdktf/provider-aws.dynamodbTable.DynamodbTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.0/docs/resources/dynamodb_table aws_dynamodb_table} Resource.
func NewDynamodbTable_Override(d DynamodbTable, scope constructs.Construct, id *string, config *DynamodbTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dynamodbTable.DynamodbTable",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DynamodbTable)SetBillingMode(val *string) {
	if err := j.validateSetBillingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingMode",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetDeletionProtectionEnabled(val interface{}) {
	if err := j.validateSetDeletionProtectionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtectionEnabled",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetHashKey(val *string) {
	if err := j.validateSetHashKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hashKey",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetRangeKey(val *string) {
	if err := j.validateSetRangeKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeKey",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetReadCapacity(val *float64) {
	if err := j.validateSetReadCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readCapacity",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetRestoreDateTime(val *string) {
	if err := j.validateSetRestoreDateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreDateTime",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetRestoreSourceName(val *string) {
	if err := j.validateSetRestoreSourceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreSourceName",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetRestoreToLatestTime(val interface{}) {
	if err := j.validateSetRestoreToLatestTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreToLatestTime",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetStreamEnabled(val interface{}) {
	if err := j.validateSetStreamEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamEnabled",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetStreamViewType(val *string) {
	if err := j.validateSetStreamViewTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamViewType",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetTableClass(val *string) {
	if err := j.validateSetTableClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableClass",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_DynamodbTable)SetWriteCapacity(val *float64) {
	if err := j.validateSetWriteCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeCapacity",
		val,
	)
}

// Generates CDKTF code for importing a DynamodbTable resource upon running "cdktf plan <stack-name>".
func DynamodbTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDynamodbTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dynamodbTable.DynamodbTable",
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
func DynamodbTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dynamodbTable.DynamodbTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dynamodbTable.DynamodbTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DynamodbTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamodbTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dynamodbTable.DynamodbTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DynamodbTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dynamodbTable.DynamodbTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DynamodbTable) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DynamodbTable) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DynamodbTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DynamodbTable) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DynamodbTable) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DynamodbTable) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DynamodbTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DynamodbTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DynamodbTable) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DynamodbTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DynamodbTable) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DynamodbTable) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DynamodbTable) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamodbTable) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DynamodbTable) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DynamodbTable) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DynamodbTable) PutAttribute(value interface{}) {
	if err := d.validatePutAttributeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAttribute",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutGlobalSecondaryIndex(value interface{}) {
	if err := d.validatePutGlobalSecondaryIndexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putGlobalSecondaryIndex",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutImportTable(value *DynamodbTableImportTable) {
	if err := d.validatePutImportTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putImportTable",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutLocalSecondaryIndex(value interface{}) {
	if err := d.validatePutLocalSecondaryIndexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putLocalSecondaryIndex",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutPointInTimeRecovery(value *DynamodbTablePointInTimeRecovery) {
	if err := d.validatePutPointInTimeRecoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPointInTimeRecovery",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutReplica(value interface{}) {
	if err := d.validatePutReplicaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putReplica",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutServerSideEncryption(value *DynamodbTableServerSideEncryption) {
	if err := d.validatePutServerSideEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putServerSideEncryption",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutTimeouts(value *DynamodbTableTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) PutTtl(value *DynamodbTableTtl) {
	if err := d.validatePutTtlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTtl",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DynamodbTable) ResetAttribute() {
	_jsii_.InvokeVoid(
		d,
		"resetAttribute",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetBillingMode() {
	_jsii_.InvokeVoid(
		d,
		"resetBillingMode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetDeletionProtectionEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDeletionProtectionEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetGlobalSecondaryIndex() {
	_jsii_.InvokeVoid(
		d,
		"resetGlobalSecondaryIndex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetHashKey() {
	_jsii_.InvokeVoid(
		d,
		"resetHashKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetImportTable() {
	_jsii_.InvokeVoid(
		d,
		"resetImportTable",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetLocalSecondaryIndex() {
	_jsii_.InvokeVoid(
		d,
		"resetLocalSecondaryIndex",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetPointInTimeRecovery() {
	_jsii_.InvokeVoid(
		d,
		"resetPointInTimeRecovery",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetRangeKey() {
	_jsii_.InvokeVoid(
		d,
		"resetRangeKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetReadCapacity() {
	_jsii_.InvokeVoid(
		d,
		"resetReadCapacity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetReplica() {
	_jsii_.InvokeVoid(
		d,
		"resetReplica",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetRestoreDateTime() {
	_jsii_.InvokeVoid(
		d,
		"resetRestoreDateTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetRestoreSourceName() {
	_jsii_.InvokeVoid(
		d,
		"resetRestoreSourceName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetRestoreToLatestTime() {
	_jsii_.InvokeVoid(
		d,
		"resetRestoreToLatestTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetServerSideEncryption() {
	_jsii_.InvokeVoid(
		d,
		"resetServerSideEncryption",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetStreamEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetStreamEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetStreamViewType() {
	_jsii_.InvokeVoid(
		d,
		"resetStreamViewType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTableClass() {
	_jsii_.InvokeVoid(
		d,
		"resetTableClass",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTagsAll() {
	_jsii_.InvokeVoid(
		d,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) ResetWriteCapacity() {
	_jsii_.InvokeVoid(
		d,
		"resetWriteCapacity",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DynamodbTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamodbTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


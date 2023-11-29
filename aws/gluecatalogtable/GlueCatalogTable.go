// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecatalogtable

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/gluecatalogtable/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/glue_catalog_table aws_glue_catalog_table}.
type GlueCatalogTable interface {
	cdktf.TerraformResource
	Arn() *string
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
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
	DatabaseName() *string
	SetDatabaseName(val *string)
	DatabaseNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	// The tree node.
	Node() constructs.Node
	OpenTableFormatInput() GlueCatalogTableOpenTableFormatInputOutputReference
	OpenTableFormatInputInput() *GlueCatalogTableOpenTableFormatInput
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	PartitionIndex() GlueCatalogTablePartitionIndexList
	PartitionIndexInput() interface{}
	PartitionKeys() GlueCatalogTablePartitionKeysList
	PartitionKeysInput() interface{}
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
	Retention() *float64
	SetRetention(val *float64)
	RetentionInput() *float64
	StorageDescriptor() GlueCatalogTableStorageDescriptorOutputReference
	StorageDescriptorInput() *GlueCatalogTableStorageDescriptor
	TableType() *string
	SetTableType(val *string)
	TableTypeInput() *string
	TargetTable() GlueCatalogTableTargetTableOutputReference
	TargetTableInput() *GlueCatalogTableTargetTable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ViewExpandedText() *string
	SetViewExpandedText(val *string)
	ViewExpandedTextInput() *string
	ViewOriginalText() *string
	SetViewOriginalText(val *string)
	ViewOriginalTextInput() *string
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
	PutOpenTableFormatInput(value *GlueCatalogTableOpenTableFormatInput)
	PutPartitionIndex(value interface{})
	PutPartitionKeys(value interface{})
	PutStorageDescriptor(value *GlueCatalogTableStorageDescriptor)
	PutTargetTable(value *GlueCatalogTableTargetTable)
	ResetCatalogId()
	ResetDescription()
	ResetId()
	ResetOpenTableFormatInput()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwner()
	ResetParameters()
	ResetPartitionIndex()
	ResetPartitionKeys()
	ResetRetention()
	ResetStorageDescriptor()
	ResetTableType()
	ResetTargetTable()
	ResetViewExpandedText()
	ResetViewOriginalText()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GlueCatalogTable
type jsiiProxy_GlueCatalogTable struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GlueCatalogTable) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) OpenTableFormatInput() GlueCatalogTableOpenTableFormatInputOutputReference {
	var returns GlueCatalogTableOpenTableFormatInputOutputReference
	_jsii_.Get(
		j,
		"openTableFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) OpenTableFormatInputInput() *GlueCatalogTableOpenTableFormatInput {
	var returns *GlueCatalogTableOpenTableFormatInput
	_jsii_.Get(
		j,
		"openTableFormatInputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) PartitionIndex() GlueCatalogTablePartitionIndexList {
	var returns GlueCatalogTablePartitionIndexList
	_jsii_.Get(
		j,
		"partitionIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) PartitionIndexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) PartitionKeys() GlueCatalogTablePartitionKeysList {
	var returns GlueCatalogTablePartitionKeysList
	_jsii_.Get(
		j,
		"partitionKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) PartitionKeysInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"partitionKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) Retention() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) RetentionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) StorageDescriptor() GlueCatalogTableStorageDescriptorOutputReference {
	var returns GlueCatalogTableStorageDescriptorOutputReference
	_jsii_.Get(
		j,
		"storageDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) StorageDescriptorInput() *GlueCatalogTableStorageDescriptor {
	var returns *GlueCatalogTableStorageDescriptor
	_jsii_.Get(
		j,
		"storageDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TableType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TableTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TargetTable() GlueCatalogTableTargetTableOutputReference {
	var returns GlueCatalogTableTargetTableOutputReference
	_jsii_.Get(
		j,
		"targetTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TargetTableInput() *GlueCatalogTableTargetTable {
	var returns *GlueCatalogTableTargetTable
	_jsii_.Get(
		j,
		"targetTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ViewExpandedText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewExpandedText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ViewExpandedTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewExpandedTextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ViewOriginalText() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewOriginalText",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueCatalogTable) ViewOriginalTextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewOriginalTextInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/glue_catalog_table aws_glue_catalog_table} Resource.
func NewGlueCatalogTable(scope constructs.Construct, id *string, config *GlueCatalogTableConfig) GlueCatalogTable {
	_init_.Initialize()

	if err := validateNewGlueCatalogTableParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueCatalogTable{}

	_jsii_.Create(
		"@cdktf/provider-aws.glueCatalogTable.GlueCatalogTable",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/glue_catalog_table aws_glue_catalog_table} Resource.
func NewGlueCatalogTable_Override(g GlueCatalogTable, scope constructs.Construct, id *string, config *GlueCatalogTableConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.glueCatalogTable.GlueCatalogTable",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetCatalogId(val *string) {
	if err := j.validateSetCatalogIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetRetention(val *float64) {
	if err := j.validateSetRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retention",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetTableType(val *string) {
	if err := j.validateSetTableTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableType",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetViewExpandedText(val *string) {
	if err := j.validateSetViewExpandedTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewExpandedText",
		val,
	)
}

func (j *jsiiProxy_GlueCatalogTable)SetViewOriginalText(val *string) {
	if err := j.validateSetViewOriginalTextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewOriginalText",
		val,
	)
}

// Generates CDKTF code for importing a GlueCatalogTable resource upon running "cdktf plan <stack-name>".
func GlueCatalogTable_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGlueCatalogTable_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.glueCatalogTable.GlueCatalogTable",
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
func GlueCatalogTable_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueCatalogTable_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.glueCatalogTable.GlueCatalogTable",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlueCatalogTable_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueCatalogTable_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.glueCatalogTable.GlueCatalogTable",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GlueCatalogTable_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGlueCatalogTable_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.glueCatalogTable.GlueCatalogTable",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GlueCatalogTable_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.glueCatalogTable.GlueCatalogTable",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GlueCatalogTable) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GlueCatalogTable) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GlueCatalogTable) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GlueCatalogTable) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GlueCatalogTable) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GlueCatalogTable) PutOpenTableFormatInput(value *GlueCatalogTableOpenTableFormatInput) {
	if err := g.validatePutOpenTableFormatInputParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putOpenTableFormatInput",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTable) PutPartitionIndex(value interface{}) {
	if err := g.validatePutPartitionIndexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPartitionIndex",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTable) PutPartitionKeys(value interface{}) {
	if err := g.validatePutPartitionKeysParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPartitionKeys",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTable) PutStorageDescriptor(value *GlueCatalogTableStorageDescriptor) {
	if err := g.validatePutStorageDescriptorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorageDescriptor",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTable) PutTargetTable(value *GlueCatalogTableTargetTable) {
	if err := g.validatePutTargetTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTargetTable",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetCatalogId() {
	_jsii_.InvokeVoid(
		g,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetDescription() {
	_jsii_.InvokeVoid(
		g,
		"resetDescription",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetOpenTableFormatInput() {
	_jsii_.InvokeVoid(
		g,
		"resetOpenTableFormatInput",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetOwner() {
	_jsii_.InvokeVoid(
		g,
		"resetOwner",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetParameters() {
	_jsii_.InvokeVoid(
		g,
		"resetParameters",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetPartitionIndex() {
	_jsii_.InvokeVoid(
		g,
		"resetPartitionIndex",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetPartitionKeys() {
	_jsii_.InvokeVoid(
		g,
		"resetPartitionKeys",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetRetention() {
	_jsii_.InvokeVoid(
		g,
		"resetRetention",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetStorageDescriptor() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageDescriptor",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetTableType() {
	_jsii_.InvokeVoid(
		g,
		"resetTableType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetTargetTable() {
	_jsii_.InvokeVoid(
		g,
		"resetTargetTable",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetViewExpandedText() {
	_jsii_.InvokeVoid(
		g,
		"resetViewExpandedText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) ResetViewOriginalText() {
	_jsii_.InvokeVoid(
		g,
		"resetViewOriginalText",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GlueCatalogTable) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueCatalogTable) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


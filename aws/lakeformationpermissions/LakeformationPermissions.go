// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lakeformationpermissions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/lakeformation_permissions aws_lakeformation_permissions}.
type LakeformationPermissions interface {
	cdktf.TerraformResource
	CatalogId() *string
	SetCatalogId(val *string)
	CatalogIdInput() *string
	CatalogResource() interface{}
	SetCatalogResource(val interface{})
	CatalogResourceInput() interface{}
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
	Database() LakeformationPermissionsDatabaseOutputReference
	DatabaseInput() *LakeformationPermissionsDatabase
	DataCellsFilter() LakeformationPermissionsDataCellsFilterOutputReference
	DataCellsFilterInput() *LakeformationPermissionsDataCellsFilter
	DataLocation() LakeformationPermissionsDataLocationOutputReference
	DataLocationInput() *LakeformationPermissionsDataLocation
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
	LfTag() LakeformationPermissionsLfTagOutputReference
	LfTagInput() *LakeformationPermissionsLfTag
	LfTagPolicy() LakeformationPermissionsLfTagPolicyOutputReference
	LfTagPolicyInput() *LakeformationPermissionsLfTagPolicy
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Permissions() *[]*string
	SetPermissions(val *[]*string)
	PermissionsInput() *[]*string
	PermissionsWithGrantOption() *[]*string
	SetPermissionsWithGrantOption(val *[]*string)
	PermissionsWithGrantOptionInput() *[]*string
	Principal() *string
	SetPrincipal(val *string)
	PrincipalInput() *string
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
	Table() LakeformationPermissionsTableOutputReference
	TableInput() *LakeformationPermissionsTable
	TableWithColumns() LakeformationPermissionsTableWithColumnsOutputReference
	TableWithColumnsInput() *LakeformationPermissionsTableWithColumns
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutDatabase(value *LakeformationPermissionsDatabase)
	PutDataCellsFilter(value *LakeformationPermissionsDataCellsFilter)
	PutDataLocation(value *LakeformationPermissionsDataLocation)
	PutLfTag(value *LakeformationPermissionsLfTag)
	PutLfTagPolicy(value *LakeformationPermissionsLfTagPolicy)
	PutTable(value *LakeformationPermissionsTable)
	PutTableWithColumns(value *LakeformationPermissionsTableWithColumns)
	ResetCatalogId()
	ResetCatalogResource()
	ResetDatabase()
	ResetDataCellsFilter()
	ResetDataLocation()
	ResetId()
	ResetLfTag()
	ResetLfTagPolicy()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermissionsWithGrantOption()
	ResetTable()
	ResetTableWithColumns()
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

// The jsii proxy struct for LakeformationPermissions
type jsiiProxy_LakeformationPermissions struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LakeformationPermissions) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) CatalogResource() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) CatalogResourceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Database() LakeformationPermissionsDatabaseOutputReference {
	var returns LakeformationPermissionsDatabaseOutputReference
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) DatabaseInput() *LakeformationPermissionsDatabase {
	var returns *LakeformationPermissionsDatabase
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) DataCellsFilter() LakeformationPermissionsDataCellsFilterOutputReference {
	var returns LakeformationPermissionsDataCellsFilterOutputReference
	_jsii_.Get(
		j,
		"dataCellsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) DataCellsFilterInput() *LakeformationPermissionsDataCellsFilter {
	var returns *LakeformationPermissionsDataCellsFilter
	_jsii_.Get(
		j,
		"dataCellsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) DataLocation() LakeformationPermissionsDataLocationOutputReference {
	var returns LakeformationPermissionsDataLocationOutputReference
	_jsii_.Get(
		j,
		"dataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) DataLocationInput() *LakeformationPermissionsDataLocation {
	var returns *LakeformationPermissionsDataLocation
	_jsii_.Get(
		j,
		"dataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) LfTag() LakeformationPermissionsLfTagOutputReference {
	var returns LakeformationPermissionsLfTagOutputReference
	_jsii_.Get(
		j,
		"lfTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) LfTagInput() *LakeformationPermissionsLfTag {
	var returns *LakeformationPermissionsLfTag
	_jsii_.Get(
		j,
		"lfTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) LfTagPolicy() LakeformationPermissionsLfTagPolicyOutputReference {
	var returns LakeformationPermissionsLfTagPolicyOutputReference
	_jsii_.Get(
		j,
		"lfTagPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) LfTagPolicyInput() *LakeformationPermissionsLfTagPolicy {
	var returns *LakeformationPermissionsLfTagPolicy
	_jsii_.Get(
		j,
		"lfTagPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Permissions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) PermissionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) PermissionsWithGrantOption() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) PermissionsWithGrantOptionInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permissionsWithGrantOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) Table() LakeformationPermissionsTableOutputReference {
	var returns LakeformationPermissionsTableOutputReference
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TableInput() *LakeformationPermissionsTable {
	var returns *LakeformationPermissionsTable
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TableWithColumns() LakeformationPermissionsTableWithColumnsOutputReference {
	var returns LakeformationPermissionsTableWithColumnsOutputReference
	_jsii_.Get(
		j,
		"tableWithColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TableWithColumnsInput() *LakeformationPermissionsTableWithColumns {
	var returns *LakeformationPermissionsTableWithColumns
	_jsii_.Get(
		j,
		"tableWithColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationPermissions) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/lakeformation_permissions aws_lakeformation_permissions} Resource.
func NewLakeformationPermissions(scope constructs.Construct, id *string, config *LakeformationPermissionsConfig) LakeformationPermissions {
	_init_.Initialize()

	if err := validateNewLakeformationPermissionsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LakeformationPermissions{}

	_jsii_.Create(
		"@cdktf/provider-aws.lakeformationPermissions.LakeformationPermissions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.54.1/docs/resources/lakeformation_permissions aws_lakeformation_permissions} Resource.
func NewLakeformationPermissions_Override(l LakeformationPermissions, scope constructs.Construct, id *string, config *LakeformationPermissionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lakeformationPermissions.LakeformationPermissions",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetCatalogId(val *string) {
	if err := j.validateSetCatalogIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetCatalogResource(val interface{}) {
	if err := j.validateSetCatalogResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogResource",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetPermissions(val *[]*string) {
	if err := j.validateSetPermissionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissions",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetPermissionsWithGrantOption(val *[]*string) {
	if err := j.validateSetPermissionsWithGrantOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionsWithGrantOption",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetPrincipal(val *string) {
	if err := j.validateSetPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LakeformationPermissions)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a LakeformationPermissions resource upon running "cdktf plan <stack-name>".
func LakeformationPermissions_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLakeformationPermissions_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lakeformationPermissions.LakeformationPermissions",
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
func LakeformationPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLakeformationPermissions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lakeformationPermissions.LakeformationPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LakeformationPermissions_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLakeformationPermissions_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lakeformationPermissions.LakeformationPermissions",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LakeformationPermissions_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLakeformationPermissions_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lakeformationPermissions.LakeformationPermissions",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LakeformationPermissions_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lakeformationPermissions.LakeformationPermissions",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LakeformationPermissions) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LakeformationPermissions) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LakeformationPermissions) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LakeformationPermissions) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LakeformationPermissions) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LakeformationPermissions) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LakeformationPermissions) PutDatabase(value *LakeformationPermissionsDatabase) {
	if err := l.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDatabase",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) PutDataCellsFilter(value *LakeformationPermissionsDataCellsFilter) {
	if err := l.validatePutDataCellsFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDataCellsFilter",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) PutDataLocation(value *LakeformationPermissionsDataLocation) {
	if err := l.validatePutDataLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDataLocation",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) PutLfTag(value *LakeformationPermissionsLfTag) {
	if err := l.validatePutLfTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLfTag",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) PutLfTagPolicy(value *LakeformationPermissionsLfTagPolicy) {
	if err := l.validatePutLfTagPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putLfTagPolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) PutTable(value *LakeformationPermissionsTable) {
	if err := l.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTable",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) PutTableWithColumns(value *LakeformationPermissionsTableWithColumns) {
	if err := l.validatePutTableWithColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTableWithColumns",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetCatalogId() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetCatalogResource() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalogResource",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetDatabase() {
	_jsii_.InvokeVoid(
		l,
		"resetDatabase",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetDataCellsFilter() {
	_jsii_.InvokeVoid(
		l,
		"resetDataCellsFilter",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetDataLocation() {
	_jsii_.InvokeVoid(
		l,
		"resetDataLocation",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetLfTag() {
	_jsii_.InvokeVoid(
		l,
		"resetLfTag",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetLfTagPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetLfTagPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetPermissionsWithGrantOption() {
	_jsii_.InvokeVoid(
		l,
		"resetPermissionsWithGrantOption",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetTable() {
	_jsii_.InvokeVoid(
		l,
		"resetTable",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) ResetTableWithColumns() {
	_jsii_.InvokeVoid(
		l,
		"resetTableWithColumns",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationPermissions) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationPermissions) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lakeformationdatalakesettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lakeformationdatalakesettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings}.
type LakeformationDataLakeSettings interface {
	cdktf.TerraformResource
	Admins() *[]*string
	SetAdmins(val *[]*string)
	AdminsInput() *[]*string
	AllowExternalDataFiltering() interface{}
	SetAllowExternalDataFiltering(val interface{})
	AllowExternalDataFilteringInput() interface{}
	AllowFullTableExternalDataAccess() interface{}
	SetAllowFullTableExternalDataAccess(val interface{})
	AllowFullTableExternalDataAccessInput() interface{}
	AuthorizedSessionTagValueList() *[]*string
	SetAuthorizedSessionTagValueList(val *[]*string)
	AuthorizedSessionTagValueListInput() *[]*string
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
	CreateDatabaseDefaultPermissions() LakeformationDataLakeSettingsCreateDatabaseDefaultPermissionsList
	CreateDatabaseDefaultPermissionsInput() interface{}
	CreateTableDefaultPermissions() LakeformationDataLakeSettingsCreateTableDefaultPermissionsList
	CreateTableDefaultPermissionsInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExternalDataFilteringAllowList() *[]*string
	SetExternalDataFilteringAllowList(val *[]*string)
	ExternalDataFilteringAllowListInput() *[]*string
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
	// The tree node.
	Node() constructs.Node
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
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
	ReadOnlyAdmins() *[]*string
	SetReadOnlyAdmins(val *[]*string)
	ReadOnlyAdminsInput() *[]*string
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TrustedResourceOwners() *[]*string
	SetTrustedResourceOwners(val *[]*string)
	TrustedResourceOwnersInput() *[]*string
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
	PutCreateDatabaseDefaultPermissions(value interface{})
	PutCreateTableDefaultPermissions(value interface{})
	ResetAdmins()
	ResetAllowExternalDataFiltering()
	ResetAllowFullTableExternalDataAccess()
	ResetAuthorizedSessionTagValueList()
	ResetCatalogId()
	ResetCreateDatabaseDefaultPermissions()
	ResetCreateTableDefaultPermissions()
	ResetExternalDataFilteringAllowList()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetReadOnlyAdmins()
	ResetRegion()
	ResetTrustedResourceOwners()
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

// The jsii proxy struct for LakeformationDataLakeSettings
type jsiiProxy_LakeformationDataLakeSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Admins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"admins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) AdminsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) AllowExternalDataFiltering() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExternalDataFiltering",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) AllowExternalDataFilteringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExternalDataFilteringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) AllowFullTableExternalDataAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFullTableExternalDataAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) AllowFullTableExternalDataAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowFullTableExternalDataAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) AuthorizedSessionTagValueList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedSessionTagValueList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) AuthorizedSessionTagValueListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"authorizedSessionTagValueListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CreateDatabaseDefaultPermissions() LakeformationDataLakeSettingsCreateDatabaseDefaultPermissionsList {
	var returns LakeformationDataLakeSettingsCreateDatabaseDefaultPermissionsList
	_jsii_.Get(
		j,
		"createDatabaseDefaultPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CreateDatabaseDefaultPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createDatabaseDefaultPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CreateTableDefaultPermissions() LakeformationDataLakeSettingsCreateTableDefaultPermissionsList {
	var returns LakeformationDataLakeSettingsCreateTableDefaultPermissionsList
	_jsii_.Get(
		j,
		"createTableDefaultPermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) CreateTableDefaultPermissionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createTableDefaultPermissionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) ExternalDataFilteringAllowList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalDataFilteringAllowList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) ExternalDataFilteringAllowListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"externalDataFilteringAllowListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) ReadOnlyAdmins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readOnlyAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) ReadOnlyAdminsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"readOnlyAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) TrustedResourceOwners() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedResourceOwners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LakeformationDataLakeSettings) TrustedResourceOwnersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedResourceOwnersInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings} Resource.
func NewLakeformationDataLakeSettings(scope constructs.Construct, id *string, config *LakeformationDataLakeSettingsConfig) LakeformationDataLakeSettings {
	_init_.Initialize()

	if err := validateNewLakeformationDataLakeSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LakeformationDataLakeSettings{}

	_jsii_.Create(
		"@cdktf/provider-aws.lakeformationDataLakeSettings.LakeformationDataLakeSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/lakeformation_data_lake_settings aws_lakeformation_data_lake_settings} Resource.
func NewLakeformationDataLakeSettings_Override(l LakeformationDataLakeSettings, scope constructs.Construct, id *string, config *LakeformationDataLakeSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lakeformationDataLakeSettings.LakeformationDataLakeSettings",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetAdmins(val *[]*string) {
	if err := j.validateSetAdminsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"admins",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetAllowExternalDataFiltering(val interface{}) {
	if err := j.validateSetAllowExternalDataFilteringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowExternalDataFiltering",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetAllowFullTableExternalDataAccess(val interface{}) {
	if err := j.validateSetAllowFullTableExternalDataAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowFullTableExternalDataAccess",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetAuthorizedSessionTagValueList(val *[]*string) {
	if err := j.validateSetAuthorizedSessionTagValueListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizedSessionTagValueList",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetCatalogId(val *string) {
	if err := j.validateSetCatalogIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetExternalDataFilteringAllowList(val *[]*string) {
	if err := j.validateSetExternalDataFilteringAllowListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalDataFilteringAllowList",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetReadOnlyAdmins(val *[]*string) {
	if err := j.validateSetReadOnlyAdminsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyAdmins",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_LakeformationDataLakeSettings)SetTrustedResourceOwners(val *[]*string) {
	if err := j.validateSetTrustedResourceOwnersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedResourceOwners",
		val,
	)
}

// Generates CDKTF code for importing a LakeformationDataLakeSettings resource upon running "cdktf plan <stack-name>".
func LakeformationDataLakeSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLakeformationDataLakeSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lakeformationDataLakeSettings.LakeformationDataLakeSettings",
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
func LakeformationDataLakeSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLakeformationDataLakeSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lakeformationDataLakeSettings.LakeformationDataLakeSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LakeformationDataLakeSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLakeformationDataLakeSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lakeformationDataLakeSettings.LakeformationDataLakeSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LakeformationDataLakeSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLakeformationDataLakeSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lakeformationDataLakeSettings.LakeformationDataLakeSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LakeformationDataLakeSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lakeformationDataLakeSettings.LakeformationDataLakeSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LakeformationDataLakeSettings) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LakeformationDataLakeSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LakeformationDataLakeSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LakeformationDataLakeSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LakeformationDataLakeSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LakeformationDataLakeSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LakeformationDataLakeSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LakeformationDataLakeSettings) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LakeformationDataLakeSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LakeformationDataLakeSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LakeformationDataLakeSettings) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) PutCreateDatabaseDefaultPermissions(value interface{}) {
	if err := l.validatePutCreateDatabaseDefaultPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCreateDatabaseDefaultPermissions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) PutCreateTableDefaultPermissions(value interface{}) {
	if err := l.validatePutCreateTableDefaultPermissionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCreateTableDefaultPermissions",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetAdmins() {
	_jsii_.InvokeVoid(
		l,
		"resetAdmins",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetAllowExternalDataFiltering() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowExternalDataFiltering",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetAllowFullTableExternalDataAccess() {
	_jsii_.InvokeVoid(
		l,
		"resetAllowFullTableExternalDataAccess",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetAuthorizedSessionTagValueList() {
	_jsii_.InvokeVoid(
		l,
		"resetAuthorizedSessionTagValueList",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetCatalogId() {
	_jsii_.InvokeVoid(
		l,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetCreateDatabaseDefaultPermissions() {
	_jsii_.InvokeVoid(
		l,
		"resetCreateDatabaseDefaultPermissions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetCreateTableDefaultPermissions() {
	_jsii_.InvokeVoid(
		l,
		"resetCreateTableDefaultPermissions",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetExternalDataFilteringAllowList() {
	_jsii_.InvokeVoid(
		l,
		"resetExternalDataFilteringAllowList",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetParameters() {
	_jsii_.InvokeVoid(
		l,
		"resetParameters",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetReadOnlyAdmins() {
	_jsii_.InvokeVoid(
		l,
		"resetReadOnlyAdmins",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetRegion() {
	_jsii_.InvokeVoid(
		l,
		"resetRegion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ResetTrustedResourceOwners() {
	_jsii_.InvokeVoid(
		l,
		"resetTrustedResourceOwners",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LakeformationDataLakeSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationDataLakeSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LakeformationDataLakeSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


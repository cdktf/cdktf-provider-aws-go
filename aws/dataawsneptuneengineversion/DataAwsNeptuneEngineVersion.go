// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsneptuneengineversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataawsneptuneengineversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/data-sources/neptune_engine_version aws_neptune_engine_version}.
type DataAwsNeptuneEngineVersion interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultCharacterSet() *string
	DefaultOnly() interface{}
	SetDefaultOnly(val interface{})
	DefaultOnlyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Engine() *string
	SetEngine(val *string)
	EngineDescription() *string
	EngineInput() *string
	ExportableLogTypes() *[]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HasMajorTarget() interface{}
	SetHasMajorTarget(val interface{})
	HasMajorTargetInput() interface{}
	HasMinorTarget() interface{}
	SetHasMinorTarget(val interface{})
	HasMinorTargetInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Latest() interface{}
	SetLatest(val interface{})
	LatestInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	ParameterGroupFamily() *string
	SetParameterGroupFamily(val *string)
	ParameterGroupFamilyInput() *string
	PreferredMajorTargets() *[]*string
	SetPreferredMajorTargets(val *[]*string)
	PreferredMajorTargetsInput() *[]*string
	PreferredUpgradeTargets() *[]*string
	SetPreferredUpgradeTargets(val *[]*string)
	PreferredUpgradeTargetsInput() *[]*string
	PreferredVersions() *[]*string
	SetPreferredVersions(val *[]*string)
	PreferredVersionsInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SupportedCharacterSets() *[]*string
	SupportedTimezones() *[]*string
	SupportsGlobalDatabases() cdktf.IResolvable
	SupportsLogExportsToCloudwatch() cdktf.IResolvable
	SupportsReadReplica() cdktf.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ValidMajorTargets() *[]*string
	ValidMinorTargets() *[]*string
	ValidUpgradeTargets() *[]*string
	Version() *string
	SetVersion(val *string)
	VersionActual() *string
	VersionDescription() *string
	VersionInput() *string
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
	ResetDefaultOnly()
	ResetEngine()
	ResetHasMajorTarget()
	ResetHasMinorTarget()
	ResetId()
	ResetLatest()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameterGroupFamily()
	ResetPreferredMajorTargets()
	ResetPreferredUpgradeTargets()
	ResetPreferredVersions()
	ResetRegion()
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
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

// The jsii proxy struct for DataAwsNeptuneEngineVersion
type jsiiProxy_DataAwsNeptuneEngineVersion struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) DefaultCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) DefaultOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) DefaultOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) EngineDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) ExportableLogTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exportableLogTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) HasMajorTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMajorTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) HasMajorTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMajorTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) HasMinorTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMinorTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) HasMinorTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMinorTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) Latest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) LatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) ParameterGroupFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) ParameterGroupFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) PreferredMajorTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredMajorTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) PreferredMajorTargetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredMajorTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) PreferredUpgradeTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredUpgradeTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) PreferredUpgradeTargetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredUpgradeTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) PreferredVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) PreferredVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) SupportedCharacterSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedCharacterSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) SupportedTimezones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTimezones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) SupportsGlobalDatabases() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportsGlobalDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) SupportsLogExportsToCloudwatch() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportsLogExportsToCloudwatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) SupportsReadReplica() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportsReadReplica",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) ValidMajorTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validMajorTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) ValidMinorTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validMinorTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) ValidUpgradeTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUpgradeTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) VersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/data-sources/neptune_engine_version aws_neptune_engine_version} Data Source.
func NewDataAwsNeptuneEngineVersion(scope constructs.Construct, id *string, config *DataAwsNeptuneEngineVersionConfig) DataAwsNeptuneEngineVersion {
	_init_.Initialize()

	if err := validateNewDataAwsNeptuneEngineVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsNeptuneEngineVersion{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsNeptuneEngineVersion.DataAwsNeptuneEngineVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.0.0/docs/data-sources/neptune_engine_version aws_neptune_engine_version} Data Source.
func NewDataAwsNeptuneEngineVersion_Override(d DataAwsNeptuneEngineVersion, scope constructs.Construct, id *string, config *DataAwsNeptuneEngineVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsNeptuneEngineVersion.DataAwsNeptuneEngineVersion",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetDefaultOnly(val interface{}) {
	if err := j.validateSetDefaultOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultOnly",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetHasMajorTarget(val interface{}) {
	if err := j.validateSetHasMajorTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasMajorTarget",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetHasMinorTarget(val interface{}) {
	if err := j.validateSetHasMinorTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasMinorTarget",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetLatest(val interface{}) {
	if err := j.validateSetLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latest",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetParameterGroupFamily(val *string) {
	if err := j.validateSetParameterGroupFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterGroupFamily",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetPreferredMajorTargets(val *[]*string) {
	if err := j.validateSetPreferredMajorTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMajorTargets",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetPreferredUpgradeTargets(val *[]*string) {
	if err := j.validateSetPreferredUpgradeTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredUpgradeTargets",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetPreferredVersions(val *[]*string) {
	if err := j.validateSetPreferredVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredVersions",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_DataAwsNeptuneEngineVersion)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsNeptuneEngineVersion resource upon running "cdktf plan <stack-name>".
func DataAwsNeptuneEngineVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsNeptuneEngineVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsNeptuneEngineVersion.DataAwsNeptuneEngineVersion",
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
func DataAwsNeptuneEngineVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsNeptuneEngineVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsNeptuneEngineVersion.DataAwsNeptuneEngineVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsNeptuneEngineVersion_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsNeptuneEngineVersion_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsNeptuneEngineVersion.DataAwsNeptuneEngineVersion",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsNeptuneEngineVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsNeptuneEngineVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsNeptuneEngineVersion.DataAwsNeptuneEngineVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsNeptuneEngineVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsNeptuneEngineVersion.DataAwsNeptuneEngineVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetDefaultOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetEngine() {
	_jsii_.InvokeVoid(
		d,
		"resetEngine",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetHasMajorTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetHasMajorTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetHasMinorTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetHasMinorTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetLatest() {
	_jsii_.InvokeVoid(
		d,
		"resetLatest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetParameterGroupFamily() {
	_jsii_.InvokeVoid(
		d,
		"resetParameterGroupFamily",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetPreferredMajorTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredMajorTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetPreferredUpgradeTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredUpgradeTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetPreferredVersions() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredVersions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsNeptuneEngineVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


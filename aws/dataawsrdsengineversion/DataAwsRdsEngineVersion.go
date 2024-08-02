// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsrdsengineversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataawsrdsengineversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/data-sources/rds_engine_version aws_rds_engine_version}.
type DataAwsRdsEngineVersion interface {
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
	Filter() DataAwsRdsEngineVersionFilterList
	FilterInput() interface{}
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
	IncludeAll() interface{}
	SetIncludeAll(val interface{})
	IncludeAllInput() interface{}
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
	Status() *string
	SupportedCharacterSets() *[]*string
	SupportedFeatureNames() *[]*string
	SupportedModes() *[]*string
	SupportedTimezones() *[]*string
	SupportsGlobalDatabases() cdktf.IResolvable
	SupportsLimitlessDatabase() cdktf.IResolvable
	SupportsLogExportsToCloudwatch() cdktf.IResolvable
	SupportsParallelQuery() cdktf.IResolvable
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
	PutFilter(value interface{})
	ResetDefaultOnly()
	ResetFilter()
	ResetHasMajorTarget()
	ResetHasMinorTarget()
	ResetId()
	ResetIncludeAll()
	ResetLatest()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameterGroupFamily()
	ResetPreferredMajorTargets()
	ResetPreferredUpgradeTargets()
	ResetPreferredVersions()
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

// The jsii proxy struct for DataAwsRdsEngineVersion
type jsiiProxy_DataAwsRdsEngineVersion struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) DefaultCharacterSet() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultCharacterSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) DefaultOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) DefaultOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) EngineDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) ExportableLogTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exportableLogTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) Filter() DataAwsRdsEngineVersionFilterList {
	var returns DataAwsRdsEngineVersionFilterList
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) FilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) HasMajorTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMajorTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) HasMajorTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMajorTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) HasMinorTarget() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMinorTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) HasMinorTargetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasMinorTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) IncludeAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) IncludeAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) Latest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) LatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"latestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) ParameterGroupFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) ParameterGroupFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) PreferredMajorTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredMajorTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) PreferredMajorTargetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredMajorTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) PreferredUpgradeTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredUpgradeTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) PreferredUpgradeTargetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredUpgradeTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) PreferredVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) PreferredVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) SupportedCharacterSets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedCharacterSets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) SupportedFeatureNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedFeatureNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) SupportedModes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedModes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) SupportedTimezones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"supportedTimezones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) SupportsGlobalDatabases() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportsGlobalDatabases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) SupportsLimitlessDatabase() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportsLimitlessDatabase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) SupportsLogExportsToCloudwatch() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportsLogExportsToCloudwatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) SupportsParallelQuery() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportsParallelQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) SupportsReadReplica() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportsReadReplica",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) ValidMajorTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validMajorTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) ValidMinorTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validMinorTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) ValidUpgradeTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUpgradeTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) VersionActual() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsRdsEngineVersion) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/data-sources/rds_engine_version aws_rds_engine_version} Data Source.
func NewDataAwsRdsEngineVersion(scope constructs.Construct, id *string, config *DataAwsRdsEngineVersionConfig) DataAwsRdsEngineVersion {
	_init_.Initialize()

	if err := validateNewDataAwsRdsEngineVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsRdsEngineVersion{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRdsEngineVersion.DataAwsRdsEngineVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.61.0/docs/data-sources/rds_engine_version aws_rds_engine_version} Data Source.
func NewDataAwsRdsEngineVersion_Override(d DataAwsRdsEngineVersion, scope constructs.Construct, id *string, config *DataAwsRdsEngineVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsRdsEngineVersion.DataAwsRdsEngineVersion",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetDefaultOnly(val interface{}) {
	if err := j.validateSetDefaultOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultOnly",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetHasMajorTarget(val interface{}) {
	if err := j.validateSetHasMajorTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasMajorTarget",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetHasMinorTarget(val interface{}) {
	if err := j.validateSetHasMinorTargetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasMinorTarget",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetIncludeAll(val interface{}) {
	if err := j.validateSetIncludeAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeAll",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetLatest(val interface{}) {
	if err := j.validateSetLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"latest",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetParameterGroupFamily(val *string) {
	if err := j.validateSetParameterGroupFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterGroupFamily",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetPreferredMajorTargets(val *[]*string) {
	if err := j.validateSetPreferredMajorTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMajorTargets",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetPreferredUpgradeTargets(val *[]*string) {
	if err := j.validateSetPreferredUpgradeTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredUpgradeTargets",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetPreferredVersions(val *[]*string) {
	if err := j.validateSetPreferredVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredVersions",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsRdsEngineVersion)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsRdsEngineVersion resource upon running "cdktf plan <stack-name>".
func DataAwsRdsEngineVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsRdsEngineVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRdsEngineVersion.DataAwsRdsEngineVersion",
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
func DataAwsRdsEngineVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRdsEngineVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRdsEngineVersion.DataAwsRdsEngineVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsRdsEngineVersion_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRdsEngineVersion_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRdsEngineVersion.DataAwsRdsEngineVersion",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsRdsEngineVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsRdsEngineVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsRdsEngineVersion.DataAwsRdsEngineVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsRdsEngineVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsRdsEngineVersion.DataAwsRdsEngineVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsRdsEngineVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsRdsEngineVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsRdsEngineVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsRdsEngineVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsRdsEngineVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsRdsEngineVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsRdsEngineVersion) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsRdsEngineVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsRdsEngineVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsRdsEngineVersion) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) PutFilter(value interface{}) {
	if err := d.validatePutFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFilter",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetDefaultOnly() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultOnly",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetFilter() {
	_jsii_.InvokeVoid(
		d,
		"resetFilter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetHasMajorTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetHasMajorTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetHasMinorTarget() {
	_jsii_.InvokeVoid(
		d,
		"resetHasMinorTarget",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetIncludeAll() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeAll",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetLatest() {
	_jsii_.InvokeVoid(
		d,
		"resetLatest",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetParameterGroupFamily() {
	_jsii_.InvokeVoid(
		d,
		"resetParameterGroupFamily",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetPreferredMajorTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredMajorTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetPreferredUpgradeTargets() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredUpgradeTargets",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetPreferredVersions() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredVersions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsRdsEngineVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


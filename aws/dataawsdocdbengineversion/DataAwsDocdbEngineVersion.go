// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsdocdbengineversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataawsdocdbengineversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/docdb_engine_version aws_docdb_engine_version}.
type DataAwsDocdbEngineVersion interface {
	cdktf.TerraformDataSource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	ParameterGroupFamily() *string
	SetParameterGroupFamily(val *string)
	ParameterGroupFamilyInput() *string
	PreferredVersions() *[]*string
	SetPreferredVersions(val *[]*string)
	PreferredVersionsInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SupportsLogExportsToCloudwatch() cdktf.IResolvable
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ValidUpgradeTargets() *[]*string
	Version() *string
	SetVersion(val *string)
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
	ResetEngine()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameterGroupFamily()
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

// The jsii proxy struct for DataAwsDocdbEngineVersion
type jsiiProxy_DataAwsDocdbEngineVersion struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) EngineDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) ExportableLogTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exportableLogTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) ParameterGroupFamily() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupFamily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) ParameterGroupFamilyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parameterGroupFamilyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) PreferredVersions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) PreferredVersionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"preferredVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) SupportsLogExportsToCloudwatch() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"supportsLogExportsToCloudwatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) ValidUpgradeTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"validUpgradeTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) VersionDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/docdb_engine_version aws_docdb_engine_version} Data Source.
func NewDataAwsDocdbEngineVersion(scope constructs.Construct, id *string, config *DataAwsDocdbEngineVersionConfig) DataAwsDocdbEngineVersion {
	_init_.Initialize()

	if err := validateNewDataAwsDocdbEngineVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsDocdbEngineVersion{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsDocdbEngineVersion.DataAwsDocdbEngineVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/data-sources/docdb_engine_version aws_docdb_engine_version} Data Source.
func NewDataAwsDocdbEngineVersion_Override(d DataAwsDocdbEngineVersion, scope constructs.Construct, id *string, config *DataAwsDocdbEngineVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsDocdbEngineVersion.DataAwsDocdbEngineVersion",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion)SetParameterGroupFamily(val *string) {
	if err := j.validateSetParameterGroupFamilyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameterGroupFamily",
		val,
	)
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion)SetPreferredVersions(val *[]*string) {
	if err := j.validateSetPreferredVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredVersions",
		val,
	)
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsDocdbEngineVersion)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsDocdbEngineVersion resource upon running "cdktf plan <stack-name>".
func DataAwsDocdbEngineVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsDocdbEngineVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsDocdbEngineVersion.DataAwsDocdbEngineVersion",
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
func DataAwsDocdbEngineVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsDocdbEngineVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsDocdbEngineVersion.DataAwsDocdbEngineVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsDocdbEngineVersion_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsDocdbEngineVersion_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsDocdbEngineVersion.DataAwsDocdbEngineVersion",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsDocdbEngineVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsDocdbEngineVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsDocdbEngineVersion.DataAwsDocdbEngineVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsDocdbEngineVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsDocdbEngineVersion.DataAwsDocdbEngineVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsDocdbEngineVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsDocdbEngineVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsDocdbEngineVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsDocdbEngineVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsDocdbEngineVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsDocdbEngineVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsDocdbEngineVersion) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsDocdbEngineVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsDocdbEngineVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsDocdbEngineVersion) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) ResetEngine() {
	_jsii_.InvokeVoid(
		d,
		"resetEngine",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) ResetParameterGroupFamily() {
	_jsii_.InvokeVoid(
		d,
		"resetParameterGroupFamily",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) ResetPreferredVersions() {
	_jsii_.InvokeVoid(
		d,
		"resetPreferredVersions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) ResetVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDocdbEngineVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


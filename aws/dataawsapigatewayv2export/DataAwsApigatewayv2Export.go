// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsapigatewayv2export

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/dataawsapigatewayv2export/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/data-sources/apigatewayv2_export aws_apigatewayv2_export}.
type DataAwsApigatewayv2Export interface {
	cdktf.TerraformDataSource
	ApiId() *string
	SetApiId(val *string)
	ApiIdInput() *string
	Body() *string
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
	ExportVersion() *string
	SetExportVersion(val *string)
	ExportVersionInput() *string
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
	IncludeExtensions() interface{}
	SetIncludeExtensions(val interface{})
	IncludeExtensionsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OutputType() *string
	SetOutputType(val *string)
	OutputTypeInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Specification() *string
	SetSpecification(val *string)
	SpecificationInput() *string
	StageName() *string
	SetStageName(val *string)
	StageNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetExportVersion()
	ResetId()
	ResetIncludeExtensions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStageName()
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

// The jsii proxy struct for DataAwsApigatewayv2Export
type jsiiProxy_DataAwsApigatewayv2Export struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) ApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) ApiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) ExportVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) ExportVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) IncludeExtensions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExtensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) IncludeExtensionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeExtensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) Specification() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) SpecificationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"specificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) StageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) StageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsApigatewayv2Export) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/data-sources/apigatewayv2_export aws_apigatewayv2_export} Data Source.
func NewDataAwsApigatewayv2Export(scope constructs.Construct, id *string, config *DataAwsApigatewayv2ExportConfig) DataAwsApigatewayv2Export {
	_init_.Initialize()

	if err := validateNewDataAwsApigatewayv2ExportParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsApigatewayv2Export{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsApigatewayv2Export.DataAwsApigatewayv2Export",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/data-sources/apigatewayv2_export aws_apigatewayv2_export} Data Source.
func NewDataAwsApigatewayv2Export_Override(d DataAwsApigatewayv2Export, scope constructs.Construct, id *string, config *DataAwsApigatewayv2ExportConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsApigatewayv2Export.DataAwsApigatewayv2Export",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetApiId(val *string) {
	if err := j.validateSetApiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiId",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetExportVersion(val *string) {
	if err := j.validateSetExportVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportVersion",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetIncludeExtensions(val interface{}) {
	if err := j.validateSetIncludeExtensionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeExtensions",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetOutputType(val *string) {
	if err := j.validateSetOutputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetSpecification(val *string) {
	if err := j.validateSetSpecificationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"specification",
		val,
	)
}

func (j *jsiiProxy_DataAwsApigatewayv2Export)SetStageName(val *string) {
	if err := j.validateSetStageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stageName",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsApigatewayv2Export resource upon running "cdktf plan <stack-name>".
func DataAwsApigatewayv2Export_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsApigatewayv2Export_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsApigatewayv2Export.DataAwsApigatewayv2Export",
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
func DataAwsApigatewayv2Export_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsApigatewayv2Export_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsApigatewayv2Export.DataAwsApigatewayv2Export",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsApigatewayv2Export_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsApigatewayv2Export_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsApigatewayv2Export.DataAwsApigatewayv2Export",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsApigatewayv2Export_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsApigatewayv2Export_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsApigatewayv2Export.DataAwsApigatewayv2Export",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsApigatewayv2Export_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsApigatewayv2Export.DataAwsApigatewayv2Export",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsApigatewayv2Export) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsApigatewayv2Export) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsApigatewayv2Export) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsApigatewayv2Export) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsApigatewayv2Export) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsApigatewayv2Export) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsApigatewayv2Export) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsApigatewayv2Export) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsApigatewayv2Export) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsApigatewayv2Export) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ResetExportVersion() {
	_jsii_.InvokeVoid(
		d,
		"resetExportVersion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ResetIncludeExtensions() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludeExtensions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ResetStageName() {
	_jsii_.InvokeVoid(
		d,
		"resetStageName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsApigatewayv2Export) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


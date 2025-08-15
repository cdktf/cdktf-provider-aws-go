// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsbedrockfoundationmodels

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/dataawsbedrockfoundationmodels/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/bedrock_foundation_models aws_bedrock_foundation_models}.
type DataAwsBedrockFoundationModels interface {
	cdktf.TerraformDataSource
	ByCustomizationType() *string
	SetByCustomizationType(val *string)
	ByCustomizationTypeInput() *string
	ByInferenceType() *string
	SetByInferenceType(val *string)
	ByInferenceTypeInput() *string
	ByOutputModality() *string
	SetByOutputModality(val *string)
	ByOutputModalityInput() *string
	ByProvider() *string
	SetByProvider(val *string)
	ByProviderInput() *string
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ModelSummaries() DataAwsBedrockFoundationModelsModelSummariesList
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
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
	ResetByCustomizationType()
	ResetByInferenceType()
	ResetByOutputModality()
	ResetByProvider()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
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

// The jsii proxy struct for DataAwsBedrockFoundationModels
type jsiiProxy_DataAwsBedrockFoundationModels struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) ByCustomizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"byCustomizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) ByCustomizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"byCustomizationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) ByInferenceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"byInferenceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) ByInferenceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"byInferenceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) ByOutputModality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"byOutputModality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) ByOutputModalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"byOutputModalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) ByProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"byProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) ByProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"byProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) ModelSummaries() DataAwsBedrockFoundationModelsModelSummariesList {
	var returns DataAwsBedrockFoundationModelsModelSummariesList
	_jsii_.Get(
		j,
		"modelSummaries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/bedrock_foundation_models aws_bedrock_foundation_models} Data Source.
func NewDataAwsBedrockFoundationModels(scope constructs.Construct, id *string, config *DataAwsBedrockFoundationModelsConfig) DataAwsBedrockFoundationModels {
	_init_.Initialize()

	if err := validateNewDataAwsBedrockFoundationModelsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsBedrockFoundationModels{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsBedrockFoundationModels.DataAwsBedrockFoundationModels",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/data-sources/bedrock_foundation_models aws_bedrock_foundation_models} Data Source.
func NewDataAwsBedrockFoundationModels_Override(d DataAwsBedrockFoundationModels, scope constructs.Construct, id *string, config *DataAwsBedrockFoundationModelsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsBedrockFoundationModels.DataAwsBedrockFoundationModels",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels)SetByCustomizationType(val *string) {
	if err := j.validateSetByCustomizationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"byCustomizationType",
		val,
	)
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels)SetByInferenceType(val *string) {
	if err := j.validateSetByInferenceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"byInferenceType",
		val,
	)
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels)SetByOutputModality(val *string) {
	if err := j.validateSetByOutputModalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"byOutputModality",
		val,
	)
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels)SetByProvider(val *string) {
	if err := j.validateSetByProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"byProvider",
		val,
	)
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsBedrockFoundationModels)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsBedrockFoundationModels resource upon running "cdktf plan <stack-name>".
func DataAwsBedrockFoundationModels_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsBedrockFoundationModels_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsBedrockFoundationModels.DataAwsBedrockFoundationModels",
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
func DataAwsBedrockFoundationModels_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsBedrockFoundationModels_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsBedrockFoundationModels.DataAwsBedrockFoundationModels",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsBedrockFoundationModels_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsBedrockFoundationModels_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsBedrockFoundationModels.DataAwsBedrockFoundationModels",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsBedrockFoundationModels_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsBedrockFoundationModels_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsBedrockFoundationModels.DataAwsBedrockFoundationModels",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsBedrockFoundationModels_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsBedrockFoundationModels.DataAwsBedrockFoundationModels",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsBedrockFoundationModels) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsBedrockFoundationModels) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsBedrockFoundationModels) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsBedrockFoundationModels) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsBedrockFoundationModels) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsBedrockFoundationModels) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsBedrockFoundationModels) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsBedrockFoundationModels) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsBedrockFoundationModels) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsBedrockFoundationModels) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) ResetByCustomizationType() {
	_jsii_.InvokeVoid(
		d,
		"resetByCustomizationType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) ResetByInferenceType() {
	_jsii_.InvokeVoid(
		d,
		"resetByInferenceType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) ResetByOutputModality() {
	_jsii_.InvokeVoid(
		d,
		"resetByOutputModality",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) ResetByProvider() {
	_jsii_.InvokeVoid(
		d,
		"resetByProvider",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) ResetRegion() {
	_jsii_.InvokeVoid(
		d,
		"resetRegion",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsBedrockFoundationModels) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


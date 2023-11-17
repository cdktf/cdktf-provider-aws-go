// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsmqbrokerinstancetypeofferings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/dataawsmqbrokerinstancetypeofferings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/mq_broker_instance_type_offerings aws_mq_broker_instance_type_offerings}.
type DataAwsMqBrokerInstanceTypeOfferings interface {
	cdktf.TerraformDataSource
	BrokerInstanceOptions() DataAwsMqBrokerInstanceTypeOfferingsBrokerInstanceOptionsList
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
	EngineType() *string
	SetEngineType(val *string)
	EngineTypeInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HostInstanceType() *string
	SetHostInstanceType(val *string)
	HostInstanceTypeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
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
	ResetEngineType()
	ResetHostInstanceType()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStorageType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsMqBrokerInstanceTypeOfferings
type jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) BrokerInstanceOptions() DataAwsMqBrokerInstanceTypeOfferingsBrokerInstanceOptionsList {
	var returns DataAwsMqBrokerInstanceTypeOfferingsBrokerInstanceOptionsList
	_jsii_.Get(
		j,
		"brokerInstanceOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) EngineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) EngineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) HostInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) HostInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/mq_broker_instance_type_offerings aws_mq_broker_instance_type_offerings} Data Source.
func NewDataAwsMqBrokerInstanceTypeOfferings(scope constructs.Construct, id *string, config *DataAwsMqBrokerInstanceTypeOfferingsConfig) DataAwsMqBrokerInstanceTypeOfferings {
	_init_.Initialize()

	if err := validateNewDataAwsMqBrokerInstanceTypeOfferingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings{}

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsMqBrokerInstanceTypeOfferings.DataAwsMqBrokerInstanceTypeOfferings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/data-sources/mq_broker_instance_type_offerings aws_mq_broker_instance_type_offerings} Data Source.
func NewDataAwsMqBrokerInstanceTypeOfferings_Override(d DataAwsMqBrokerInstanceTypeOfferings, scope constructs.Construct, id *string, config *DataAwsMqBrokerInstanceTypeOfferingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.dataAwsMqBrokerInstanceTypeOfferings.DataAwsMqBrokerInstanceTypeOfferings",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings)SetEngineType(val *string) {
	if err := j.validateSetEngineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineType",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings)SetHostInstanceType(val *string) {
	if err := j.validateSetHostInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostInstanceType",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsMqBrokerInstanceTypeOfferings resource upon running "cdktf plan <stack-name>".
func DataAwsMqBrokerInstanceTypeOfferings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsMqBrokerInstanceTypeOfferings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsMqBrokerInstanceTypeOfferings.DataAwsMqBrokerInstanceTypeOfferings",
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
func DataAwsMqBrokerInstanceTypeOfferings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsMqBrokerInstanceTypeOfferings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsMqBrokerInstanceTypeOfferings.DataAwsMqBrokerInstanceTypeOfferings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsMqBrokerInstanceTypeOfferings_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsMqBrokerInstanceTypeOfferings_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsMqBrokerInstanceTypeOfferings.DataAwsMqBrokerInstanceTypeOfferings",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsMqBrokerInstanceTypeOfferings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsMqBrokerInstanceTypeOfferings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.dataAwsMqBrokerInstanceTypeOfferings.DataAwsMqBrokerInstanceTypeOfferings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsMqBrokerInstanceTypeOfferings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.dataAwsMqBrokerInstanceTypeOfferings.DataAwsMqBrokerInstanceTypeOfferings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) ResetEngineType() {
	_jsii_.InvokeVoid(
		d,
		"resetEngineType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) ResetHostInstanceType() {
	_jsii_.InvokeVoid(
		d,
		"resetHostInstanceType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) ResetStorageType() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsMqBrokerInstanceTypeOfferings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


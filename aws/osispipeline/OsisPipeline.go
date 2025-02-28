// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osispipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/osispipeline/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/osis_pipeline aws_osis_pipeline}.
type OsisPipeline interface {
	cdktf.TerraformResource
	BufferOptions() OsisPipelineBufferOptionsList
	BufferOptionsInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EncryptionAtRestOptions() OsisPipelineEncryptionAtRestOptionsList
	EncryptionAtRestOptionsInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	IngestEndpointUrls() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogPublishingOptions() OsisPipelineLogPublishingOptionsList
	LogPublishingOptionsInput() interface{}
	MaxUnits() *float64
	SetMaxUnits(val *float64)
	MaxUnitsInput() *float64
	MinUnits() *float64
	SetMinUnits(val *float64)
	MinUnitsInput() *float64
	// The tree node.
	Node() constructs.Node
	PipelineArn() *string
	PipelineConfigurationBody() *string
	SetPipelineConfigurationBody(val *string)
	PipelineConfigurationBodyInput() *string
	PipelineName() *string
	SetPipelineName(val *string)
	PipelineNameInput() *string
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() cdktf.StringMap
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() OsisPipelineTimeoutsOutputReference
	TimeoutsInput() interface{}
	VpcOptions() OsisPipelineVpcOptionsList
	VpcOptionsInput() interface{}
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
	PutBufferOptions(value interface{})
	PutEncryptionAtRestOptions(value interface{})
	PutLogPublishingOptions(value interface{})
	PutTimeouts(value *OsisPipelineTimeouts)
	PutVpcOptions(value interface{})
	ResetBufferOptions()
	ResetEncryptionAtRestOptions()
	ResetLogPublishingOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	ResetTimeouts()
	ResetVpcOptions()
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

// The jsii proxy struct for OsisPipeline
type jsiiProxy_OsisPipeline struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OsisPipeline) BufferOptions() OsisPipelineBufferOptionsList {
	var returns OsisPipelineBufferOptionsList
	_jsii_.Get(
		j,
		"bufferOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) BufferOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bufferOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) EncryptionAtRestOptions() OsisPipelineEncryptionAtRestOptionsList {
	var returns OsisPipelineEncryptionAtRestOptionsList
	_jsii_.Get(
		j,
		"encryptionAtRestOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) EncryptionAtRestOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtRestOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) IngestEndpointUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ingestEndpointUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) LogPublishingOptions() OsisPipelineLogPublishingOptionsList {
	var returns OsisPipelineLogPublishingOptionsList
	_jsii_.Get(
		j,
		"logPublishingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) LogPublishingOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logPublishingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) MaxUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) MaxUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) MinUnits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUnits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) MinUnitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUnitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) PipelineArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) PipelineConfigurationBody() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineConfigurationBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) PipelineConfigurationBodyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineConfigurationBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) PipelineName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) PipelineNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pipelineNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) Timeouts() OsisPipelineTimeoutsOutputReference {
	var returns OsisPipelineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) VpcOptions() OsisPipelineVpcOptionsList {
	var returns OsisPipelineVpcOptionsList
	_jsii_.Get(
		j,
		"vpcOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OsisPipeline) VpcOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcOptionsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/osis_pipeline aws_osis_pipeline} Resource.
func NewOsisPipeline(scope constructs.Construct, id *string, config *OsisPipelineConfig) OsisPipeline {
	_init_.Initialize()

	if err := validateNewOsisPipelineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OsisPipeline{}

	_jsii_.Create(
		"@cdktf/provider-aws.osisPipeline.OsisPipeline",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/osis_pipeline aws_osis_pipeline} Resource.
func NewOsisPipeline_Override(o OsisPipeline, scope constructs.Construct, id *string, config *OsisPipelineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.osisPipeline.OsisPipeline",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OsisPipeline)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OsisPipeline)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OsisPipeline)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OsisPipeline)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OsisPipeline)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OsisPipeline)SetMaxUnits(val *float64) {
	if err := j.validateSetMaxUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnits",
		val,
	)
}

func (j *jsiiProxy_OsisPipeline)SetMinUnits(val *float64) {
	if err := j.validateSetMinUnitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minUnits",
		val,
	)
}

func (j *jsiiProxy_OsisPipeline)SetPipelineConfigurationBody(val *string) {
	if err := j.validateSetPipelineConfigurationBodyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineConfigurationBody",
		val,
	)
}

func (j *jsiiProxy_OsisPipeline)SetPipelineName(val *string) {
	if err := j.validateSetPipelineNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pipelineName",
		val,
	)
}

func (j *jsiiProxy_OsisPipeline)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OsisPipeline)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OsisPipeline)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a OsisPipeline resource upon running "cdktf plan <stack-name>".
func OsisPipeline_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateOsisPipeline_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.osisPipeline.OsisPipeline",
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
func OsisPipeline_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOsisPipeline_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.osisPipeline.OsisPipeline",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OsisPipeline_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOsisPipeline_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.osisPipeline.OsisPipeline",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OsisPipeline_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOsisPipeline_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.osisPipeline.OsisPipeline",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OsisPipeline_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.osisPipeline.OsisPipeline",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OsisPipeline) AddMoveTarget(moveTarget *string) {
	if err := o.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (o *jsiiProxy_OsisPipeline) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OsisPipeline) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := o.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (o *jsiiProxy_OsisPipeline) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) MoveFromId(id *string) {
	if err := o.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveFromId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OsisPipeline) MoveTo(moveTarget *string, index interface{}) {
	if err := o.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (o *jsiiProxy_OsisPipeline) MoveToId(id *string) {
	if err := o.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"moveToId",
		[]interface{}{id},
	)
}

func (o *jsiiProxy_OsisPipeline) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OsisPipeline) PutBufferOptions(value interface{}) {
	if err := o.validatePutBufferOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBufferOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsisPipeline) PutEncryptionAtRestOptions(value interface{}) {
	if err := o.validatePutEncryptionAtRestOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putEncryptionAtRestOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsisPipeline) PutLogPublishingOptions(value interface{}) {
	if err := o.validatePutLogPublishingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putLogPublishingOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsisPipeline) PutTimeouts(value *OsisPipelineTimeouts) {
	if err := o.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsisPipeline) PutVpcOptions(value interface{}) {
	if err := o.validatePutVpcOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putVpcOptions",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OsisPipeline) ResetBufferOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetBufferOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsisPipeline) ResetEncryptionAtRestOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetEncryptionAtRestOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsisPipeline) ResetLogPublishingOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetLogPublishingOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsisPipeline) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsisPipeline) ResetTags() {
	_jsii_.InvokeVoid(
		o,
		"resetTags",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsisPipeline) ResetTimeouts() {
	_jsii_.InvokeVoid(
		o,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsisPipeline) ResetVpcOptions() {
	_jsii_.InvokeVoid(
		o,
		"resetVpcOptions",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OsisPipeline) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OsisPipeline) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


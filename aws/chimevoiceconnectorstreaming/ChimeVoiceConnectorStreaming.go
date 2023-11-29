// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimevoiceconnectorstreaming

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/chimevoiceconnectorstreaming/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/chime_voice_connector_streaming aws_chime_voice_connector_streaming}.
type ChimeVoiceConnectorStreaming interface {
	cdktf.TerraformResource
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
	DataRetention() *float64
	SetDataRetention(val *float64)
	DataRetentionInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disabled() interface{}
	SetDisabled(val interface{})
	DisabledInput() interface{}
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
	MediaInsightsConfiguration() ChimeVoiceConnectorStreamingMediaInsightsConfigurationOutputReference
	MediaInsightsConfigurationInput() *ChimeVoiceConnectorStreamingMediaInsightsConfiguration
	// The tree node.
	Node() constructs.Node
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
	StreamingNotificationTargets() *[]*string
	SetStreamingNotificationTargets(val *[]*string)
	StreamingNotificationTargetsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VoiceConnectorId() *string
	SetVoiceConnectorId(val *string)
	VoiceConnectorIdInput() *string
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
	PutMediaInsightsConfiguration(value *ChimeVoiceConnectorStreamingMediaInsightsConfiguration)
	ResetDisabled()
	ResetId()
	ResetMediaInsightsConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStreamingNotificationTargets()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ChimeVoiceConnectorStreaming
type jsiiProxy_ChimeVoiceConnectorStreaming struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) DataRetention() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) DataRetentionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Disabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) DisabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) MediaInsightsConfiguration() ChimeVoiceConnectorStreamingMediaInsightsConfigurationOutputReference {
	var returns ChimeVoiceConnectorStreamingMediaInsightsConfigurationOutputReference
	_jsii_.Get(
		j,
		"mediaInsightsConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) MediaInsightsConfigurationInput() *ChimeVoiceConnectorStreamingMediaInsightsConfiguration {
	var returns *ChimeVoiceConnectorStreamingMediaInsightsConfiguration
	_jsii_.Get(
		j,
		"mediaInsightsConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) StreamingNotificationTargets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamingNotificationTargets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) StreamingNotificationTargetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamingNotificationTargetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) VoiceConnectorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming) VoiceConnectorIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceConnectorIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/chime_voice_connector_streaming aws_chime_voice_connector_streaming} Resource.
func NewChimeVoiceConnectorStreaming(scope constructs.Construct, id *string, config *ChimeVoiceConnectorStreamingConfig) ChimeVoiceConnectorStreaming {
	_init_.Initialize()

	if err := validateNewChimeVoiceConnectorStreamingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ChimeVoiceConnectorStreaming{}

	_jsii_.Create(
		"@cdktf/provider-aws.chimeVoiceConnectorStreaming.ChimeVoiceConnectorStreaming",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/chime_voice_connector_streaming aws_chime_voice_connector_streaming} Resource.
func NewChimeVoiceConnectorStreaming_Override(c ChimeVoiceConnectorStreaming, scope constructs.Construct, id *string, config *ChimeVoiceConnectorStreamingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.chimeVoiceConnectorStreaming.ChimeVoiceConnectorStreaming",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetDataRetention(val *float64) {
	if err := j.validateSetDataRetentionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataRetention",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetDisabled(val interface{}) {
	if err := j.validateSetDisabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabled",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetStreamingNotificationTargets(val *[]*string) {
	if err := j.validateSetStreamingNotificationTargetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streamingNotificationTargets",
		val,
	)
}

func (j *jsiiProxy_ChimeVoiceConnectorStreaming)SetVoiceConnectorId(val *string) {
	if err := j.validateSetVoiceConnectorIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"voiceConnectorId",
		val,
	)
}

// Generates CDKTF code for importing a ChimeVoiceConnectorStreaming resource upon running "cdktf plan <stack-name>".
func ChimeVoiceConnectorStreaming_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateChimeVoiceConnectorStreaming_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.chimeVoiceConnectorStreaming.ChimeVoiceConnectorStreaming",
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
func ChimeVoiceConnectorStreaming_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChimeVoiceConnectorStreaming_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.chimeVoiceConnectorStreaming.ChimeVoiceConnectorStreaming",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ChimeVoiceConnectorStreaming_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChimeVoiceConnectorStreaming_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.chimeVoiceConnectorStreaming.ChimeVoiceConnectorStreaming",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ChimeVoiceConnectorStreaming_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateChimeVoiceConnectorStreaming_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.chimeVoiceConnectorStreaming.ChimeVoiceConnectorStreaming",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ChimeVoiceConnectorStreaming_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.chimeVoiceConnectorStreaming.ChimeVoiceConnectorStreaming",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) PutMediaInsightsConfiguration(value *ChimeVoiceConnectorStreamingMediaInsightsConfiguration) {
	if err := c.validatePutMediaInsightsConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMediaInsightsConfiguration",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ResetDisabled() {
	_jsii_.InvokeVoid(
		c,
		"resetDisabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ResetMediaInsightsConfiguration() {
	_jsii_.InvokeVoid(
		c,
		"resetMediaInsightsConfiguration",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ResetStreamingNotificationTargets() {
	_jsii_.InvokeVoid(
		c,
		"resetStreamingNotificationTargets",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ChimeVoiceConnectorStreaming) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


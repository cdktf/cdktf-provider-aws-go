// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagentcollaborator

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/bedrockagentagentcollaborator/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/bedrockagent_agent_collaborator aws_bedrockagent_agent_collaborator}.
type BedrockagentAgentCollaborator interface {
	cdktf.TerraformResource
	AgentDescriptor() BedrockagentAgentCollaboratorAgentDescriptorList
	AgentDescriptorInput() interface{}
	AgentId() *string
	SetAgentId(val *string)
	AgentIdInput() *string
	AgentVersion() *string
	SetAgentVersion(val *string)
	AgentVersionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CollaborationInstruction() *string
	SetCollaborationInstruction(val *string)
	CollaborationInstructionInput() *string
	CollaboratorId() *string
	CollaboratorName() *string
	SetCollaboratorName(val *string)
	CollaboratorNameInput() *string
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
	// The tree node.
	Node() constructs.Node
	PrepareAgent() interface{}
	SetPrepareAgent(val interface{})
	PrepareAgentInput() interface{}
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RelayConversationHistory() *string
	SetRelayConversationHistory(val *string)
	RelayConversationHistoryInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() BedrockagentAgentCollaboratorTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutAgentDescriptor(value interface{})
	PutTimeouts(value *BedrockagentAgentCollaboratorTimeouts)
	ResetAgentDescriptor()
	ResetAgentVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrepareAgent()
	ResetRegion()
	ResetRelayConversationHistory()
	ResetTimeouts()
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

// The jsii proxy struct for BedrockagentAgentCollaborator
type jsiiProxy_BedrockagentAgentCollaborator struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) AgentDescriptor() BedrockagentAgentCollaboratorAgentDescriptorList {
	var returns BedrockagentAgentCollaboratorAgentDescriptorList
	_jsii_.Get(
		j,
		"agentDescriptor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) AgentDescriptorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentDescriptorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) AgentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) AgentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) CollaborationInstruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaborationInstruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) CollaborationInstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaborationInstructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) CollaboratorId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) CollaboratorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) CollaboratorNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"collaboratorNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) PrepareAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prepareAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) PrepareAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prepareAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) RelayConversationHistory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relayConversationHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) RelayConversationHistoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relayConversationHistoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) Timeouts() BedrockagentAgentCollaboratorTimeoutsOutputReference {
	var returns BedrockagentAgentCollaboratorTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentCollaborator) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/bedrockagent_agent_collaborator aws_bedrockagent_agent_collaborator} Resource.
func NewBedrockagentAgentCollaborator(scope constructs.Construct, id *string, config *BedrockagentAgentCollaboratorConfig) BedrockagentAgentCollaborator {
	_init_.Initialize()

	if err := validateNewBedrockagentAgentCollaboratorParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentAgentCollaborator{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentAgentCollaborator.BedrockagentAgentCollaborator",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/bedrockagent_agent_collaborator aws_bedrockagent_agent_collaborator} Resource.
func NewBedrockagentAgentCollaborator_Override(b BedrockagentAgentCollaborator, scope constructs.Construct, id *string, config *BedrockagentAgentCollaboratorConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentAgentCollaborator.BedrockagentAgentCollaborator",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetAgentId(val *string) {
	if err := j.validateSetAgentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentId",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetAgentVersion(val *string) {
	if err := j.validateSetAgentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetCollaborationInstruction(val *string) {
	if err := j.validateSetCollaborationInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collaborationInstruction",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetCollaboratorName(val *string) {
	if err := j.validateSetCollaboratorNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"collaboratorName",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetPrepareAgent(val interface{}) {
	if err := j.validateSetPrepareAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prepareAgent",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentCollaborator)SetRelayConversationHistory(val *string) {
	if err := j.validateSetRelayConversationHistoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"relayConversationHistory",
		val,
	)
}

// Generates CDKTF code for importing a BedrockagentAgentCollaborator resource upon running "cdktf plan <stack-name>".
func BedrockagentAgentCollaborator_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBedrockagentAgentCollaborator_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgentCollaborator.BedrockagentAgentCollaborator",
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
func BedrockagentAgentCollaborator_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentAgentCollaborator_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgentCollaborator.BedrockagentAgentCollaborator",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentAgentCollaborator_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentAgentCollaborator_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgentCollaborator.BedrockagentAgentCollaborator",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentAgentCollaborator_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentAgentCollaborator_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgentCollaborator.BedrockagentAgentCollaborator",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BedrockagentAgentCollaborator_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.bedrockagentAgentCollaborator.BedrockagentAgentCollaborator",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) PutAgentDescriptor(value interface{}) {
	if err := b.validatePutAgentDescriptorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putAgentDescriptor",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) PutTimeouts(value *BedrockagentAgentCollaboratorTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ResetAgentDescriptor() {
	_jsii_.InvokeVoid(
		b,
		"resetAgentDescriptor",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ResetAgentVersion() {
	_jsii_.InvokeVoid(
		b,
		"resetAgentVersion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ResetPrepareAgent() {
	_jsii_.InvokeVoid(
		b,
		"resetPrepareAgent",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ResetRegion() {
	_jsii_.InvokeVoid(
		b,
		"resetRegion",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ResetRelayConversationHistory() {
	_jsii_.InvokeVoid(
		b,
		"resetRelayConversationHistory",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentCollaborator) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


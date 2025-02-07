// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/bedrockagentagent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/bedrockagent_agent aws_bedrockagent_agent}.
type BedrockagentAgent interface {
	cdktf.TerraformResource
	AgentArn() *string
	AgentCollaboration() *string
	SetAgentCollaboration(val *string)
	AgentCollaborationInput() *string
	AgentId() *string
	AgentName() *string
	SetAgentName(val *string)
	AgentNameInput() *string
	AgentResourceRoleArn() *string
	SetAgentResourceRoleArn(val *string)
	AgentResourceRoleArnInput() *string
	AgentVersion() *string
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
	CustomerEncryptionKeyArn() *string
	SetCustomerEncryptionKeyArn(val *string)
	CustomerEncryptionKeyArnInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	FoundationModel() *string
	SetFoundationModel(val *string)
	FoundationModelInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GuardrailConfiguration() BedrockagentAgentGuardrailConfigurationList
	GuardrailConfigurationInput() interface{}
	Id() *string
	IdleSessionTtlInSeconds() *float64
	SetIdleSessionTtlInSeconds(val *float64)
	IdleSessionTtlInSecondsInput() *float64
	Instruction() *string
	SetInstruction(val *string)
	InstructionInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PrepareAgent() interface{}
	SetPrepareAgent(val interface{})
	PrepareAgentInput() interface{}
	PromptOverrideConfiguration() BedrockagentAgentPromptOverrideConfigurationList
	PromptOverrideConfigurationInput() interface{}
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
	SkipResourceInUseCheck() interface{}
	SetSkipResourceInUseCheck(val interface{})
	SkipResourceInUseCheckInput() interface{}
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
	Timeouts() BedrockagentAgentTimeoutsOutputReference
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
	PutGuardrailConfiguration(value interface{})
	PutPromptOverrideConfiguration(value interface{})
	PutTimeouts(value *BedrockagentAgentTimeouts)
	ResetAgentCollaboration()
	ResetCustomerEncryptionKeyArn()
	ResetDescription()
	ResetGuardrailConfiguration()
	ResetIdleSessionTtlInSeconds()
	ResetInstruction()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrepareAgent()
	ResetPromptOverrideConfiguration()
	ResetSkipResourceInUseCheck()
	ResetTags()
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

// The jsii proxy struct for BedrockagentAgent
type jsiiProxy_BedrockagentAgent struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BedrockagentAgent) AgentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) AgentCollaboration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentCollaboration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) AgentCollaborationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentCollaborationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) AgentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) AgentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) AgentResourceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentResourceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) AgentResourceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentResourceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) CustomerEncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEncryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) CustomerEncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEncryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) FoundationModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foundationModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) FoundationModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foundationModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) GuardrailConfiguration() BedrockagentAgentGuardrailConfigurationList {
	var returns BedrockagentAgentGuardrailConfigurationList
	_jsii_.Get(
		j,
		"guardrailConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) GuardrailConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guardrailConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) IdleSessionTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) IdleSessionTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) InstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) PrepareAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prepareAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) PrepareAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prepareAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) PromptOverrideConfiguration() BedrockagentAgentPromptOverrideConfigurationList {
	var returns BedrockagentAgentPromptOverrideConfigurationList
	_jsii_.Get(
		j,
		"promptOverrideConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) PromptOverrideConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptOverrideConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) SkipResourceInUseCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) SkipResourceInUseCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) TagsAll() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) Timeouts() BedrockagentAgentTimeoutsOutputReference {
	var returns BedrockagentAgentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/bedrockagent_agent aws_bedrockagent_agent} Resource.
func NewBedrockagentAgent(scope constructs.Construct, id *string, config *BedrockagentAgentConfig) BedrockagentAgent {
	_init_.Initialize()

	if err := validateNewBedrockagentAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentAgent{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentAgent.BedrockagentAgent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.86.0/docs/resources/bedrockagent_agent aws_bedrockagent_agent} Resource.
func NewBedrockagentAgent_Override(b BedrockagentAgent, scope constructs.Construct, id *string, config *BedrockagentAgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentAgent.BedrockagentAgent",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetAgentCollaboration(val *string) {
	if err := j.validateSetAgentCollaborationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentCollaboration",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetAgentName(val *string) {
	if err := j.validateSetAgentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentName",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetAgentResourceRoleArn(val *string) {
	if err := j.validateSetAgentResourceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentResourceRoleArn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetCustomerEncryptionKeyArn(val *string) {
	if err := j.validateSetCustomerEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerEncryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetFoundationModel(val *string) {
	if err := j.validateSetFoundationModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"foundationModel",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetIdleSessionTtlInSeconds(val *float64) {
	if err := j.validateSetIdleSessionTtlInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleSessionTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetInstruction(val *string) {
	if err := j.validateSetInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instruction",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetPrepareAgent(val interface{}) {
	if err := j.validateSetPrepareAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prepareAgent",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetSkipResourceInUseCheck(val interface{}) {
	if err := j.validateSetSkipResourceInUseCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipResourceInUseCheck",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgent)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a BedrockagentAgent resource upon running "cdktf plan <stack-name>".
func BedrockagentAgent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBedrockagentAgent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgent.BedrockagentAgent",
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
func BedrockagentAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgent.BedrockagentAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentAgent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgent.BedrockagentAgent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentAgent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentAgent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgent.BedrockagentAgent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BedrockagentAgent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.bedrockagentAgent.BedrockagentAgent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockagentAgent) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BedrockagentAgent) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BedrockagentAgent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentAgent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockagentAgent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentAgent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentAgent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentAgent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentAgent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentAgent) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentAgent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentAgent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgent) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BedrockagentAgent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockagentAgent) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentAgent) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BedrockagentAgent) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentAgent) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BedrockagentAgent) PutGuardrailConfiguration(value interface{}) {
	if err := b.validatePutGuardrailConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putGuardrailConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentAgent) PutPromptOverrideConfiguration(value interface{}) {
	if err := b.validatePutPromptOverrideConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putPromptOverrideConfiguration",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentAgent) PutTimeouts(value *BedrockagentAgentTimeouts) {
	if err := b.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetAgentCollaboration() {
	_jsii_.InvokeVoid(
		b,
		"resetAgentCollaboration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetCustomerEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		b,
		"resetCustomerEncryptionKeyArn",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetGuardrailConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetGuardrailConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetIdleSessionTtlInSeconds() {
	_jsii_.InvokeVoid(
		b,
		"resetIdleSessionTtlInSeconds",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetInstruction() {
	_jsii_.InvokeVoid(
		b,
		"resetInstruction",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetPrepareAgent() {
	_jsii_.InvokeVoid(
		b,
		"resetPrepareAgent",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetPromptOverrideConfiguration() {
	_jsii_.InvokeVoid(
		b,
		"resetPromptOverrideConfiguration",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetSkipResourceInUseCheck() {
	_jsii_.InvokeVoid(
		b,
		"resetSkipResourceInUseCheck",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetTags() {
	_jsii_.InvokeVoid(
		b,
		"resetTags",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		b,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


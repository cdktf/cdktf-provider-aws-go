// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentagentactiongroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/bedrockagentagentactiongroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/bedrockagent_agent_action_group aws_bedrockagent_agent_action_group}.
type BedrockagentAgentActionGroup interface {
	cdktf.TerraformResource
	ActionGroupExecutor() BedrockagentAgentActionGroupActionGroupExecutorList
	ActionGroupExecutorInput() interface{}
	ActionGroupId() *string
	ActionGroupName() *string
	SetActionGroupName(val *string)
	ActionGroupNameInput() *string
	ActionGroupState() *string
	SetActionGroupState(val *string)
	ActionGroupStateInput() *string
	AgentId() *string
	SetAgentId(val *string)
	AgentIdInput() *string
	AgentVersion() *string
	SetAgentVersion(val *string)
	AgentVersionInput() *string
	ApiSchema() BedrockagentAgentActionGroupApiSchemaList
	ApiSchemaInput() interface{}
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	ParentActionGroupSignature() *string
	SetParentActionGroupSignature(val *string)
	ParentActionGroupSignatureInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutActionGroupExecutor(value interface{})
	PutApiSchema(value interface{})
	ResetActionGroupExecutor()
	ResetActionGroupState()
	ResetApiSchema()
	ResetDescription()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentActionGroupSignature()
	ResetSkipResourceInUseCheck()
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

// The jsii proxy struct for BedrockagentAgentActionGroup
type jsiiProxy_BedrockagentAgentActionGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ActionGroupExecutor() BedrockagentAgentActionGroupActionGroupExecutorList {
	var returns BedrockagentAgentActionGroupActionGroupExecutorList
	_jsii_.Get(
		j,
		"actionGroupExecutor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ActionGroupExecutorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionGroupExecutorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ActionGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ActionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ActionGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ActionGroupState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ActionGroupStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) AgentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) AgentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ApiSchema() BedrockagentAgentActionGroupApiSchemaList {
	var returns BedrockagentAgentActionGroupApiSchemaList
	_jsii_.Get(
		j,
		"apiSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ApiSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ParentActionGroupSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentActionGroupSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) ParentActionGroupSignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentActionGroupSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) SkipResourceInUseCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) SkipResourceInUseCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BedrockagentAgentActionGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/bedrockagent_agent_action_group aws_bedrockagent_agent_action_group} Resource.
func NewBedrockagentAgentActionGroup(scope constructs.Construct, id *string, config *BedrockagentAgentActionGroupConfig) BedrockagentAgentActionGroup {
	_init_.Initialize()

	if err := validateNewBedrockagentAgentActionGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BedrockagentAgentActionGroup{}

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentAgentActionGroup.BedrockagentAgentActionGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/bedrockagent_agent_action_group aws_bedrockagent_agent_action_group} Resource.
func NewBedrockagentAgentActionGroup_Override(b BedrockagentAgentActionGroup, scope constructs.Construct, id *string, config *BedrockagentAgentActionGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.bedrockagentAgentActionGroup.BedrockagentAgentActionGroup",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetActionGroupName(val *string) {
	if err := j.validateSetActionGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionGroupName",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetActionGroupState(val *string) {
	if err := j.validateSetActionGroupStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionGroupState",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetAgentId(val *string) {
	if err := j.validateSetAgentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentId",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetAgentVersion(val *string) {
	if err := j.validateSetAgentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetParentActionGroupSignature(val *string) {
	if err := j.validateSetParentActionGroupSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentActionGroupSignature",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BedrockagentAgentActionGroup)SetSkipResourceInUseCheck(val interface{}) {
	if err := j.validateSetSkipResourceInUseCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipResourceInUseCheck",
		val,
	)
}

// Generates CDKTF code for importing a BedrockagentAgentActionGroup resource upon running "cdktf plan <stack-name>".
func BedrockagentAgentActionGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBedrockagentAgentActionGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgentActionGroup.BedrockagentAgentActionGroup",
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
func BedrockagentAgentActionGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentAgentActionGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgentActionGroup.BedrockagentAgentActionGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentAgentActionGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentAgentActionGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgentActionGroup.BedrockagentAgentActionGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BedrockagentAgentActionGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBedrockagentAgentActionGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.bedrockagentAgentActionGroup.BedrockagentAgentActionGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BedrockagentAgentActionGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.bedrockagentAgentActionGroup.BedrockagentAgentActionGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BedrockagentAgentActionGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockagentAgentActionGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BedrockagentAgentActionGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BedrockagentAgentActionGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BedrockagentAgentActionGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BedrockagentAgentActionGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BedrockagentAgentActionGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BedrockagentAgentActionGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BedrockagentAgentActionGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BedrockagentAgentActionGroup) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) PutActionGroupExecutor(value interface{}) {
	if err := b.validatePutActionGroupExecutorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putActionGroupExecutor",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) PutApiSchema(value interface{}) {
	if err := b.validatePutApiSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putApiSchema",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ResetActionGroupExecutor() {
	_jsii_.InvokeVoid(
		b,
		"resetActionGroupExecutor",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ResetActionGroupState() {
	_jsii_.InvokeVoid(
		b,
		"resetActionGroupState",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ResetApiSchema() {
	_jsii_.InvokeVoid(
		b,
		"resetApiSchema",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		b,
		"resetDescription",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ResetParentActionGroupSignature() {
	_jsii_.InvokeVoid(
		b,
		"resetParentActionGroupSignature",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ResetSkipResourceInUseCheck() {
	_jsii_.InvokeVoid(
		b,
		"resetSkipResourceInUseCheck",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BedrockagentAgentActionGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


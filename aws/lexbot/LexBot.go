// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexbot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lexbot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/lex_bot aws_lex_bot}.
type LexBot interface {
	cdktf.TerraformResource
	AbortStatement() LexBotAbortStatementOutputReference
	AbortStatementInput() *LexBotAbortStatement
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Checksum() *string
	ChildDirected() interface{}
	SetChildDirected(val interface{})
	ChildDirectedInput() interface{}
	ClarificationPrompt() LexBotClarificationPromptOutputReference
	ClarificationPromptInput() *LexBotClarificationPrompt
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
	CreatedDate() *string
	CreateVersion() interface{}
	SetCreateVersion(val interface{})
	CreateVersionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DetectSentiment() interface{}
	SetDetectSentiment(val interface{})
	DetectSentimentInput() interface{}
	EnableModelImprovements() interface{}
	SetEnableModelImprovements(val interface{})
	EnableModelImprovementsInput() interface{}
	FailureReason() *string
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
	IdleSessionTtlInSeconds() *float64
	SetIdleSessionTtlInSeconds(val *float64)
	IdleSessionTtlInSecondsInput() *float64
	Intent() LexBotIntentList
	IntentInput() interface{}
	LastUpdatedDate() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locale() *string
	SetLocale(val *string)
	LocaleInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NluIntentConfidenceThreshold() *float64
	SetNluIntentConfidenceThreshold(val *float64)
	NluIntentConfidenceThresholdInput() *float64
	// The tree node.
	Node() constructs.Node
	ProcessBehavior() *string
	SetProcessBehavior(val *string)
	ProcessBehaviorInput() *string
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
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LexBotTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	VoiceId() *string
	SetVoiceId(val *string)
	VoiceIdInput() *string
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
	PutAbortStatement(value *LexBotAbortStatement)
	PutClarificationPrompt(value *LexBotClarificationPrompt)
	PutIntent(value interface{})
	PutTimeouts(value *LexBotTimeouts)
	ResetClarificationPrompt()
	ResetCreateVersion()
	ResetDescription()
	ResetDetectSentiment()
	ResetEnableModelImprovements()
	ResetId()
	ResetIdleSessionTtlInSeconds()
	ResetLocale()
	ResetNluIntentConfidenceThreshold()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProcessBehavior()
	ResetTimeouts()
	ResetVoiceId()
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

// The jsii proxy struct for LexBot
type jsiiProxy_LexBot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LexBot) AbortStatement() LexBotAbortStatementOutputReference {
	var returns LexBotAbortStatementOutputReference
	_jsii_.Get(
		j,
		"abortStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) AbortStatementInput() *LexBotAbortStatement {
	var returns *LexBotAbortStatement
	_jsii_.Get(
		j,
		"abortStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ChildDirected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"childDirected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ChildDirectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"childDirectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ClarificationPrompt() LexBotClarificationPromptOutputReference {
	var returns LexBotClarificationPromptOutputReference
	_jsii_.Get(
		j,
		"clarificationPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ClarificationPromptInput() *LexBotClarificationPrompt {
	var returns *LexBotClarificationPrompt
	_jsii_.Get(
		j,
		"clarificationPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) CreateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) CreateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) DetectSentiment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectSentiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) DetectSentimentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectSentimentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) EnableModelImprovements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableModelImprovements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) EnableModelImprovementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableModelImprovementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) IdleSessionTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) IdleSessionTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Intent() LexBotIntentList {
	var returns LexBotIntentList
	_jsii_.Get(
		j,
		"intent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) IntentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) NluIntentConfidenceThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nluIntentConfidenceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) NluIntentConfidenceThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nluIntentConfidenceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ProcessBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) ProcessBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Timeouts() LexBotTimeoutsOutputReference {
	var returns LexBotTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) VoiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LexBot) VoiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/lex_bot aws_lex_bot} Resource.
func NewLexBot(scope constructs.Construct, id *string, config *LexBotConfig) LexBot {
	_init_.Initialize()

	if err := validateNewLexBotParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LexBot{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexBot.LexBot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.62.0/docs/resources/lex_bot aws_lex_bot} Resource.
func NewLexBot_Override(l LexBot, scope constructs.Construct, id *string, config *LexBotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexBot.LexBot",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LexBot)SetChildDirected(val interface{}) {
	if err := j.validateSetChildDirectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childDirected",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetCreateVersion(val interface{}) {
	if err := j.validateSetCreateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createVersion",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetDetectSentiment(val interface{}) {
	if err := j.validateSetDetectSentimentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectSentiment",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetEnableModelImprovements(val interface{}) {
	if err := j.validateSetEnableModelImprovementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableModelImprovements",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetIdleSessionTtlInSeconds(val *float64) {
	if err := j.validateSetIdleSessionTtlInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleSessionTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetNluIntentConfidenceThreshold(val *float64) {
	if err := j.validateSetNluIntentConfidenceThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nluIntentConfidenceThreshold",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetProcessBehavior(val *string) {
	if err := j.validateSetProcessBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processBehavior",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LexBot)SetVoiceId(val *string) {
	if err := j.validateSetVoiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"voiceId",
		val,
	)
}

// Generates CDKTF code for importing a LexBot resource upon running "cdktf plan <stack-name>".
func LexBot_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLexBot_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexBot.LexBot",
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
func LexBot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexBot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexBot.LexBot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LexBot_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexBot_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexBot.LexBot",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LexBot_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexBot_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexBot.LexBot",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LexBot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lexBot.LexBot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LexBot) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LexBot) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LexBot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LexBot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LexBot) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LexBot) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LexBot) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LexBot) PutAbortStatement(value *LexBotAbortStatement) {
	if err := l.validatePutAbortStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAbortStatement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBot) PutClarificationPrompt(value *LexBotClarificationPrompt) {
	if err := l.validatePutClarificationPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putClarificationPrompt",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBot) PutIntent(value interface{}) {
	if err := l.validatePutIntentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIntent",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBot) PutTimeouts(value *LexBotTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LexBot) ResetClarificationPrompt() {
	_jsii_.InvokeVoid(
		l,
		"resetClarificationPrompt",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetCreateVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetCreateVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetDetectSentiment() {
	_jsii_.InvokeVoid(
		l,
		"resetDetectSentiment",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetEnableModelImprovements() {
	_jsii_.InvokeVoid(
		l,
		"resetEnableModelImprovements",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetIdleSessionTtlInSeconds() {
	_jsii_.InvokeVoid(
		l,
		"resetIdleSessionTtlInSeconds",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetLocale() {
	_jsii_.InvokeVoid(
		l,
		"resetLocale",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetNluIntentConfidenceThreshold() {
	_jsii_.InvokeVoid(
		l,
		"resetNluIntentConfidenceThreshold",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetProcessBehavior() {
	_jsii_.InvokeVoid(
		l,
		"resetProcessBehavior",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) ResetVoiceId() {
	_jsii_.InvokeVoid(
		l,
		"resetVoiceId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LexBot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LexBot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


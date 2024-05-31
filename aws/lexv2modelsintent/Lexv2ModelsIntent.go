// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsintent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lexv2modelsintent/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lexv2models_intent aws_lexv2models_intent}.
type Lexv2ModelsIntent interface {
	cdktf.TerraformResource
	BotId() *string
	SetBotId(val *string)
	BotIdInput() *string
	BotVersion() *string
	SetBotVersion(val *string)
	BotVersionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClosingSetting() Lexv2ModelsIntentClosingSettingList
	ClosingSettingInput() interface{}
	ConfirmationSetting() Lexv2ModelsIntentConfirmationSettingList
	ConfirmationSettingInput() interface{}
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
	CreationDateTime() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DialogCodeHook() Lexv2ModelsIntentDialogCodeHookList
	DialogCodeHookInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FulfillmentCodeHook() Lexv2ModelsIntentFulfillmentCodeHookList
	FulfillmentCodeHookInput() interface{}
	Id() *string
	InitialResponseSetting() Lexv2ModelsIntentInitialResponseSettingList
	InitialResponseSettingInput() interface{}
	InputContext() Lexv2ModelsIntentInputContextList
	InputContextInput() interface{}
	IntentId() *string
	KendraConfiguration() Lexv2ModelsIntentKendraConfigurationList
	KendraConfigurationInput() interface{}
	LastUpdatedDateTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocaleId() *string
	SetLocaleId(val *string)
	LocaleIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutputContext() Lexv2ModelsIntentOutputContextList
	OutputContextInput() interface{}
	ParentIntentSignature() *string
	SetParentIntentSignature(val *string)
	ParentIntentSignatureInput() *string
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
	SampleUtterance() Lexv2ModelsIntentSampleUtteranceList
	SampleUtteranceInput() interface{}
	SlotPriority() Lexv2ModelsIntentSlotPriorityList
	SlotPriorityInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Lexv2ModelsIntentTimeoutsOutputReference
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
	PutClosingSetting(value interface{})
	PutConfirmationSetting(value interface{})
	PutDialogCodeHook(value interface{})
	PutFulfillmentCodeHook(value interface{})
	PutInitialResponseSetting(value interface{})
	PutInputContext(value interface{})
	PutKendraConfiguration(value interface{})
	PutOutputContext(value interface{})
	PutSampleUtterance(value interface{})
	PutSlotPriority(value interface{})
	PutTimeouts(value *Lexv2ModelsIntentTimeouts)
	ResetClosingSetting()
	ResetConfirmationSetting()
	ResetDescription()
	ResetDialogCodeHook()
	ResetFulfillmentCodeHook()
	ResetInitialResponseSetting()
	ResetInputContext()
	ResetKendraConfiguration()
	ResetOutputContext()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentIntentSignature()
	ResetSampleUtterance()
	ResetSlotPriority()
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

// The jsii proxy struct for Lexv2ModelsIntent
type jsiiProxy_Lexv2ModelsIntent struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Lexv2ModelsIntent) BotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) BotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) BotVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) BotVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) ClosingSetting() Lexv2ModelsIntentClosingSettingList {
	var returns Lexv2ModelsIntentClosingSettingList
	_jsii_.Get(
		j,
		"closingSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) ClosingSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closingSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) ConfirmationSetting() Lexv2ModelsIntentConfirmationSettingList {
	var returns Lexv2ModelsIntentConfirmationSettingList
	_jsii_.Get(
		j,
		"confirmationSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) ConfirmationSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) CreationDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) DialogCodeHook() Lexv2ModelsIntentDialogCodeHookList {
	var returns Lexv2ModelsIntentDialogCodeHookList
	_jsii_.Get(
		j,
		"dialogCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) DialogCodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dialogCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) FulfillmentCodeHook() Lexv2ModelsIntentFulfillmentCodeHookList {
	var returns Lexv2ModelsIntentFulfillmentCodeHookList
	_jsii_.Get(
		j,
		"fulfillmentCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) FulfillmentCodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fulfillmentCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) InitialResponseSetting() Lexv2ModelsIntentInitialResponseSettingList {
	var returns Lexv2ModelsIntentInitialResponseSettingList
	_jsii_.Get(
		j,
		"initialResponseSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) InitialResponseSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialResponseSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) InputContext() Lexv2ModelsIntentInputContextList {
	var returns Lexv2ModelsIntentInputContextList
	_jsii_.Get(
		j,
		"inputContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) InputContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) IntentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) KendraConfiguration() Lexv2ModelsIntentKendraConfigurationList {
	var returns Lexv2ModelsIntentKendraConfigurationList
	_jsii_.Get(
		j,
		"kendraConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) KendraConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kendraConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) LastUpdatedDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) LocaleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) LocaleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) OutputContext() Lexv2ModelsIntentOutputContextList {
	var returns Lexv2ModelsIntentOutputContextList
	_jsii_.Get(
		j,
		"outputContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) OutputContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) ParentIntentSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIntentSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) ParentIntentSignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIntentSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) SampleUtterance() Lexv2ModelsIntentSampleUtteranceList {
	var returns Lexv2ModelsIntentSampleUtteranceList
	_jsii_.Get(
		j,
		"sampleUtterance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) SampleUtteranceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sampleUtteranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) SlotPriority() Lexv2ModelsIntentSlotPriorityList {
	var returns Lexv2ModelsIntentSlotPriorityList
	_jsii_.Get(
		j,
		"slotPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) SlotPriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slotPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) Timeouts() Lexv2ModelsIntentTimeoutsOutputReference {
	var returns Lexv2ModelsIntentTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsIntent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lexv2models_intent aws_lexv2models_intent} Resource.
func NewLexv2ModelsIntent(scope constructs.Construct, id *string, config *Lexv2ModelsIntentConfig) Lexv2ModelsIntent {
	_init_.Initialize()

	if err := validateNewLexv2ModelsIntentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsIntent{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/lexv2models_intent aws_lexv2models_intent} Resource.
func NewLexv2ModelsIntent_Override(l Lexv2ModelsIntent, scope constructs.Construct, id *string, config *Lexv2ModelsIntentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntent",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetBotId(val *string) {
	if err := j.validateSetBotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botId",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetBotVersion(val *string) {
	if err := j.validateSetBotVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botVersion",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetLocaleId(val *string) {
	if err := j.validateSetLocaleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localeId",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetParentIntentSignature(val *string) {
	if err := j.validateSetParentIntentSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentIntentSignature",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsIntent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a Lexv2ModelsIntent resource upon running "cdktf plan <stack-name>".
func Lexv2ModelsIntent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLexv2ModelsIntent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntent",
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
func Lexv2ModelsIntent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexv2ModelsIntent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Lexv2ModelsIntent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexv2ModelsIntent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Lexv2ModelsIntent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexv2ModelsIntent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Lexv2ModelsIntent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lexv2ModelsIntent.Lexv2ModelsIntent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntent) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsIntent) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsIntent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsIntent) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsIntent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsIntent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntent) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsIntent) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) PutClosingSetting(value interface{}) {
	if err := l.validatePutClosingSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putClosingSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) PutConfirmationSetting(value interface{}) {
	if err := l.validatePutConfirmationSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putConfirmationSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) PutDialogCodeHook(value interface{}) {
	if err := l.validatePutDialogCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDialogCodeHook",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) PutFulfillmentCodeHook(value interface{}) {
	if err := l.validatePutFulfillmentCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putFulfillmentCodeHook",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) PutInitialResponseSetting(value interface{}) {
	if err := l.validatePutInitialResponseSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putInitialResponseSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) PutInputContext(value interface{}) {
	if err := l.validatePutInputContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putInputContext",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) PutKendraConfiguration(value interface{}) {
	if err := l.validatePutKendraConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putKendraConfiguration",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) PutOutputContext(value interface{}) {
	if err := l.validatePutOutputContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOutputContext",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) PutSampleUtterance(value interface{}) {
	if err := l.validatePutSampleUtteranceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSampleUtterance",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) PutSlotPriority(value interface{}) {
	if err := l.validatePutSlotPriorityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSlotPriority",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) PutTimeouts(value *Lexv2ModelsIntentTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetClosingSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetClosingSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetConfirmationSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetConfirmationSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetDialogCodeHook() {
	_jsii_.InvokeVoid(
		l,
		"resetDialogCodeHook",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetFulfillmentCodeHook() {
	_jsii_.InvokeVoid(
		l,
		"resetFulfillmentCodeHook",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetInitialResponseSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetInitialResponseSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetInputContext() {
	_jsii_.InvokeVoid(
		l,
		"resetInputContext",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetKendraConfiguration() {
	_jsii_.InvokeVoid(
		l,
		"resetKendraConfiguration",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetOutputContext() {
	_jsii_.InvokeVoid(
		l,
		"resetOutputContext",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetParentIntentSignature() {
	_jsii_.InvokeVoid(
		l,
		"resetParentIntentSignature",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetSampleUtterance() {
	_jsii_.InvokeVoid(
		l,
		"resetSampleUtterance",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetSlotPriority() {
	_jsii_.InvokeVoid(
		l,
		"resetSlotPriority",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsIntent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsIntent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


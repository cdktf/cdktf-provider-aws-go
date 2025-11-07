// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lexv2modelsslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/lexv2models_slot aws_lexv2models_slot}.
type Lexv2ModelsSlot interface {
	cdktf.TerraformResource
	BotId() *string
	SetBotId(val *string)
	BotIdInput() *string
	BotVersion() *string
	SetBotVersion(val *string)
	BotVersionInput() *string
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
	IntentId() *string
	SetIntentId(val *string)
	IntentIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocaleId() *string
	SetLocaleId(val *string)
	LocaleIdInput() *string
	MultipleValuesSetting() Lexv2ModelsSlotMultipleValuesSettingList
	MultipleValuesSettingInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ObfuscationSetting() Lexv2ModelsSlotObfuscationSettingList
	ObfuscationSettingInput() interface{}
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
	SlotId() *string
	SlotTypeId() *string
	SetSlotTypeId(val *string)
	SlotTypeIdInput() *string
	SubSlotSetting() Lexv2ModelsSlotSubSlotSettingList
	SubSlotSettingInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Lexv2ModelsSlotTimeoutsOutputReference
	TimeoutsInput() interface{}
	ValueElicitationSetting() Lexv2ModelsSlotValueElicitationSettingList
	ValueElicitationSettingInput() interface{}
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
	PutMultipleValuesSetting(value interface{})
	PutObfuscationSetting(value interface{})
	PutSubSlotSetting(value interface{})
	PutTimeouts(value *Lexv2ModelsSlotTimeouts)
	PutValueElicitationSetting(value interface{})
	ResetDescription()
	ResetMultipleValuesSetting()
	ResetObfuscationSetting()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetSlotTypeId()
	ResetSubSlotSetting()
	ResetTimeouts()
	ResetValueElicitationSetting()
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

// The jsii proxy struct for Lexv2ModelsSlot
type jsiiProxy_Lexv2ModelsSlot struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Lexv2ModelsSlot) BotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) BotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) BotVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) BotVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) IntentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) IntentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) LocaleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) LocaleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) MultipleValuesSetting() Lexv2ModelsSlotMultipleValuesSettingList {
	var returns Lexv2ModelsSlotMultipleValuesSettingList
	_jsii_.Get(
		j,
		"multipleValuesSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) MultipleValuesSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multipleValuesSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) ObfuscationSetting() Lexv2ModelsSlotObfuscationSettingList {
	var returns Lexv2ModelsSlotObfuscationSettingList
	_jsii_.Get(
		j,
		"obfuscationSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) ObfuscationSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"obfuscationSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) SlotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) SlotTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) SlotTypeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotTypeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) SubSlotSetting() Lexv2ModelsSlotSubSlotSettingList {
	var returns Lexv2ModelsSlotSubSlotSettingList
	_jsii_.Get(
		j,
		"subSlotSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) SubSlotSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"subSlotSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) Timeouts() Lexv2ModelsSlotTimeoutsOutputReference {
	var returns Lexv2ModelsSlotTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) ValueElicitationSetting() Lexv2ModelsSlotValueElicitationSettingList {
	var returns Lexv2ModelsSlotValueElicitationSettingList
	_jsii_.Get(
		j,
		"valueElicitationSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlot) ValueElicitationSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueElicitationSettingInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/lexv2models_slot aws_lexv2models_slot} Resource.
func NewLexv2ModelsSlot(scope constructs.Construct, id *string, config *Lexv2ModelsSlotConfig) Lexv2ModelsSlot {
	_init_.Initialize()

	if err := validateNewLexv2ModelsSlotParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsSlot{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsSlot.Lexv2ModelsSlot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/lexv2models_slot aws_lexv2models_slot} Resource.
func NewLexv2ModelsSlot_Override(l Lexv2ModelsSlot, scope constructs.Construct, id *string, config *Lexv2ModelsSlotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsSlot.Lexv2ModelsSlot",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetBotId(val *string) {
	if err := j.validateSetBotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botId",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetBotVersion(val *string) {
	if err := j.validateSetBotVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botVersion",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetIntentId(val *string) {
	if err := j.validateSetIntentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intentId",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetLocaleId(val *string) {
	if err := j.validateSetLocaleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localeId",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlot)SetSlotTypeId(val *string) {
	if err := j.validateSetSlotTypeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"slotTypeId",
		val,
	)
}

// Generates CDKTF code for importing a Lexv2ModelsSlot resource upon running "cdktf plan <stack-name>".
func Lexv2ModelsSlot_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLexv2ModelsSlot_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsSlot.Lexv2ModelsSlot",
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
func Lexv2ModelsSlot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexv2ModelsSlot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsSlot.Lexv2ModelsSlot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Lexv2ModelsSlot_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexv2ModelsSlot_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsSlot.Lexv2ModelsSlot",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Lexv2ModelsSlot_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexv2ModelsSlot_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsSlot.Lexv2ModelsSlot",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Lexv2ModelsSlot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lexv2ModelsSlot.Lexv2ModelsSlot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlot) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsSlot) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsSlot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsSlot) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsSlot) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlot) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsSlot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsSlot) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlot) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsSlot) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) PutMultipleValuesSetting(value interface{}) {
	if err := l.validatePutMultipleValuesSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putMultipleValuesSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) PutObfuscationSetting(value interface{}) {
	if err := l.validatePutObfuscationSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putObfuscationSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) PutSubSlotSetting(value interface{}) {
	if err := l.validatePutSubSlotSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSubSlotSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) PutTimeouts(value *Lexv2ModelsSlotTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) PutValueElicitationSetting(value interface{}) {
	if err := l.validatePutValueElicitationSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putValueElicitationSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) ResetMultipleValuesSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetMultipleValuesSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) ResetObfuscationSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetObfuscationSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) ResetRegion() {
	_jsii_.InvokeVoid(
		l,
		"resetRegion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) ResetSlotTypeId() {
	_jsii_.InvokeVoid(
		l,
		"resetSlotTypeId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) ResetSubSlotSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetSubSlotSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) ResetValueElicitationSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetValueElicitationSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlot) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlot) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsslottype

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/lexv2modelsslottype/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/lexv2models_slot_type aws_lexv2models_slot_type}.
type Lexv2ModelsSlotType interface {
	cdktf.TerraformResource
	BotId() *string
	SetBotId(val *string)
	BotIdInput() *string
	BotVersion() *string
	SetBotVersion(val *string)
	BotVersionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CompositeSlotTypeSetting() Lexv2ModelsSlotTypeCompositeSlotTypeSettingList
	CompositeSlotTypeSettingInput() interface{}
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
	ExternalSourceSetting() Lexv2ModelsSlotTypeExternalSourceSettingList
	ExternalSourceSettingInput() interface{}
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
	LocaleId() *string
	SetLocaleId(val *string)
	LocaleIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ParentSlotTypeSignature() *string
	SetParentSlotTypeSignature(val *string)
	ParentSlotTypeSignatureInput() *string
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
	SlotTypeId() *string
	SlotTypeValues() Lexv2ModelsSlotTypeSlotTypeValuesList
	SlotTypeValuesInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() Lexv2ModelsSlotTypeTimeoutsOutputReference
	TimeoutsInput() interface{}
	ValueSelectionSetting() Lexv2ModelsSlotTypeValueSelectionSettingList
	ValueSelectionSettingInput() interface{}
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
	PutCompositeSlotTypeSetting(value interface{})
	PutExternalSourceSetting(value interface{})
	PutSlotTypeValues(value interface{})
	PutTimeouts(value *Lexv2ModelsSlotTypeTimeouts)
	PutValueSelectionSetting(value interface{})
	ResetCompositeSlotTypeSetting()
	ResetDescription()
	ResetExternalSourceSetting()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParentSlotTypeSignature()
	ResetSlotTypeValues()
	ResetTimeouts()
	ResetValueSelectionSetting()
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

// The jsii proxy struct for Lexv2ModelsSlotType
type jsiiProxy_Lexv2ModelsSlotType struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Lexv2ModelsSlotType) BotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) BotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) BotVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) BotVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) CompositeSlotTypeSetting() Lexv2ModelsSlotTypeCompositeSlotTypeSettingList {
	var returns Lexv2ModelsSlotTypeCompositeSlotTypeSettingList
	_jsii_.Get(
		j,
		"compositeSlotTypeSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) CompositeSlotTypeSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compositeSlotTypeSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) ExternalSourceSetting() Lexv2ModelsSlotTypeExternalSourceSettingList {
	var returns Lexv2ModelsSlotTypeExternalSourceSettingList
	_jsii_.Get(
		j,
		"externalSourceSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) ExternalSourceSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalSourceSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) LocaleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) LocaleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) ParentSlotTypeSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentSlotTypeSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) ParentSlotTypeSignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentSlotTypeSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) SlotTypeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"slotTypeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) SlotTypeValues() Lexv2ModelsSlotTypeSlotTypeValuesList {
	var returns Lexv2ModelsSlotTypeSlotTypeValuesList
	_jsii_.Get(
		j,
		"slotTypeValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) SlotTypeValuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slotTypeValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) Timeouts() Lexv2ModelsSlotTypeTimeoutsOutputReference {
	var returns Lexv2ModelsSlotTypeTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) ValueSelectionSetting() Lexv2ModelsSlotTypeValueSelectionSettingList {
	var returns Lexv2ModelsSlotTypeValueSelectionSettingList
	_jsii_.Get(
		j,
		"valueSelectionSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Lexv2ModelsSlotType) ValueSelectionSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueSelectionSettingInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/lexv2models_slot_type aws_lexv2models_slot_type} Resource.
func NewLexv2ModelsSlotType(scope constructs.Construct, id *string, config *Lexv2ModelsSlotTypeConfig) Lexv2ModelsSlotType {
	_init_.Initialize()

	if err := validateNewLexv2ModelsSlotTypeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Lexv2ModelsSlotType{}

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsSlotType.Lexv2ModelsSlotType",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/lexv2models_slot_type aws_lexv2models_slot_type} Resource.
func NewLexv2ModelsSlotType_Override(l Lexv2ModelsSlotType, scope constructs.Construct, id *string, config *Lexv2ModelsSlotTypeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lexv2ModelsSlotType.Lexv2ModelsSlotType",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetBotId(val *string) {
	if err := j.validateSetBotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botId",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetBotVersion(val *string) {
	if err := j.validateSetBotVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botVersion",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetLocaleId(val *string) {
	if err := j.validateSetLocaleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localeId",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetParentSlotTypeSignature(val *string) {
	if err := j.validateSetParentSlotTypeSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentSlotTypeSignature",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Lexv2ModelsSlotType)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a Lexv2ModelsSlotType resource upon running "cdktf plan <stack-name>".
func Lexv2ModelsSlotType_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLexv2ModelsSlotType_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsSlotType.Lexv2ModelsSlotType",
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
func Lexv2ModelsSlotType_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexv2ModelsSlotType_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsSlotType.Lexv2ModelsSlotType",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Lexv2ModelsSlotType_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexv2ModelsSlotType_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsSlotType.Lexv2ModelsSlotType",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Lexv2ModelsSlotType_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLexv2ModelsSlotType_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lexv2ModelsSlotType.Lexv2ModelsSlotType",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Lexv2ModelsSlotType_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lexv2ModelsSlotType.Lexv2ModelsSlotType",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotType) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_Lexv2ModelsSlotType) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsSlotType) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_Lexv2ModelsSlotType) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_Lexv2ModelsSlotType) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlotType) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlotType) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_Lexv2ModelsSlotType) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_Lexv2ModelsSlotType) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_Lexv2ModelsSlotType) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_Lexv2ModelsSlotType) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) PutCompositeSlotTypeSetting(value interface{}) {
	if err := l.validatePutCompositeSlotTypeSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCompositeSlotTypeSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) PutExternalSourceSetting(value interface{}) {
	if err := l.validatePutExternalSourceSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putExternalSourceSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) PutSlotTypeValues(value interface{}) {
	if err := l.validatePutSlotTypeValuesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSlotTypeValues",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) PutTimeouts(value *Lexv2ModelsSlotTypeTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) PutValueSelectionSetting(value interface{}) {
	if err := l.validatePutValueSelectionSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putValueSelectionSetting",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ResetCompositeSlotTypeSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetCompositeSlotTypeSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ResetDescription() {
	_jsii_.InvokeVoid(
		l,
		"resetDescription",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ResetExternalSourceSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetExternalSourceSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ResetParentSlotTypeSignature() {
	_jsii_.InvokeVoid(
		l,
		"resetParentSlotTypeSignature",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ResetSlotTypeValues() {
	_jsii_.InvokeVoid(
		l,
		"resetSlotTypeValues",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ResetValueSelectionSetting() {
	_jsii_.InvokeVoid(
		l,
		"resetValueSelectionSetting",
		nil, // no parameters
	)
}

func (l *jsiiProxy_Lexv2ModelsSlotType) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotType) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_Lexv2ModelsSlotType) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


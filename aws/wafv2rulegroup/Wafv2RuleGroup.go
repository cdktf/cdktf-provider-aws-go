// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2rulegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/wafv2rulegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/wafv2_rule_group aws_wafv2_rule_group}.
type Wafv2RuleGroup interface {
	cdktf.TerraformResource
	Arn() *string
	Capacity() *float64
	SetCapacity(val *float64)
	CapacityInput() *float64
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
	CustomResponseBody() Wafv2RuleGroupCustomResponseBodyList
	CustomResponseBodyInput() interface{}
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
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LockToken() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	NamePrefix() *string
	SetNamePrefix(val *string)
	NamePrefixInput() *string
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
	Rule() Wafv2RuleGroupRuleList
	RuleInput() interface{}
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsAll() *map[string]*string
	SetTagsAll(val *map[string]*string)
	TagsAllInput() *map[string]*string
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VisibilityConfig() Wafv2RuleGroupVisibilityConfigOutputReference
	VisibilityConfigInput() *Wafv2RuleGroupVisibilityConfig
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
	PutCustomResponseBody(value interface{})
	PutRule(value interface{})
	PutVisibilityConfig(value *Wafv2RuleGroupVisibilityConfig)
	ResetCustomResponseBody()
	ResetDescription()
	ResetId()
	ResetName()
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRule()
	ResetTags()
	ResetTagsAll()
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

// The jsii proxy struct for Wafv2RuleGroup
type jsiiProxy_Wafv2RuleGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Wafv2RuleGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) CapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) CustomResponseBody() Wafv2RuleGroupCustomResponseBodyList {
	var returns Wafv2RuleGroupCustomResponseBodyList
	_jsii_.Get(
		j,
		"customResponseBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) CustomResponseBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customResponseBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) LockToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Rule() Wafv2RuleGroupRuleList {
	var returns Wafv2RuleGroupRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) VisibilityConfig() Wafv2RuleGroupVisibilityConfigOutputReference {
	var returns Wafv2RuleGroupVisibilityConfigOutputReference
	_jsii_.Get(
		j,
		"visibilityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2RuleGroup) VisibilityConfigInput() *Wafv2RuleGroupVisibilityConfig {
	var returns *Wafv2RuleGroupVisibilityConfig
	_jsii_.Get(
		j,
		"visibilityConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/wafv2_rule_group aws_wafv2_rule_group} Resource.
func NewWafv2RuleGroup(scope constructs.Construct, id *string, config *Wafv2RuleGroupConfig) Wafv2RuleGroup {
	_init_.Initialize()

	if err := validateNewWafv2RuleGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2RuleGroup{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/wafv2_rule_group aws_wafv2_rule_group} Resource.
func NewWafv2RuleGroup_Override(w Wafv2RuleGroup, scope constructs.Construct, id *string, config *Wafv2RuleGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroup",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetCapacity(val *float64) {
	if err := j.validateSetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacity",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Wafv2RuleGroup)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTF code for importing a Wafv2RuleGroup resource upon running "cdktf plan <stack-name>".
func Wafv2RuleGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWafv2RuleGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroup",
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
func Wafv2RuleGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafv2RuleGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Wafv2RuleGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafv2RuleGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Wafv2RuleGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafv2RuleGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Wafv2RuleGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.wafv2RuleGroup.Wafv2RuleGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) PutCustomResponseBody(value interface{}) {
	if err := w.validatePutCustomResponseBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCustomResponseBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) PutRule(value interface{}) {
	if err := w.validatePutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRule",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) PutVisibilityConfig(value *Wafv2RuleGroupVisibilityConfig) {
	if err := w.validatePutVisibilityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putVisibilityConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) ResetCustomResponseBody() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomResponseBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetDescription",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) ResetName() {
	_jsii_.InvokeVoid(
		w,
		"resetName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		w,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) ResetRule() {
	_jsii_.InvokeVoid(
		w,
		"resetRule",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) ResetTagsAll() {
	_jsii_.InvokeVoid(
		w,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2RuleGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2RuleGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


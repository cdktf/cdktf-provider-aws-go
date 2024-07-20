// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53recoverycontrolconfigsafetyrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/route53recoverycontrolconfigsafetyrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/route53recoverycontrolconfig_safety_rule aws_route53recoverycontrolconfig_safety_rule}.
type Route53RecoverycontrolconfigSafetyRule interface {
	cdktf.TerraformResource
	Arn() *string
	AssertedControls() *[]*string
	SetAssertedControls(val *[]*string)
	AssertedControlsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ControlPanelArn() *string
	SetControlPanelArn(val *string)
	ControlPanelArnInput() *string
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
	GatingControls() *[]*string
	SetGatingControls(val *[]*string)
	GatingControlsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	RuleConfig() Route53RecoverycontrolconfigSafetyRuleRuleConfigOutputReference
	RuleConfigInput() *Route53RecoverycontrolconfigSafetyRuleRuleConfig
	Status() *string
	TargetControls() *[]*string
	SetTargetControls(val *[]*string)
	TargetControlsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WaitPeriodMs() *float64
	SetWaitPeriodMs(val *float64)
	WaitPeriodMsInput() *float64
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
	PutRuleConfig(value *Route53RecoverycontrolconfigSafetyRuleRuleConfig)
	ResetAssertedControls()
	ResetGatingControls()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTargetControls()
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

// The jsii proxy struct for Route53RecoverycontrolconfigSafetyRule
type jsiiProxy_Route53RecoverycontrolconfigSafetyRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) AssertedControls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assertedControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) AssertedControlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"assertedControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ControlPanelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPanelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ControlPanelArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"controlPanelArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) GatingControls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gatingControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) GatingControlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gatingControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) RuleConfig() Route53RecoverycontrolconfigSafetyRuleRuleConfigOutputReference {
	var returns Route53RecoverycontrolconfigSafetyRuleRuleConfigOutputReference
	_jsii_.Get(
		j,
		"ruleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) RuleConfigInput() *Route53RecoverycontrolconfigSafetyRuleRuleConfig {
	var returns *Route53RecoverycontrolconfigSafetyRuleRuleConfig
	_jsii_.Get(
		j,
		"ruleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) TargetControls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetControls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) TargetControlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"targetControlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) WaitPeriodMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitPeriodMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) WaitPeriodMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitPeriodMsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/route53recoverycontrolconfig_safety_rule aws_route53recoverycontrolconfig_safety_rule} Resource.
func NewRoute53RecoverycontrolconfigSafetyRule(scope constructs.Construct, id *string, config *Route53RecoverycontrolconfigSafetyRuleConfig) Route53RecoverycontrolconfigSafetyRule {
	_init_.Initialize()

	if err := validateNewRoute53RecoverycontrolconfigSafetyRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53RecoverycontrolconfigSafetyRule{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecoverycontrolconfigSafetyRule.Route53RecoverycontrolconfigSafetyRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/route53recoverycontrolconfig_safety_rule aws_route53recoverycontrolconfig_safety_rule} Resource.
func NewRoute53RecoverycontrolconfigSafetyRule_Override(r Route53RecoverycontrolconfigSafetyRule, scope constructs.Construct, id *string, config *Route53RecoverycontrolconfigSafetyRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53RecoverycontrolconfigSafetyRule.Route53RecoverycontrolconfigSafetyRule",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetAssertedControls(val *[]*string) {
	if err := j.validateSetAssertedControlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assertedControls",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetControlPanelArn(val *string) {
	if err := j.validateSetControlPanelArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"controlPanelArn",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetGatingControls(val *[]*string) {
	if err := j.validateSetGatingControlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatingControls",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetTargetControls(val *[]*string) {
	if err := j.validateSetTargetControlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetControls",
		val,
	)
}

func (j *jsiiProxy_Route53RecoverycontrolconfigSafetyRule)SetWaitPeriodMs(val *float64) {
	if err := j.validateSetWaitPeriodMsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitPeriodMs",
		val,
	)
}

// Generates CDKTF code for importing a Route53RecoverycontrolconfigSafetyRule resource upon running "cdktf plan <stack-name>".
func Route53RecoverycontrolconfigSafetyRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRoute53RecoverycontrolconfigSafetyRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53RecoverycontrolconfigSafetyRule.Route53RecoverycontrolconfigSafetyRule",
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
func Route53RecoverycontrolconfigSafetyRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53RecoverycontrolconfigSafetyRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53RecoverycontrolconfigSafetyRule.Route53RecoverycontrolconfigSafetyRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53RecoverycontrolconfigSafetyRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53RecoverycontrolconfigSafetyRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53RecoverycontrolconfigSafetyRule.Route53RecoverycontrolconfigSafetyRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Route53RecoverycontrolconfigSafetyRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRoute53RecoverycontrolconfigSafetyRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.route53RecoverycontrolconfigSafetyRule.Route53RecoverycontrolconfigSafetyRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Route53RecoverycontrolconfigSafetyRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.route53RecoverycontrolconfigSafetyRule.Route53RecoverycontrolconfigSafetyRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) PutRuleConfig(value *Route53RecoverycontrolconfigSafetyRuleRuleConfig) {
	if err := r.validatePutRuleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRuleConfig",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ResetAssertedControls() {
	_jsii_.InvokeVoid(
		r,
		"resetAssertedControls",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ResetGatingControls() {
	_jsii_.InvokeVoid(
		r,
		"resetGatingControls",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ResetTargetControls() {
	_jsii_.InvokeVoid(
		r,
		"resetTargetControls",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53RecoverycontrolconfigSafetyRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


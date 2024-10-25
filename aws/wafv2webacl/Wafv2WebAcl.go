// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafv2webacl

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/wafv2webacl/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/wafv2_web_acl aws_wafv2_web_acl}.
type Wafv2WebAcl interface {
	cdktf.TerraformResource
	ApplicationIntegrationUrl() *string
	Arn() *string
	AssociationConfig() Wafv2WebAclAssociationConfigOutputReference
	AssociationConfigInput() *Wafv2WebAclAssociationConfig
	Capacity() *float64
	CaptchaConfig() Wafv2WebAclCaptchaConfigOutputReference
	CaptchaConfigInput() *Wafv2WebAclCaptchaConfig
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChallengeConfig() Wafv2WebAclChallengeConfigOutputReference
	ChallengeConfigInput() *Wafv2WebAclChallengeConfig
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
	CustomResponseBody() Wafv2WebAclCustomResponseBodyList
	CustomResponseBodyInput() interface{}
	DefaultAction() Wafv2WebAclDefaultActionOutputReference
	DefaultActionInput() *Wafv2WebAclDefaultAction
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
	Rule() Wafv2WebAclRuleList
	RuleInput() interface{}
	RuleJson() *string
	SetRuleJson(val *string)
	RuleJsonInput() *string
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
	TokenDomains() *[]*string
	SetTokenDomains(val *[]*string)
	TokenDomainsInput() *[]*string
	VisibilityConfig() Wafv2WebAclVisibilityConfigOutputReference
	VisibilityConfigInput() *Wafv2WebAclVisibilityConfig
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
	PutAssociationConfig(value *Wafv2WebAclAssociationConfig)
	PutCaptchaConfig(value *Wafv2WebAclCaptchaConfig)
	PutChallengeConfig(value *Wafv2WebAclChallengeConfig)
	PutCustomResponseBody(value interface{})
	PutDefaultAction(value *Wafv2WebAclDefaultAction)
	PutRule(value interface{})
	PutVisibilityConfig(value *Wafv2WebAclVisibilityConfig)
	ResetAssociationConfig()
	ResetCaptchaConfig()
	ResetChallengeConfig()
	ResetCustomResponseBody()
	ResetDescription()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRule()
	ResetRuleJson()
	ResetTags()
	ResetTagsAll()
	ResetTokenDomains()
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

// The jsii proxy struct for Wafv2WebAcl
type jsiiProxy_Wafv2WebAcl struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Wafv2WebAcl) ApplicationIntegrationUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationIntegrationUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) AssociationConfig() Wafv2WebAclAssociationConfigOutputReference {
	var returns Wafv2WebAclAssociationConfigOutputReference
	_jsii_.Get(
		j,
		"associationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) AssociationConfigInput() *Wafv2WebAclAssociationConfig {
	var returns *Wafv2WebAclAssociationConfig
	_jsii_.Get(
		j,
		"associationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Capacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"capacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) CaptchaConfig() Wafv2WebAclCaptchaConfigOutputReference {
	var returns Wafv2WebAclCaptchaConfigOutputReference
	_jsii_.Get(
		j,
		"captchaConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) CaptchaConfigInput() *Wafv2WebAclCaptchaConfig {
	var returns *Wafv2WebAclCaptchaConfig
	_jsii_.Get(
		j,
		"captchaConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) ChallengeConfig() Wafv2WebAclChallengeConfigOutputReference {
	var returns Wafv2WebAclChallengeConfigOutputReference
	_jsii_.Get(
		j,
		"challengeConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) ChallengeConfigInput() *Wafv2WebAclChallengeConfig {
	var returns *Wafv2WebAclChallengeConfig
	_jsii_.Get(
		j,
		"challengeConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) CustomResponseBody() Wafv2WebAclCustomResponseBodyList {
	var returns Wafv2WebAclCustomResponseBodyList
	_jsii_.Get(
		j,
		"customResponseBody",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) CustomResponseBodyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customResponseBodyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) DefaultAction() Wafv2WebAclDefaultActionOutputReference {
	var returns Wafv2WebAclDefaultActionOutputReference
	_jsii_.Get(
		j,
		"defaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) DefaultActionInput() *Wafv2WebAclDefaultAction {
	var returns *Wafv2WebAclDefaultAction
	_jsii_.Get(
		j,
		"defaultActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) LockToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Rule() Wafv2WebAclRuleList {
	var returns Wafv2WebAclRuleList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) RuleJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) RuleJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) TokenDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) TokenDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) VisibilityConfig() Wafv2WebAclVisibilityConfigOutputReference {
	var returns Wafv2WebAclVisibilityConfigOutputReference
	_jsii_.Get(
		j,
		"visibilityConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Wafv2WebAcl) VisibilityConfigInput() *Wafv2WebAclVisibilityConfig {
	var returns *Wafv2WebAclVisibilityConfig
	_jsii_.Get(
		j,
		"visibilityConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/wafv2_web_acl aws_wafv2_web_acl} Resource.
func NewWafv2WebAcl(scope constructs.Construct, id *string, config *Wafv2WebAclConfig) Wafv2WebAcl {
	_init_.Initialize()

	if err := validateNewWafv2WebAclParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Wafv2WebAcl{}

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAcl",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/wafv2_web_acl aws_wafv2_web_acl} Resource.
func NewWafv2WebAcl_Override(w Wafv2WebAcl, scope constructs.Construct, id *string, config *Wafv2WebAclConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAcl",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetRuleJson(val *string) {
	if err := j.validateSetRuleJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ruleJson",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_Wafv2WebAcl)SetTokenDomains(val *[]*string) {
	if err := j.validateSetTokenDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenDomains",
		val,
	)
}

// Generates CDKTF code for importing a Wafv2WebAcl resource upon running "cdktf plan <stack-name>".
func Wafv2WebAcl_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWafv2WebAcl_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAcl",
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
func Wafv2WebAcl_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafv2WebAcl_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAcl",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Wafv2WebAcl_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafv2WebAcl_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAcl",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Wafv2WebAcl_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWafv2WebAcl_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAcl",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Wafv2WebAcl_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.wafv2WebAcl.Wafv2WebAcl",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_Wafv2WebAcl) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_Wafv2WebAcl) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAcl) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_Wafv2WebAcl) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_Wafv2WebAcl) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_Wafv2WebAcl) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_Wafv2WebAcl) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_Wafv2WebAcl) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_Wafv2WebAcl) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_Wafv2WebAcl) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAcl) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_Wafv2WebAcl) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) PutAssociationConfig(value *Wafv2WebAclAssociationConfig) {
	if err := w.validatePutAssociationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAssociationConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) PutCaptchaConfig(value *Wafv2WebAclCaptchaConfig) {
	if err := w.validatePutCaptchaConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCaptchaConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) PutChallengeConfig(value *Wafv2WebAclChallengeConfig) {
	if err := w.validatePutChallengeConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putChallengeConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) PutCustomResponseBody(value interface{}) {
	if err := w.validatePutCustomResponseBodyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putCustomResponseBody",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) PutDefaultAction(value *Wafv2WebAclDefaultAction) {
	if err := w.validatePutDefaultActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDefaultAction",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) PutRule(value interface{}) {
	if err := w.validatePutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRule",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) PutVisibilityConfig(value *Wafv2WebAclVisibilityConfig) {
	if err := w.validatePutVisibilityConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putVisibilityConfig",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetAssociationConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetAssociationConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetCaptchaConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetCaptchaConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetChallengeConfig() {
	_jsii_.InvokeVoid(
		w,
		"resetChallengeConfig",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetCustomResponseBody() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomResponseBody",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetDescription() {
	_jsii_.InvokeVoid(
		w,
		"resetDescription",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetRule() {
	_jsii_.InvokeVoid(
		w,
		"resetRule",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetRuleJson() {
	_jsii_.InvokeVoid(
		w,
		"resetRuleJson",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetTagsAll() {
	_jsii_.InvokeVoid(
		w,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) ResetTokenDomains() {
	_jsii_.InvokeVoid(
		w,
		"resetTokenDomains",
		nil, // no parameters
	)
}

func (w *jsiiProxy_Wafv2WebAcl) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAcl) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAcl) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAcl) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAcl) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_Wafv2WebAcl) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


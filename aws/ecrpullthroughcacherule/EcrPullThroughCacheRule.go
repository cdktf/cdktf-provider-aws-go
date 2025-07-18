// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrpullthroughcacherule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/ecrpullthroughcacherule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_pull_through_cache_rule aws_ecr_pull_through_cache_rule}.
type EcrPullThroughCacheRule interface {
	cdktf.TerraformResource
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
	CredentialArn() *string
	SetCredentialArn(val *string)
	CredentialArnInput() *string
	CustomRoleArn() *string
	SetCustomRoleArn(val *string)
	CustomRoleArnInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EcrRepositoryPrefix() *string
	SetEcrRepositoryPrefix(val *string)
	EcrRepositoryPrefixInput() *string
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
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	RegistryId() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpstreamRegistryUrl() *string
	SetUpstreamRegistryUrl(val *string)
	UpstreamRegistryUrlInput() *string
	UpstreamRepositoryPrefix() *string
	SetUpstreamRepositoryPrefix(val *string)
	UpstreamRepositoryPrefixInput() *string
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
	ResetCredentialArn()
	ResetCustomRoleArn()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegion()
	ResetUpstreamRepositoryPrefix()
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

// The jsii proxy struct for EcrPullThroughCacheRule
type jsiiProxy_EcrPullThroughCacheRule struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EcrPullThroughCacheRule) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) CredentialArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) CredentialArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) CustomRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) CustomRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) EcrRepositoryPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecrRepositoryPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) EcrRepositoryPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ecrRepositoryPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) RegistryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"registryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) UpstreamRegistryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upstreamRegistryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) UpstreamRegistryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upstreamRegistryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) UpstreamRepositoryPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upstreamRepositoryPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EcrPullThroughCacheRule) UpstreamRepositoryPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upstreamRepositoryPrefixInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_pull_through_cache_rule aws_ecr_pull_through_cache_rule} Resource.
func NewEcrPullThroughCacheRule(scope constructs.Construct, id *string, config *EcrPullThroughCacheRuleConfig) EcrPullThroughCacheRule {
	_init_.Initialize()

	if err := validateNewEcrPullThroughCacheRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EcrPullThroughCacheRule{}

	_jsii_.Create(
		"@cdktf/provider-aws.ecrPullThroughCacheRule.EcrPullThroughCacheRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_pull_through_cache_rule aws_ecr_pull_through_cache_rule} Resource.
func NewEcrPullThroughCacheRule_Override(e EcrPullThroughCacheRule, scope constructs.Construct, id *string, config *EcrPullThroughCacheRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.ecrPullThroughCacheRule.EcrPullThroughCacheRule",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetCredentialArn(val *string) {
	if err := j.validateSetCredentialArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialArn",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetCustomRoleArn(val *string) {
	if err := j.validateSetCustomRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customRoleArn",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetEcrRepositoryPrefix(val *string) {
	if err := j.validateSetEcrRepositoryPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ecrRepositoryPrefix",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetUpstreamRegistryUrl(val *string) {
	if err := j.validateSetUpstreamRegistryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upstreamRegistryUrl",
		val,
	)
}

func (j *jsiiProxy_EcrPullThroughCacheRule)SetUpstreamRepositoryPrefix(val *string) {
	if err := j.validateSetUpstreamRepositoryPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upstreamRepositoryPrefix",
		val,
	)
}

// Generates CDKTF code for importing a EcrPullThroughCacheRule resource upon running "cdktf plan <stack-name>".
func EcrPullThroughCacheRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEcrPullThroughCacheRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecrPullThroughCacheRule.EcrPullThroughCacheRule",
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
func EcrPullThroughCacheRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcrPullThroughCacheRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecrPullThroughCacheRule.EcrPullThroughCacheRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcrPullThroughCacheRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcrPullThroughCacheRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecrPullThroughCacheRule.EcrPullThroughCacheRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EcrPullThroughCacheRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEcrPullThroughCacheRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.ecrPullThroughCacheRule.EcrPullThroughCacheRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EcrPullThroughCacheRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.ecrPullThroughCacheRule.EcrPullThroughCacheRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) ResetCredentialArn() {
	_jsii_.InvokeVoid(
		e,
		"resetCredentialArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) ResetCustomRoleArn() {
	_jsii_.InvokeVoid(
		e,
		"resetCustomRoleArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) ResetRegion() {
	_jsii_.InvokeVoid(
		e,
		"resetRegion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) ResetUpstreamRepositoryPrefix() {
	_jsii_.InvokeVoid(
		e,
		"resetUpstreamRepositoryPrefix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EcrPullThroughCacheRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EcrPullThroughCacheRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


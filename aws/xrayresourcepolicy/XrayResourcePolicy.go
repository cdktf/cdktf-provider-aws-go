// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package xrayresourcepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/xrayresourcepolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/xray_resource_policy aws_xray_resource_policy}.
type XrayResourcePolicy interface {
	cdktf.TerraformResource
	BypassPolicyLockoutCheck() interface{}
	SetBypassPolicyLockoutCheck(val interface{})
	BypassPolicyLockoutCheckInput() interface{}
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	LastUpdatedTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PolicyDocument() *string
	SetPolicyDocument(val *string)
	PolicyDocumentInput() *string
	PolicyName() *string
	SetPolicyName(val *string)
	PolicyNameInput() *string
	PolicyRevisionId() *string
	SetPolicyRevisionId(val *string)
	PolicyRevisionIdInput() *string
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
	ResetBypassPolicyLockoutCheck()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPolicyRevisionId()
	ResetRegion()
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

// The jsii proxy struct for XrayResourcePolicy
type jsiiProxy_XrayResourcePolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_XrayResourcePolicy) BypassPolicyLockoutCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassPolicyLockoutCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) BypassPolicyLockoutCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassPolicyLockoutCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) LastUpdatedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) PolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) PolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) PolicyRevisionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyRevisionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) PolicyRevisionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyRevisionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_XrayResourcePolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/xray_resource_policy aws_xray_resource_policy} Resource.
func NewXrayResourcePolicy(scope constructs.Construct, id *string, config *XrayResourcePolicyConfig) XrayResourcePolicy {
	_init_.Initialize()

	if err := validateNewXrayResourcePolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_XrayResourcePolicy{}

	_jsii_.Create(
		"@cdktf/provider-aws.xrayResourcePolicy.XrayResourcePolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.8.0/docs/resources/xray_resource_policy aws_xray_resource_policy} Resource.
func NewXrayResourcePolicy_Override(x XrayResourcePolicy, scope constructs.Construct, id *string, config *XrayResourcePolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.xrayResourcePolicy.XrayResourcePolicy",
		[]interface{}{scope, id, config},
		x,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetBypassPolicyLockoutCheck(val interface{}) {
	if err := j.validateSetBypassPolicyLockoutCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bypassPolicyLockoutCheck",
		val,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetPolicyDocument(val *string) {
	if err := j.validateSetPolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetPolicyName(val *string) {
	if err := j.validateSetPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetPolicyRevisionId(val *string) {
	if err := j.validateSetPolicyRevisionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyRevisionId",
		val,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_XrayResourcePolicy)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTF code for importing a XrayResourcePolicy resource upon running "cdktf plan <stack-name>".
func XrayResourcePolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateXrayResourcePolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.xrayResourcePolicy.XrayResourcePolicy",
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
func XrayResourcePolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateXrayResourcePolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.xrayResourcePolicy.XrayResourcePolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func XrayResourcePolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateXrayResourcePolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.xrayResourcePolicy.XrayResourcePolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func XrayResourcePolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateXrayResourcePolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.xrayResourcePolicy.XrayResourcePolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func XrayResourcePolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.xrayResourcePolicy.XrayResourcePolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) AddMoveTarget(moveTarget *string) {
	if err := x.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (x *jsiiProxy_XrayResourcePolicy) AddOverride(path *string, value interface{}) {
	if err := x.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (x *jsiiProxy_XrayResourcePolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := x.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		x,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := x.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		x,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := x.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		x,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := x.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		x,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := x.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		x,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := x.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		x,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := x.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		x,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := x.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		x,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := x.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		x,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		x,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := x.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (x *jsiiProxy_XrayResourcePolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := x.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		x,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) MoveFromId(id *string) {
	if err := x.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"moveFromId",
		[]interface{}{id},
	)
}

func (x *jsiiProxy_XrayResourcePolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := x.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (x *jsiiProxy_XrayResourcePolicy) MoveToId(id *string) {
	if err := x.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"moveToId",
		[]interface{}{id},
	)
}

func (x *jsiiProxy_XrayResourcePolicy) OverrideLogicalId(newLogicalId *string) {
	if err := x.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		x,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (x *jsiiProxy_XrayResourcePolicy) ResetBypassPolicyLockoutCheck() {
	_jsii_.InvokeVoid(
		x,
		"resetBypassPolicyLockoutCheck",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XrayResourcePolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		x,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XrayResourcePolicy) ResetPolicyRevisionId() {
	_jsii_.InvokeVoid(
		x,
		"resetPolicyRevisionId",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XrayResourcePolicy) ResetRegion() {
	_jsii_.InvokeVoid(
		x,
		"resetRegion",
		nil, // no parameters
	)
}

func (x *jsiiProxy_XrayResourcePolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		x,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		x,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		x,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		x,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		x,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XrayResourcePolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		x,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


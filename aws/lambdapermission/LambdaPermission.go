// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdapermission

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/lambdapermission/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/lambda_permission aws_lambda_permission}.
type LambdaPermission interface {
	cdktf.TerraformResource
	Action() *string
	SetAction(val *string)
	ActionInput() *string
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
	EventSourceToken() *string
	SetEventSourceToken(val *string)
	EventSourceTokenInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FunctionName() *string
	SetFunctionName(val *string)
	FunctionNameInput() *string
	FunctionUrlAuthType() *string
	SetFunctionUrlAuthType(val *string)
	FunctionUrlAuthTypeInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Principal() *string
	SetPrincipal(val *string)
	PrincipalInput() *string
	PrincipalOrgId() *string
	SetPrincipalOrgId(val *string)
	PrincipalOrgIdInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	Qualifier() *string
	SetQualifier(val *string)
	QualifierInput() *string
	// Experimental.
	RawOverrides() interface{}
	Region() *string
	SetRegion(val *string)
	RegionInput() *string
	SourceAccount() *string
	SetSourceAccount(val *string)
	SourceAccountInput() *string
	SourceArn() *string
	SetSourceArn(val *string)
	SourceArnInput() *string
	StatementId() *string
	SetStatementId(val *string)
	StatementIdInput() *string
	StatementIdPrefix() *string
	SetStatementIdPrefix(val *string)
	StatementIdPrefixInput() *string
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
	ResetEventSourceToken()
	ResetFunctionUrlAuthType()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrincipalOrgId()
	ResetQualifier()
	ResetRegion()
	ResetSourceAccount()
	ResetSourceArn()
	ResetStatementId()
	ResetStatementIdPrefix()
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

// The jsii proxy struct for LambdaPermission
type jsiiProxy_LambdaPermission struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LambdaPermission) Action() *string {
	var returns *string
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) ActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) EventSourceToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) EventSourceTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) FunctionUrlAuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionUrlAuthType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) FunctionUrlAuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionUrlAuthTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Principal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) PrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) PrincipalOrgId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalOrgId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) PrincipalOrgIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"principalOrgIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Qualifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) QualifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"qualifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) SourceAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) SourceAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) SourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) SourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) StatementId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) StatementIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) StatementIdPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementIdPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) StatementIdPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statementIdPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaPermission) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/lambda_permission aws_lambda_permission} Resource.
func NewLambdaPermission(scope constructs.Construct, id *string, config *LambdaPermissionConfig) LambdaPermission {
	_init_.Initialize()

	if err := validateNewLambdaPermissionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaPermission{}

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaPermission.LambdaPermission",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/lambda_permission aws_lambda_permission} Resource.
func NewLambdaPermission_Override(l LambdaPermission, scope constructs.Construct, id *string, config *LambdaPermissionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.lambdaPermission.LambdaPermission",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LambdaPermission)SetAction(val *string) {
	if err := j.validateSetActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"action",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetEventSourceToken(val *string) {
	if err := j.validateSetEventSourceTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSourceToken",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetFunctionName(val *string) {
	if err := j.validateSetFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetFunctionUrlAuthType(val *string) {
	if err := j.validateSetFunctionUrlAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionUrlAuthType",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetPrincipal(val *string) {
	if err := j.validateSetPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principal",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetPrincipalOrgId(val *string) {
	if err := j.validateSetPrincipalOrgIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"principalOrgId",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetQualifier(val *string) {
	if err := j.validateSetQualifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qualifier",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetSourceAccount(val *string) {
	if err := j.validateSetSourceAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceAccount",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetSourceArn(val *string) {
	if err := j.validateSetSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceArn",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetStatementId(val *string) {
	if err := j.validateSetStatementIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementId",
		val,
	)
}

func (j *jsiiProxy_LambdaPermission)SetStatementIdPrefix(val *string) {
	if err := j.validateSetStatementIdPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statementIdPrefix",
		val,
	)
}

// Generates CDKTF code for importing a LambdaPermission resource upon running "cdktf plan <stack-name>".
func LambdaPermission_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLambdaPermission_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaPermission.LambdaPermission",
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
func LambdaPermission_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaPermission_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaPermission.LambdaPermission",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LambdaPermission_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaPermission_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaPermission.LambdaPermission",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LambdaPermission_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLambdaPermission_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.lambdaPermission.LambdaPermission",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LambdaPermission_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.lambdaPermission.LambdaPermission",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LambdaPermission) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LambdaPermission) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LambdaPermission) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LambdaPermission) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LambdaPermission) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LambdaPermission) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LambdaPermission) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LambdaPermission) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LambdaPermission) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LambdaPermission) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LambdaPermission) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LambdaPermission) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaPermission) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LambdaPermission) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LambdaPermission) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LambdaPermission) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LambdaPermission) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LambdaPermission) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LambdaPermission) ResetEventSourceToken() {
	_jsii_.InvokeVoid(
		l,
		"resetEventSourceToken",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetFunctionUrlAuthType() {
	_jsii_.InvokeVoid(
		l,
		"resetFunctionUrlAuthType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetPrincipalOrgId() {
	_jsii_.InvokeVoid(
		l,
		"resetPrincipalOrgId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetQualifier() {
	_jsii_.InvokeVoid(
		l,
		"resetQualifier",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetRegion() {
	_jsii_.InvokeVoid(
		l,
		"resetRegion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetSourceAccount() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceAccount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetSourceArn() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceArn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetStatementId() {
	_jsii_.InvokeVoid(
		l,
		"resetStatementId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) ResetStatementIdPrefix() {
	_jsii_.InvokeVoid(
		l,
		"resetStatementIdPrefix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LambdaPermission) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaPermission) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaPermission) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaPermission) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaPermission) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaPermission) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


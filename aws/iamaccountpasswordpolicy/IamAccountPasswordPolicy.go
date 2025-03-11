// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iamaccountpasswordpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/iamaccountpasswordpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/iam_account_password_policy aws_iam_account_password_policy}.
type IamAccountPasswordPolicy interface {
	cdktf.TerraformResource
	AllowUsersToChangePassword() interface{}
	SetAllowUsersToChangePassword(val interface{})
	AllowUsersToChangePasswordInput() interface{}
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
	ExpirePasswords() cdktf.IResolvable
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HardExpiry() interface{}
	SetHardExpiry(val interface{})
	HardExpiryInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxPasswordAge() *float64
	SetMaxPasswordAge(val *float64)
	MaxPasswordAgeInput() *float64
	MinimumPasswordLength() *float64
	SetMinimumPasswordLength(val *float64)
	MinimumPasswordLengthInput() *float64
	// The tree node.
	Node() constructs.Node
	PasswordReusePrevention() *float64
	SetPasswordReusePrevention(val *float64)
	PasswordReusePreventionInput() *float64
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
	RequireLowercaseCharacters() interface{}
	SetRequireLowercaseCharacters(val interface{})
	RequireLowercaseCharactersInput() interface{}
	RequireNumbers() interface{}
	SetRequireNumbers(val interface{})
	RequireNumbersInput() interface{}
	RequireSymbols() interface{}
	SetRequireSymbols(val interface{})
	RequireSymbolsInput() interface{}
	RequireUppercaseCharacters() interface{}
	SetRequireUppercaseCharacters(val interface{})
	RequireUppercaseCharactersInput() interface{}
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
	ResetAllowUsersToChangePassword()
	ResetHardExpiry()
	ResetId()
	ResetMaxPasswordAge()
	ResetMinimumPasswordLength()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordReusePrevention()
	ResetRequireLowercaseCharacters()
	ResetRequireNumbers()
	ResetRequireSymbols()
	ResetRequireUppercaseCharacters()
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

// The jsii proxy struct for IamAccountPasswordPolicy
type jsiiProxy_IamAccountPasswordPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IamAccountPasswordPolicy) AllowUsersToChangePassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUsersToChangePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) AllowUsersToChangePasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowUsersToChangePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) ExpirePasswords() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"expirePasswords",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) HardExpiry() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) HardExpiryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hardExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) MaxPasswordAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPasswordAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) MaxPasswordAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPasswordAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) MinimumPasswordLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumPasswordLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) MinimumPasswordLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumPasswordLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) PasswordReusePrevention() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordReusePrevention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) PasswordReusePreventionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"passwordReusePreventionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) RequireLowercaseCharacters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercaseCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) RequireLowercaseCharactersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLowercaseCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) RequireNumbers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) RequireNumbersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) RequireSymbols() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) RequireSymbolsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSymbolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) RequireUppercaseCharacters() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercaseCharacters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) RequireUppercaseCharactersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireUppercaseCharactersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IamAccountPasswordPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/iam_account_password_policy aws_iam_account_password_policy} Resource.
func NewIamAccountPasswordPolicy(scope constructs.Construct, id *string, config *IamAccountPasswordPolicyConfig) IamAccountPasswordPolicy {
	_init_.Initialize()

	if err := validateNewIamAccountPasswordPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IamAccountPasswordPolicy{}

	_jsii_.Create(
		"@cdktf/provider-aws.iamAccountPasswordPolicy.IamAccountPasswordPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/iam_account_password_policy aws_iam_account_password_policy} Resource.
func NewIamAccountPasswordPolicy_Override(i IamAccountPasswordPolicy, scope constructs.Construct, id *string, config *IamAccountPasswordPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.iamAccountPasswordPolicy.IamAccountPasswordPolicy",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetAllowUsersToChangePassword(val interface{}) {
	if err := j.validateSetAllowUsersToChangePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowUsersToChangePassword",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetHardExpiry(val interface{}) {
	if err := j.validateSetHardExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hardExpiry",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetMaxPasswordAge(val *float64) {
	if err := j.validateSetMaxPasswordAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPasswordAge",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetMinimumPasswordLength(val *float64) {
	if err := j.validateSetMinimumPasswordLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumPasswordLength",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetPasswordReusePrevention(val *float64) {
	if err := j.validateSetPasswordReusePreventionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordReusePrevention",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetRequireLowercaseCharacters(val interface{}) {
	if err := j.validateSetRequireLowercaseCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireLowercaseCharacters",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetRequireNumbers(val interface{}) {
	if err := j.validateSetRequireNumbersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireNumbers",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetRequireSymbols(val interface{}) {
	if err := j.validateSetRequireSymbolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireSymbols",
		val,
	)
}

func (j *jsiiProxy_IamAccountPasswordPolicy)SetRequireUppercaseCharacters(val interface{}) {
	if err := j.validateSetRequireUppercaseCharactersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireUppercaseCharacters",
		val,
	)
}

// Generates CDKTF code for importing a IamAccountPasswordPolicy resource upon running "cdktf plan <stack-name>".
func IamAccountPasswordPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIamAccountPasswordPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.iamAccountPasswordPolicy.IamAccountPasswordPolicy",
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
func IamAccountPasswordPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIamAccountPasswordPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.iamAccountPasswordPolicy.IamAccountPasswordPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IamAccountPasswordPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIamAccountPasswordPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.iamAccountPasswordPolicy.IamAccountPasswordPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IamAccountPasswordPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIamAccountPasswordPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.iamAccountPasswordPolicy.IamAccountPasswordPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IamAccountPasswordPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.iamAccountPasswordPolicy.IamAccountPasswordPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) MoveFromId(id *string) {
	if err := i.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveFromId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) MoveToId(id *string) {
	if err := i.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveToId",
		[]interface{}{id},
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ResetAllowUsersToChangePassword() {
	_jsii_.InvokeVoid(
		i,
		"resetAllowUsersToChangePassword",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ResetHardExpiry() {
	_jsii_.InvokeVoid(
		i,
		"resetHardExpiry",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ResetMaxPasswordAge() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxPasswordAge",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ResetMinimumPasswordLength() {
	_jsii_.InvokeVoid(
		i,
		"resetMinimumPasswordLength",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ResetPasswordReusePrevention() {
	_jsii_.InvokeVoid(
		i,
		"resetPasswordReusePrevention",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ResetRequireLowercaseCharacters() {
	_jsii_.InvokeVoid(
		i,
		"resetRequireLowercaseCharacters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ResetRequireNumbers() {
	_jsii_.InvokeVoid(
		i,
		"resetRequireNumbers",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ResetRequireSymbols() {
	_jsii_.InvokeVoid(
		i,
		"resetRequireSymbols",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ResetRequireUppercaseCharacters() {
	_jsii_.InvokeVoid(
		i,
		"resetRequireUppercaseCharacters",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IamAccountPasswordPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IamAccountPasswordPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


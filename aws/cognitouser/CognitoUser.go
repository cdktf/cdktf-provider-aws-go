// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cognitouser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/cognitouser/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/cognito_user aws_cognito_user}.
type CognitoUser interface {
	cdktf.TerraformResource
	Attributes() *map[string]*string
	SetAttributes(val *map[string]*string)
	AttributesInput() *map[string]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientMetadata() *map[string]*string
	SetClientMetadata(val *map[string]*string)
	ClientMetadataInput() *map[string]*string
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
	CreationDate() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DesiredDeliveryMediums() *[]*string
	SetDesiredDeliveryMediums(val *[]*string)
	DesiredDeliveryMediumsInput() *[]*string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ForceAliasCreation() interface{}
	SetForceAliasCreation(val interface{})
	ForceAliasCreationInput() interface{}
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
	LastModifiedDate() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MessageAction() *string
	SetMessageAction(val *string)
	MessageActionInput() *string
	MfaSettingList() *[]*string
	// The tree node.
	Node() constructs.Node
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PreferredMfaSetting() *string
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
	Status() *string
	Sub() *string
	TemporaryPassword() *string
	SetTemporaryPassword(val *string)
	TemporaryPasswordInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	UserPoolId() *string
	SetUserPoolId(val *string)
	UserPoolIdInput() *string
	ValidationData() *map[string]*string
	SetValidationData(val *map[string]*string)
	ValidationDataInput() *map[string]*string
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
	ResetAttributes()
	ResetClientMetadata()
	ResetDesiredDeliveryMediums()
	ResetEnabled()
	ResetForceAliasCreation()
	ResetId()
	ResetMessageAction()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPassword()
	ResetTemporaryPassword()
	ResetValidationData()
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

// The jsii proxy struct for CognitoUser
type jsiiProxy_CognitoUser struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CognitoUser) Attributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) AttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) ClientMetadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"clientMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) ClientMetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"clientMetadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) DesiredDeliveryMediums() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"desiredDeliveryMediums",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) DesiredDeliveryMediumsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"desiredDeliveryMediumsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) ForceAliasCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceAliasCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) ForceAliasCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceAliasCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) LastModifiedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) MessageAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) MessageActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"messageActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) MfaSettingList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"mfaSettingList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) PreferredMfaSetting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMfaSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Sub() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) TemporaryPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporaryPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) TemporaryPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"temporaryPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) UserPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) UserPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) ValidationData() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"validationData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CognitoUser) ValidationDataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"validationDataInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/cognito_user aws_cognito_user} Resource.
func NewCognitoUser(scope constructs.Construct, id *string, config *CognitoUserConfig) CognitoUser {
	_init_.Initialize()

	if err := validateNewCognitoUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CognitoUser{}

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUser.CognitoUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.76.0/docs/resources/cognito_user aws_cognito_user} Resource.
func NewCognitoUser_Override(c CognitoUser, scope constructs.Construct, id *string, config *CognitoUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cognitoUser.CognitoUser",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CognitoUser)SetAttributes(val *map[string]*string) {
	if err := j.validateSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetClientMetadata(val *map[string]*string) {
	if err := j.validateSetClientMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientMetadata",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetDesiredDeliveryMediums(val *[]*string) {
	if err := j.validateSetDesiredDeliveryMediumsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"desiredDeliveryMediums",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetForceAliasCreation(val interface{}) {
	if err := j.validateSetForceAliasCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceAliasCreation",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetMessageAction(val *string) {
	if err := j.validateSetMessageActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageAction",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetTemporaryPassword(val *string) {
	if err := j.validateSetTemporaryPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"temporaryPassword",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetUserPoolId(val *string) {
	if err := j.validateSetUserPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPoolId",
		val,
	)
}

func (j *jsiiProxy_CognitoUser)SetValidationData(val *map[string]*string) {
	if err := j.validateSetValidationDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validationData",
		val,
	)
}

// Generates CDKTF code for importing a CognitoUser resource upon running "cdktf plan <stack-name>".
func CognitoUser_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCognitoUser_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoUser.CognitoUser",
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
func CognitoUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoUser.CognitoUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoUser_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoUser.CognitoUser",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CognitoUser_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCognitoUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.cognitoUser.CognitoUser",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CognitoUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.cognitoUser.CognitoUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CognitoUser) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CognitoUser) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CognitoUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CognitoUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CognitoUser) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CognitoUser) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CognitoUser) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CognitoUser) ResetAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUser) ResetClientMetadata() {
	_jsii_.InvokeVoid(
		c,
		"resetClientMetadata",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUser) ResetDesiredDeliveryMediums() {
	_jsii_.InvokeVoid(
		c,
		"resetDesiredDeliveryMediums",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUser) ResetEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUser) ResetForceAliasCreation() {
	_jsii_.InvokeVoid(
		c,
		"resetForceAliasCreation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUser) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUser) ResetMessageAction() {
	_jsii_.InvokeVoid(
		c,
		"resetMessageAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUser) ResetPassword() {
	_jsii_.InvokeVoid(
		c,
		"resetPassword",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUser) ResetTemporaryPassword() {
	_jsii_.InvokeVoid(
		c,
		"resetTemporaryPassword",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUser) ResetValidationData() {
	_jsii_.InvokeVoid(
		c,
		"resetValidationData",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CognitoUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CognitoUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


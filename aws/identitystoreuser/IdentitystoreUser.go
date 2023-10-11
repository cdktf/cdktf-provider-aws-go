// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitystoreuser

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v17/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v17/identitystoreuser/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/identitystore_user aws_identitystore_user}.
type IdentitystoreUser interface {
	cdktf.TerraformResource
	Addresses() IdentitystoreUserAddressesOutputReference
	AddressesInput() *IdentitystoreUserAddresses
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
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
	Emails() IdentitystoreUserEmailsOutputReference
	EmailsInput() *IdentitystoreUserEmails
	ExternalIds() IdentitystoreUserExternalIdsList
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
	IdentityStoreId() *string
	SetIdentityStoreId(val *string)
	IdentityStoreIdInput() *string
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locale() *string
	SetLocale(val *string)
	LocaleInput() *string
	Name() IdentitystoreUserNameOutputReference
	NameInput() *IdentitystoreUserName
	Nickname() *string
	SetNickname(val *string)
	NicknameInput() *string
	// The tree node.
	Node() constructs.Node
	PhoneNumbers() IdentitystoreUserPhoneNumbersOutputReference
	PhoneNumbersInput() *IdentitystoreUserPhoneNumbers
	PreferredLanguage() *string
	SetPreferredLanguage(val *string)
	PreferredLanguageInput() *string
	ProfileUrl() *string
	SetProfileUrl(val *string)
	ProfileUrlInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	Title() *string
	SetTitle(val *string)
	TitleInput() *string
	UserId() *string
	UserName() *string
	SetUserName(val *string)
	UserNameInput() *string
	UserType() *string
	SetUserType(val *string)
	UserTypeInput() *string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAddresses(value *IdentitystoreUserAddresses)
	PutEmails(value *IdentitystoreUserEmails)
	PutName(value *IdentitystoreUserName)
	PutPhoneNumbers(value *IdentitystoreUserPhoneNumbers)
	ResetAddresses()
	ResetEmails()
	ResetId()
	ResetLocale()
	ResetNickname()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPhoneNumbers()
	ResetPreferredLanguage()
	ResetProfileUrl()
	ResetTimezone()
	ResetTitle()
	ResetUserType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for IdentitystoreUser
type jsiiProxy_IdentitystoreUser struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IdentitystoreUser) Addresses() IdentitystoreUserAddressesOutputReference {
	var returns IdentitystoreUserAddressesOutputReference
	_jsii_.Get(
		j,
		"addresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) AddressesInput() *IdentitystoreUserAddresses {
	var returns *IdentitystoreUserAddresses
	_jsii_.Get(
		j,
		"addressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Emails() IdentitystoreUserEmailsOutputReference {
	var returns IdentitystoreUserEmailsOutputReference
	_jsii_.Get(
		j,
		"emails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) EmailsInput() *IdentitystoreUserEmails {
	var returns *IdentitystoreUserEmails
	_jsii_.Get(
		j,
		"emailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) ExternalIds() IdentitystoreUserExternalIdsList {
	var returns IdentitystoreUserExternalIdsList
	_jsii_.Get(
		j,
		"externalIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) IdentityStoreId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityStoreId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) IdentityStoreIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityStoreIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Name() IdentitystoreUserNameOutputReference {
	var returns IdentitystoreUserNameOutputReference
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) NameInput() *IdentitystoreUserName {
	var returns *IdentitystoreUserName
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Nickname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nickname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) NicknameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nicknameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) PhoneNumbers() IdentitystoreUserPhoneNumbersOutputReference {
	var returns IdentitystoreUserPhoneNumbersOutputReference
	_jsii_.Get(
		j,
		"phoneNumbers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) PhoneNumbersInput() *IdentitystoreUserPhoneNumbers {
	var returns *IdentitystoreUserPhoneNumbers
	_jsii_.Get(
		j,
		"phoneNumbersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) PreferredLanguage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) PreferredLanguageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredLanguageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) ProfileUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) ProfileUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) TitleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"titleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) UserId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) UserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) UserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) UserType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentitystoreUser) UserTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/identitystore_user aws_identitystore_user} Resource.
func NewIdentitystoreUser(scope constructs.Construct, id *string, config *IdentitystoreUserConfig) IdentitystoreUser {
	_init_.Initialize()

	if err := validateNewIdentitystoreUserParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentitystoreUser{}

	_jsii_.Create(
		"@cdktf/provider-aws.identitystoreUser.IdentitystoreUser",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/identitystore_user aws_identitystore_user} Resource.
func NewIdentitystoreUser_Override(i IdentitystoreUser, scope constructs.Construct, id *string, config *IdentitystoreUserConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.identitystoreUser.IdentitystoreUser",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetIdentityStoreId(val *string) {
	if err := j.validateSetIdentityStoreIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityStoreId",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetNickname(val *string) {
	if err := j.validateSetNicknameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nickname",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetPreferredLanguage(val *string) {
	if err := j.validateSetPreferredLanguageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredLanguage",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetProfileUrl(val *string) {
	if err := j.validateSetProfileUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profileUrl",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetTitle(val *string) {
	if err := j.validateSetTitleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"title",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetUserName(val *string) {
	if err := j.validateSetUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userName",
		val,
	)
}

func (j *jsiiProxy_IdentitystoreUser)SetUserType(val *string) {
	if err := j.validateSetUserTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userType",
		val,
	)
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
func IdentitystoreUser_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentitystoreUser_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.identitystoreUser.IdentitystoreUser",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentitystoreUser_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentitystoreUser_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.identitystoreUser.IdentitystoreUser",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentitystoreUser_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentitystoreUser_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.identitystoreUser.IdentitystoreUser",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IdentitystoreUser_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.identitystoreUser.IdentitystoreUser",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IdentitystoreUser) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IdentitystoreUser) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IdentitystoreUser) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IdentitystoreUser) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IdentitystoreUser) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IdentitystoreUser) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IdentitystoreUser) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IdentitystoreUser) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IdentitystoreUser) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IdentitystoreUser) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IdentitystoreUser) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IdentitystoreUser) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IdentitystoreUser) PutAddresses(value *IdentitystoreUserAddresses) {
	if err := i.validatePutAddressesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAddresses",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IdentitystoreUser) PutEmails(value *IdentitystoreUserEmails) {
	if err := i.validatePutEmailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putEmails",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IdentitystoreUser) PutName(value *IdentitystoreUserName) {
	if err := i.validatePutNameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putName",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IdentitystoreUser) PutPhoneNumbers(value *IdentitystoreUserPhoneNumbers) {
	if err := i.validatePutPhoneNumbersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putPhoneNumbers",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetAddresses() {
	_jsii_.InvokeVoid(
		i,
		"resetAddresses",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetEmails() {
	_jsii_.InvokeVoid(
		i,
		"resetEmails",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetLocale() {
	_jsii_.InvokeVoid(
		i,
		"resetLocale",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetNickname() {
	_jsii_.InvokeVoid(
		i,
		"resetNickname",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetPhoneNumbers() {
	_jsii_.InvokeVoid(
		i,
		"resetPhoneNumbers",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetPreferredLanguage() {
	_jsii_.InvokeVoid(
		i,
		"resetPreferredLanguage",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetProfileUrl() {
	_jsii_.InvokeVoid(
		i,
		"resetProfileUrl",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetTimezone() {
	_jsii_.InvokeVoid(
		i,
		"resetTimezone",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetTitle() {
	_jsii_.InvokeVoid(
		i,
		"resetTitle",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) ResetUserType() {
	_jsii_.InvokeVoid(
		i,
		"resetUserType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentitystoreUser) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentitystoreUser) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentitystoreUser) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentitystoreUser) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package accountprimarycontact

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/accountprimarycontact/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/account_primary_contact aws_account_primary_contact}.
type AccountPrimaryContact interface {
	cdktf.TerraformResource
	AccountId() *string
	SetAccountId(val *string)
	AccountIdInput() *string
	AddressLine1() *string
	SetAddressLine1(val *string)
	AddressLine1Input() *string
	AddressLine2() *string
	SetAddressLine2(val *string)
	AddressLine2Input() *string
	AddressLine3() *string
	SetAddressLine3(val *string)
	AddressLine3Input() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	City() *string
	SetCity(val *string)
	CityInput() *string
	CompanyName() *string
	SetCompanyName(val *string)
	CompanyNameInput() *string
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
	CountryCode() *string
	SetCountryCode(val *string)
	CountryCodeInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DistrictOrCounty() *string
	SetDistrictOrCounty(val *string)
	DistrictOrCountyInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullName() *string
	SetFullName(val *string)
	FullNameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	PhoneNumber() *string
	SetPhoneNumber(val *string)
	PhoneNumberInput() *string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
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
	StateOrRegion() *string
	SetStateOrRegion(val *string)
	StateOrRegionInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WebsiteUrl() *string
	SetWebsiteUrl(val *string)
	WebsiteUrlInput() *string
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
	ResetAccountId()
	ResetAddressLine2()
	ResetAddressLine3()
	ResetCompanyName()
	ResetDistrictOrCounty()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetStateOrRegion()
	ResetWebsiteUrl()
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

// The jsii proxy struct for AccountPrimaryContact
type jsiiProxy_AccountPrimaryContact struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AccountPrimaryContact) AccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) AccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) AddressLine1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) AddressLine1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) AddressLine2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) AddressLine2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) AddressLine3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) AddressLine3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) CompanyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) CompanyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) CountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) DistrictOrCounty() *string {
	var returns *string
	_jsii_.Get(
		j,
		"districtOrCounty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) DistrictOrCountyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"districtOrCountyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) FullNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) StateOrRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateOrRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) StateOrRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateOrRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) WebsiteUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AccountPrimaryContact) WebsiteUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"websiteUrlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/account_primary_contact aws_account_primary_contact} Resource.
func NewAccountPrimaryContact(scope constructs.Construct, id *string, config *AccountPrimaryContactConfig) AccountPrimaryContact {
	_init_.Initialize()

	if err := validateNewAccountPrimaryContactParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AccountPrimaryContact{}

	_jsii_.Create(
		"@cdktf/provider-aws.accountPrimaryContact.AccountPrimaryContact",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/account_primary_contact aws_account_primary_contact} Resource.
func NewAccountPrimaryContact_Override(a AccountPrimaryContact, scope constructs.Construct, id *string, config *AccountPrimaryContactConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.accountPrimaryContact.AccountPrimaryContact",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetAccountId(val *string) {
	if err := j.validateSetAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountId",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetAddressLine1(val *string) {
	if err := j.validateSetAddressLine1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine1",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetAddressLine2(val *string) {
	if err := j.validateSetAddressLine2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine2",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetAddressLine3(val *string) {
	if err := j.validateSetAddressLine3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine3",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetCompanyName(val *string) {
	if err := j.validateSetCompanyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"companyName",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetCountryCode(val *string) {
	if err := j.validateSetCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countryCode",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetDistrictOrCounty(val *string) {
	if err := j.validateSetDistrictOrCountyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"districtOrCounty",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetFullName(val *string) {
	if err := j.validateSetFullNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullName",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetStateOrRegion(val *string) {
	if err := j.validateSetStateOrRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stateOrRegion",
		val,
	)
}

func (j *jsiiProxy_AccountPrimaryContact)SetWebsiteUrl(val *string) {
	if err := j.validateSetWebsiteUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websiteUrl",
		val,
	)
}

// Generates CDKTF code for importing a AccountPrimaryContact resource upon running "cdktf plan <stack-name>".
func AccountPrimaryContact_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAccountPrimaryContact_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.accountPrimaryContact.AccountPrimaryContact",
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
func AccountPrimaryContact_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccountPrimaryContact_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.accountPrimaryContact.AccountPrimaryContact",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccountPrimaryContact_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccountPrimaryContact_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.accountPrimaryContact.AccountPrimaryContact",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AccountPrimaryContact_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAccountPrimaryContact_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.accountPrimaryContact.AccountPrimaryContact",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AccountPrimaryContact_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.accountPrimaryContact.AccountPrimaryContact",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AccountPrimaryContact) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AccountPrimaryContact) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AccountPrimaryContact) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AccountPrimaryContact) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AccountPrimaryContact) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AccountPrimaryContact) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AccountPrimaryContact) ResetAccountId() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountPrimaryContact) ResetAddressLine2() {
	_jsii_.InvokeVoid(
		a,
		"resetAddressLine2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountPrimaryContact) ResetAddressLine3() {
	_jsii_.InvokeVoid(
		a,
		"resetAddressLine3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountPrimaryContact) ResetCompanyName() {
	_jsii_.InvokeVoid(
		a,
		"resetCompanyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountPrimaryContact) ResetDistrictOrCounty() {
	_jsii_.InvokeVoid(
		a,
		"resetDistrictOrCounty",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountPrimaryContact) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountPrimaryContact) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountPrimaryContact) ResetStateOrRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetStateOrRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountPrimaryContact) ResetWebsiteUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetWebsiteUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AccountPrimaryContact) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AccountPrimaryContact) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


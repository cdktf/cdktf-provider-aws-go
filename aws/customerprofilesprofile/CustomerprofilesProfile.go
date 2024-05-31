// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/customerprofilesprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/customerprofiles_profile aws_customerprofiles_profile}.
type CustomerprofilesProfile interface {
	cdktf.TerraformResource
	AccountNumber() *string
	SetAccountNumber(val *string)
	AccountNumberInput() *string
	AdditionalInformation() *string
	SetAdditionalInformation(val *string)
	AdditionalInformationInput() *string
	Address() CustomerprofilesProfileAddressOutputReference
	AddressInput() *CustomerprofilesProfileAddress
	Attributes() *map[string]*string
	SetAttributes(val *map[string]*string)
	AttributesInput() *map[string]*string
	BillingAddress() CustomerprofilesProfileBillingAddressOutputReference
	BillingAddressInput() *CustomerprofilesProfileBillingAddress
	BirthDate() *string
	SetBirthDate(val *string)
	BirthDateInput() *string
	BusinessEmailAddress() *string
	SetBusinessEmailAddress(val *string)
	BusinessEmailAddressInput() *string
	BusinessName() *string
	SetBusinessName(val *string)
	BusinessNameInput() *string
	BusinessPhoneNumber() *string
	SetBusinessPhoneNumber(val *string)
	BusinessPhoneNumberInput() *string
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
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	EmailAddress() *string
	SetEmailAddress(val *string)
	EmailAddressInput() *string
	FirstName() *string
	SetFirstName(val *string)
	FirstNameInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GenderString() *string
	SetGenderString(val *string)
	GenderStringInput() *string
	HomePhoneNumber() *string
	SetHomePhoneNumber(val *string)
	HomePhoneNumberInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LastName() *string
	SetLastName(val *string)
	LastNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MailingAddress() CustomerprofilesProfileMailingAddressOutputReference
	MailingAddressInput() *CustomerprofilesProfileMailingAddress
	MiddleName() *string
	SetMiddleName(val *string)
	MiddleNameInput() *string
	MobilePhoneNumber() *string
	SetMobilePhoneNumber(val *string)
	MobilePhoneNumberInput() *string
	// The tree node.
	Node() constructs.Node
	PartyTypeString() *string
	SetPartyTypeString(val *string)
	PartyTypeStringInput() *string
	PersonalEmailAddress() *string
	SetPersonalEmailAddress(val *string)
	PersonalEmailAddressInput() *string
	PhoneNumber() *string
	SetPhoneNumber(val *string)
	PhoneNumberInput() *string
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
	ShippingAddress() CustomerprofilesProfileShippingAddressOutputReference
	ShippingAddressInput() *CustomerprofilesProfileShippingAddress
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
	PutAddress(value *CustomerprofilesProfileAddress)
	PutBillingAddress(value *CustomerprofilesProfileBillingAddress)
	PutMailingAddress(value *CustomerprofilesProfileMailingAddress)
	PutShippingAddress(value *CustomerprofilesProfileShippingAddress)
	ResetAccountNumber()
	ResetAdditionalInformation()
	ResetAddress()
	ResetAttributes()
	ResetBillingAddress()
	ResetBirthDate()
	ResetBusinessEmailAddress()
	ResetBusinessName()
	ResetBusinessPhoneNumber()
	ResetEmailAddress()
	ResetFirstName()
	ResetGenderString()
	ResetHomePhoneNumber()
	ResetId()
	ResetLastName()
	ResetMailingAddress()
	ResetMiddleName()
	ResetMobilePhoneNumber()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartyTypeString()
	ResetPersonalEmailAddress()
	ResetPhoneNumber()
	ResetShippingAddress()
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

// The jsii proxy struct for CustomerprofilesProfile
type jsiiProxy_CustomerprofilesProfile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_CustomerprofilesProfile) AccountNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) AccountNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) AdditionalInformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) AdditionalInformationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalInformationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) Address() CustomerprofilesProfileAddressOutputReference {
	var returns CustomerprofilesProfileAddressOutputReference
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) AddressInput() *CustomerprofilesProfileAddress {
	var returns *CustomerprofilesProfileAddress
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) Attributes() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) AttributesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"attributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) BillingAddress() CustomerprofilesProfileBillingAddressOutputReference {
	var returns CustomerprofilesProfileBillingAddressOutputReference
	_jsii_.Get(
		j,
		"billingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) BillingAddressInput() *CustomerprofilesProfileBillingAddress {
	var returns *CustomerprofilesProfileBillingAddress
	_jsii_.Get(
		j,
		"billingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) BirthDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"birthDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) BirthDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"birthDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) BusinessEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) BusinessEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) BusinessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) BusinessNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) BusinessPhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessPhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) BusinessPhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"businessPhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) EmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) EmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) GenderString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"genderString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) GenderStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"genderStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) HomePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) HomePhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) MailingAddress() CustomerprofilesProfileMailingAddressOutputReference {
	var returns CustomerprofilesProfileMailingAddressOutputReference
	_jsii_.Get(
		j,
		"mailingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) MailingAddressInput() *CustomerprofilesProfileMailingAddress {
	var returns *CustomerprofilesProfileMailingAddress
	_jsii_.Get(
		j,
		"mailingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) MiddleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) MiddleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) MobilePhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) MobilePhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobilePhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) PartyTypeString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partyTypeString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) PartyTypeStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partyTypeStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) PersonalEmailAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"personalEmailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) PersonalEmailAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"personalEmailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) ShippingAddress() CustomerprofilesProfileShippingAddressOutputReference {
	var returns CustomerprofilesProfileShippingAddressOutputReference
	_jsii_.Get(
		j,
		"shippingAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) ShippingAddressInput() *CustomerprofilesProfileShippingAddress {
	var returns *CustomerprofilesProfileShippingAddress
	_jsii_.Get(
		j,
		"shippingAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/customerprofiles_profile aws_customerprofiles_profile} Resource.
func NewCustomerprofilesProfile(scope constructs.Construct, id *string, config *CustomerprofilesProfileConfig) CustomerprofilesProfile {
	_init_.Initialize()

	if err := validateNewCustomerprofilesProfileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesProfile{}

	_jsii_.Create(
		"@cdktf/provider-aws.customerprofilesProfile.CustomerprofilesProfile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/5.52.0/docs/resources/customerprofiles_profile aws_customerprofiles_profile} Resource.
func NewCustomerprofilesProfile_Override(c CustomerprofilesProfile, scope constructs.Construct, id *string, config *CustomerprofilesProfileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.customerprofilesProfile.CustomerprofilesProfile",
		[]interface{}{scope, id, config},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetAccountNumber(val *string) {
	if err := j.validateSetAccountNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountNumber",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetAdditionalInformation(val *string) {
	if err := j.validateSetAdditionalInformationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalInformation",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetAttributes(val *map[string]*string) {
	if err := j.validateSetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributes",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetBirthDate(val *string) {
	if err := j.validateSetBirthDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"birthDate",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetBusinessEmailAddress(val *string) {
	if err := j.validateSetBusinessEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessEmailAddress",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetBusinessName(val *string) {
	if err := j.validateSetBusinessNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessName",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetBusinessPhoneNumber(val *string) {
	if err := j.validateSetBusinessPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"businessPhoneNumber",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetEmailAddress(val *string) {
	if err := j.validateSetEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddress",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetGenderString(val *string) {
	if err := j.validateSetGenderStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"genderString",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetHomePhoneNumber(val *string) {
	if err := j.validateSetHomePhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"homePhoneNumber",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetMiddleName(val *string) {
	if err := j.validateSetMiddleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"middleName",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetMobilePhoneNumber(val *string) {
	if err := j.validateSetMobilePhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobilePhoneNumber",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetPartyTypeString(val *string) {
	if err := j.validateSetPartyTypeStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partyTypeString",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetPersonalEmailAddress(val *string) {
	if err := j.validateSetPersonalEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"personalEmailAddress",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a CustomerprofilesProfile resource upon running "cdktf plan <stack-name>".
func CustomerprofilesProfile_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateCustomerprofilesProfile_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.customerprofilesProfile.CustomerprofilesProfile",
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
func CustomerprofilesProfile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomerprofilesProfile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.customerprofilesProfile.CustomerprofilesProfile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CustomerprofilesProfile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomerprofilesProfile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.customerprofilesProfile.CustomerprofilesProfile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func CustomerprofilesProfile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCustomerprofilesProfile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-aws.customerprofilesProfile.CustomerprofilesProfile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CustomerprofilesProfile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-aws.customerprofilesProfile.CustomerprofilesProfile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CustomerprofilesProfile) AddMoveTarget(moveTarget *string) {
	if err := c.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) AddOverride(path *string, value interface{}) {
	if err := c.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CustomerprofilesProfile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesProfile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CustomerprofilesProfile) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CustomerprofilesProfile) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CustomerprofilesProfile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CustomerprofilesProfile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CustomerprofilesProfile) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CustomerprofilesProfile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CustomerprofilesProfile) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesProfile) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := c.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesProfile) MoveFromId(id *string) {
	if err := c.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveFromId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) MoveTo(moveTarget *string, index interface{}) {
	if err := c.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) MoveToId(id *string) {
	if err := c.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"moveToId",
		[]interface{}{id},
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) OverrideLogicalId(newLogicalId *string) {
	if err := c.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) PutAddress(value *CustomerprofilesProfileAddress) {
	if err := c.validatePutAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) PutBillingAddress(value *CustomerprofilesProfileBillingAddress) {
	if err := c.validatePutBillingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBillingAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) PutMailingAddress(value *CustomerprofilesProfileMailingAddress) {
	if err := c.validatePutMailingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putMailingAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) PutShippingAddress(value *CustomerprofilesProfileShippingAddress) {
	if err := c.validatePutShippingAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putShippingAddress",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetAccountNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetAccountNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetAdditionalInformation() {
	_jsii_.InvokeVoid(
		c,
		"resetAdditionalInformation",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetAttributes() {
	_jsii_.InvokeVoid(
		c,
		"resetAttributes",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetBillingAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetBillingAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetBirthDate() {
	_jsii_.InvokeVoid(
		c,
		"resetBirthDate",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetBusinessEmailAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetBusinessEmailAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetBusinessName() {
	_jsii_.InvokeVoid(
		c,
		"resetBusinessName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetBusinessPhoneNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetBusinessPhoneNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetEmailAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetEmailAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetFirstName() {
	_jsii_.InvokeVoid(
		c,
		"resetFirstName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetGenderString() {
	_jsii_.InvokeVoid(
		c,
		"resetGenderString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetHomePhoneNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetHomePhoneNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetId() {
	_jsii_.InvokeVoid(
		c,
		"resetId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetLastName() {
	_jsii_.InvokeVoid(
		c,
		"resetLastName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetMailingAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetMailingAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetMiddleName() {
	_jsii_.InvokeVoid(
		c,
		"resetMiddleName",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetMobilePhoneNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetMobilePhoneNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		c,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetPartyTypeString() {
	_jsii_.InvokeVoid(
		c,
		"resetPartyTypeString",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetPersonalEmailAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetPersonalEmailAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		c,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) ResetShippingAddress() {
	_jsii_.InvokeVoid(
		c,
		"resetShippingAddress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesProfile) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesProfile) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesProfile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesProfile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesProfile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53domainsdomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v21/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v21/route53domainsdomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53DomainsDomainTechContactOutputReference interface {
	cdktf.ComplexObject
	AddressLine1() *string
	SetAddressLine1(val *string)
	AddressLine1Input() *string
	AddressLine2() *string
	SetAddressLine2(val *string)
	AddressLine2Input() *string
	City() *string
	SetCity(val *string)
	CityInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ContactType() *string
	SetContactType(val *string)
	ContactTypeInput() *string
	CountryCode() *string
	SetCountryCode(val *string)
	CountryCodeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	ExtraParam() Route53DomainsDomainTechContactExtraParamList
	ExtraParamInput() interface{}
	Fax() *string
	SetFax(val *string)
	FaxInput() *string
	FirstName() *string
	SetFirstName(val *string)
	FirstNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LastName() *string
	SetLastName(val *string)
	LastNameInput() *string
	OrganizationName() *string
	SetOrganizationName(val *string)
	OrganizationNameInput() *string
	PhoneNumber() *string
	SetPhoneNumber(val *string)
	PhoneNumberInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ZipCode() *string
	SetZipCode(val *string)
	ZipCodeInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutExtraParam(value interface{})
	ResetAddressLine1()
	ResetAddressLine2()
	ResetCity()
	ResetContactType()
	ResetCountryCode()
	ResetEmail()
	ResetExtraParam()
	ResetFax()
	ResetFirstName()
	ResetLastName()
	ResetOrganizationName()
	ResetPhoneNumber()
	ResetState()
	ResetZipCode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Route53DomainsDomainTechContactOutputReference
type jsiiProxy_Route53DomainsDomainTechContactOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) AddressLine1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) AddressLine1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) AddressLine2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) AddressLine2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ContactType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ContactTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) CountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ExtraParam() Route53DomainsDomainTechContactExtraParamList {
	var returns Route53DomainsDomainTechContactExtraParamList
	_jsii_.Get(
		j,
		"extraParam",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ExtraParamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extraParamInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) Fax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) FaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) OrganizationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ZipCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ZipCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCodeInput",
		&returns,
	)
	return returns
}


func NewRoute53DomainsDomainTechContactOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Route53DomainsDomainTechContactOutputReference {
	_init_.Initialize()

	if err := validateNewRoute53DomainsDomainTechContactOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53DomainsDomainTechContactOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53DomainsDomain.Route53DomainsDomainTechContactOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRoute53DomainsDomainTechContactOutputReference_Override(r Route53DomainsDomainTechContactOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53DomainsDomain.Route53DomainsDomainTechContactOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetAddressLine1(val *string) {
	if err := j.validateSetAddressLine1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine1",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetAddressLine2(val *string) {
	if err := j.validateSetAddressLine2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine2",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetContactType(val *string) {
	if err := j.validateSetContactTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactType",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetCountryCode(val *string) {
	if err := j.validateSetCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countryCode",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetFax(val *string) {
	if err := j.validateSetFaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fax",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetOrganizationName(val *string) {
	if err := j.validateSetOrganizationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationName",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsDomainTechContactOutputReference)SetZipCode(val *string) {
	if err := j.validateSetZipCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipCode",
		val,
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) PutExtraParam(value interface{}) {
	if err := r.validatePutExtraParamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putExtraParam",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetAddressLine1() {
	_jsii_.InvokeVoid(
		r,
		"resetAddressLine1",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetAddressLine2() {
	_jsii_.InvokeVoid(
		r,
		"resetAddressLine2",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetCity() {
	_jsii_.InvokeVoid(
		r,
		"resetCity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetContactType() {
	_jsii_.InvokeVoid(
		r,
		"resetContactType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetCountryCode() {
	_jsii_.InvokeVoid(
		r,
		"resetCountryCode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		r,
		"resetEmail",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetExtraParam() {
	_jsii_.InvokeVoid(
		r,
		"resetExtraParam",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetFax() {
	_jsii_.InvokeVoid(
		r,
		"resetFax",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetFirstName() {
	_jsii_.InvokeVoid(
		r,
		"resetFirstName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetLastName() {
	_jsii_.InvokeVoid(
		r,
		"resetLastName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetOrganizationName() {
	_jsii_.InvokeVoid(
		r,
		"resetOrganizationName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		r,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		r,
		"resetState",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ResetZipCode() {
	_jsii_.InvokeVoid(
		r,
		"resetZipCode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsDomainTechContactOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


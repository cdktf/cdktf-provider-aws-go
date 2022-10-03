package route53domainsregistereddomain

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-aws-go/aws/v10/jsii"

	"github.com/hashicorp/cdktf-provider-aws-go/aws/v10/route53domainsregistereddomain/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53DomainsRegisteredDomainAdminContactOutputReference interface {
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
	ExtraParams() *map[string]*string
	SetExtraParams(val *map[string]*string)
	ExtraParamsInput() *map[string]*string
	Fax() *string
	SetFax(val *string)
	FaxInput() *string
	FirstName() *string
	SetFirstName(val *string)
	FirstNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *Route53DomainsRegisteredDomainAdminContact
	SetInternalValue(val *Route53DomainsRegisteredDomainAdminContact)
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
	ResetAddressLine1()
	ResetAddressLine2()
	ResetCity()
	ResetContactType()
	ResetCountryCode()
	ResetEmail()
	ResetExtraParams()
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

// The jsii proxy struct for Route53DomainsRegisteredDomainAdminContactOutputReference
type jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) AddressLine1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) AddressLine1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) AddressLine2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) AddressLine2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressLine2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ContactType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ContactTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contactTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) CountryCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) CountryCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ExtraParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ExtraParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) Fax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) FaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"faxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) FirstName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) FirstNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) InternalValue() *Route53DomainsRegisteredDomainAdminContact {
	var returns *Route53DomainsRegisteredDomainAdminContact
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) LastName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) LastNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) OrganizationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) OrganizationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) PhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) PhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ZipCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ZipCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zipCodeInput",
		&returns,
	)
	return returns
}


func NewRoute53DomainsRegisteredDomainAdminContactOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Route53DomainsRegisteredDomainAdminContactOutputReference {
	_init_.Initialize()

	if err := validateNewRoute53DomainsRegisteredDomainAdminContactOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.route53DomainsRegisteredDomain.Route53DomainsRegisteredDomainAdminContactOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRoute53DomainsRegisteredDomainAdminContactOutputReference_Override(r Route53DomainsRegisteredDomainAdminContactOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.route53DomainsRegisteredDomain.Route53DomainsRegisteredDomainAdminContactOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetAddressLine1(val *string) {
	if err := j.validateSetAddressLine1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine1",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetAddressLine2(val *string) {
	if err := j.validateSetAddressLine2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addressLine2",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetContactType(val *string) {
	if err := j.validateSetContactTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contactType",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetCountryCode(val *string) {
	if err := j.validateSetCountryCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countryCode",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetExtraParams(val *map[string]*string) {
	if err := j.validateSetExtraParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraParams",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetFax(val *string) {
	if err := j.validateSetFaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fax",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetFirstName(val *string) {
	if err := j.validateSetFirstNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstName",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetInternalValue(val *Route53DomainsRegisteredDomainAdminContact) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetLastName(val *string) {
	if err := j.validateSetLastNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastName",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetOrganizationName(val *string) {
	if err := j.validateSetOrganizationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationName",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetPhoneNumber(val *string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference)SetZipCode(val *string) {
	if err := j.validateSetZipCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zipCode",
		val,
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetAddressLine1() {
	_jsii_.InvokeVoid(
		r,
		"resetAddressLine1",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetAddressLine2() {
	_jsii_.InvokeVoid(
		r,
		"resetAddressLine2",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetCity() {
	_jsii_.InvokeVoid(
		r,
		"resetCity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetContactType() {
	_jsii_.InvokeVoid(
		r,
		"resetContactType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetCountryCode() {
	_jsii_.InvokeVoid(
		r,
		"resetCountryCode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetEmail() {
	_jsii_.InvokeVoid(
		r,
		"resetEmail",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetExtraParams() {
	_jsii_.InvokeVoid(
		r,
		"resetExtraParams",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetFax() {
	_jsii_.InvokeVoid(
		r,
		"resetFax",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetFirstName() {
	_jsii_.InvokeVoid(
		r,
		"resetFirstName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetLastName() {
	_jsii_.InvokeVoid(
		r,
		"resetLastName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetOrganizationName() {
	_jsii_.InvokeVoid(
		r,
		"resetOrganizationName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		r,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		r,
		"resetState",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ResetZipCode() {
	_jsii_.InvokeVoid(
		r,
		"resetZipCode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_Route53DomainsRegisteredDomainAdminContactOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


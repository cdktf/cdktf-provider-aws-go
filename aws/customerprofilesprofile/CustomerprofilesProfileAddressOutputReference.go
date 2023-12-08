// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package customerprofilesprofile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v18/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v18/customerprofilesprofile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CustomerprofilesProfileAddressOutputReference interface {
	cdktf.ComplexObject
	Address1() *string
	SetAddress1(val *string)
	Address1Input() *string
	Address2() *string
	SetAddress2(val *string)
	Address2Input() *string
	Address3() *string
	SetAddress3(val *string)
	Address3Input() *string
	Address4() *string
	SetAddress4(val *string)
	Address4Input() *string
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
	Country() *string
	SetCountry(val *string)
	CountryInput() *string
	County() *string
	SetCounty(val *string)
	CountyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CustomerprofilesProfileAddress
	SetInternalValue(val *CustomerprofilesProfileAddress)
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
	Province() *string
	SetProvince(val *string)
	ProvinceInput() *string
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
	ResetAddress1()
	ResetAddress2()
	ResetAddress3()
	ResetAddress4()
	ResetCity()
	ResetCountry()
	ResetCounty()
	ResetPostalCode()
	ResetProvince()
	ResetState()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CustomerprofilesProfileAddressOutputReference
type jsiiProxy_CustomerprofilesProfileAddressOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Address1() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Address1Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Address2() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Address2Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Address3() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Address3Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Address4() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Address4Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address4Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) City() *string {
	var returns *string
	_jsii_.Get(
		j,
		"city",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) CityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) County() *string {
	var returns *string
	_jsii_.Get(
		j,
		"county",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) CountyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) InternalValue() *CustomerprofilesProfileAddress {
	var returns *CustomerprofilesProfileAddress
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Province() *string {
	var returns *string
	_jsii_.Get(
		j,
		"province",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ProvinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCustomerprofilesProfileAddressOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomerprofilesProfileAddressOutputReference {
	_init_.Initialize()

	if err := validateNewCustomerprofilesProfileAddressOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesProfileAddressOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.customerprofilesProfile.CustomerprofilesProfileAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomerprofilesProfileAddressOutputReference_Override(c CustomerprofilesProfileAddressOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.customerprofilesProfile.CustomerprofilesProfileAddressOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetAddress1(val *string) {
	if err := j.validateSetAddress1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address1",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetAddress2(val *string) {
	if err := j.validateSetAddress2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address2",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetAddress3(val *string) {
	if err := j.validateSetAddress3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address3",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetAddress4(val *string) {
	if err := j.validateSetAddress4Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address4",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetCity(val *string) {
	if err := j.validateSetCityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"city",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetCountry(val *string) {
	if err := j.validateSetCountryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetCounty(val *string) {
	if err := j.validateSetCountyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"county",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetInternalValue(val *CustomerprofilesProfileAddress) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetPostalCode(val *string) {
	if err := j.validateSetPostalCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetProvince(val *string) {
	if err := j.validateSetProvinceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"province",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesProfileAddressOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ResetAddress1() {
	_jsii_.InvokeVoid(
		c,
		"resetAddress1",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ResetAddress2() {
	_jsii_.InvokeVoid(
		c,
		"resetAddress2",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ResetAddress3() {
	_jsii_.InvokeVoid(
		c,
		"resetAddress3",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ResetAddress4() {
	_jsii_.InvokeVoid(
		c,
		"resetAddress4",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ResetCity() {
	_jsii_.InvokeVoid(
		c,
		"resetCity",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ResetCountry() {
	_jsii_.InvokeVoid(
		c,
		"resetCountry",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ResetCounty() {
	_jsii_.InvokeVoid(
		c,
		"resetCounty",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ResetPostalCode() {
	_jsii_.InvokeVoid(
		c,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ResetProvince() {
	_jsii_.InvokeVoid(
		c,
		"resetProvince",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ResetState() {
	_jsii_.InvokeVoid(
		c,
		"resetState",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesProfileAddressOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


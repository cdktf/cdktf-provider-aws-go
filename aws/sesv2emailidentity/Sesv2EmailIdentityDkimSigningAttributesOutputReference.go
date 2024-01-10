// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2emailidentity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-aws-go/aws/v19/jsii"

	"github.com/cdktf/cdktf-provider-aws-go/aws/v19/sesv2emailidentity/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Sesv2EmailIdentityDkimSigningAttributesOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	CurrentSigningKeyLength() *string
	DomainSigningPrivateKey() *string
	SetDomainSigningPrivateKey(val *string)
	DomainSigningPrivateKeyInput() *string
	DomainSigningSelector() *string
	SetDomainSigningSelector(val *string)
	DomainSigningSelectorInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *Sesv2EmailIdentityDkimSigningAttributes
	SetInternalValue(val *Sesv2EmailIdentityDkimSigningAttributes)
	LastKeyGenerationTimestamp() *string
	NextSigningKeyLength() *string
	SetNextSigningKeyLength(val *string)
	NextSigningKeyLengthInput() *string
	SigningAttributesOrigin() *string
	Status() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Tokens() *[]*string
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
	ResetDomainSigningPrivateKey()
	ResetDomainSigningSelector()
	ResetNextSigningKeyLength()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Sesv2EmailIdentityDkimSigningAttributesOutputReference
type jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) CurrentSigningKeyLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"currentSigningKeyLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) DomainSigningPrivateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningPrivateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) DomainSigningPrivateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningPrivateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) DomainSigningSelector() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningSelector",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) DomainSigningSelectorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainSigningSelectorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) InternalValue() *Sesv2EmailIdentityDkimSigningAttributes {
	var returns *Sesv2EmailIdentityDkimSigningAttributes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) LastKeyGenerationTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastKeyGenerationTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) NextSigningKeyLength() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSigningKeyLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) NextSigningKeyLengthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nextSigningKeyLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) SigningAttributesOrigin() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signingAttributesOrigin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) Tokens() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokens",
		&returns,
	)
	return returns
}


func NewSesv2EmailIdentityDkimSigningAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Sesv2EmailIdentityDkimSigningAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewSesv2EmailIdentityDkimSigningAttributesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.sesv2EmailIdentity.Sesv2EmailIdentityDkimSigningAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSesv2EmailIdentityDkimSigningAttributesOutputReference_Override(s Sesv2EmailIdentityDkimSigningAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.sesv2EmailIdentity.Sesv2EmailIdentityDkimSigningAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference)SetDomainSigningPrivateKey(val *string) {
	if err := j.validateSetDomainSigningPrivateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSigningPrivateKey",
		val,
	)
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference)SetDomainSigningSelector(val *string) {
	if err := j.validateSetDomainSigningSelectorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainSigningSelector",
		val,
	)
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference)SetInternalValue(val *Sesv2EmailIdentityDkimSigningAttributes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference)SetNextSigningKeyLength(val *string) {
	if err := j.validateSetNextSigningKeyLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nextSigningKeyLength",
		val,
	)
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) ResetDomainSigningPrivateKey() {
	_jsii_.InvokeVoid(
		s,
		"resetDomainSigningPrivateKey",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) ResetDomainSigningSelector() {
	_jsii_.InvokeVoid(
		s,
		"resetDomainSigningSelector",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) ResetNextSigningKeyLength() {
	_jsii_.InvokeVoid(
		s,
		"resetNextSigningKeyLength",
		nil, // no parameters
	)
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_Sesv2EmailIdentityDkimSigningAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

